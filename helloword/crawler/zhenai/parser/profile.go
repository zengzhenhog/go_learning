package parser

import (
	"helloword/crawler/engine"
	"helloword/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`>([\d]+)岁</div>`)
var workRe = regexp.MustCompile(`>工作地:([^<]+)</div>`)

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return "广州"
	}
}

func ParserProfile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))

	if err == nil {
		profile.Age = age
	} else {
		profile.Age = 29
	}

	profile.WorkPlace = extractString(contents, workRe)

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}

	return result
}
