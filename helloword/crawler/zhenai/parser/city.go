package parser

import (
	"helloword/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	result := engine.ParserResult{}
	matches := re.FindAllSubmatch(contents, -1) // 返回  [][][]byte -> [][]string
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User"+name)
		result.Requests = append(result.Requests, engine.Request{
			URL: string(m[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParserProfile(c, name)
			},
		})
	}

	return result
}
