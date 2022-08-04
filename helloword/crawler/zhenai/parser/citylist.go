package parser

import (
	"helloword/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	// matches := re.FindAll(contents, -1)
	// for _, m := range matches {
	// 	fmt.Printf("%s\n", m)
	// }
	result := engine.ParserResult{}
	matches := re.FindAllSubmatch(contents, -1) // 返回  [][][]byte -> [][]string
	limit := 3
	for _, m := range matches {
		result.Items = append(result.Items, "City"+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: ParserCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}

	return result
}
