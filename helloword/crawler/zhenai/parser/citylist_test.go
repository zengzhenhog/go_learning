package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParserCityList(contents)
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if url != result.Requests[i].URL {
			t.Errorf("expected url #%d: %s, but was %s", i, url, result.Requests[i].URL)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d Items; but had %d", resultSize, len(result.Items))
	}
	for i, city := range expectedCities {
		if city != result.Items[i].(string) {
			t.Errorf("expected city #%d: %s, but was %s", i, city, result.Items[i].(string))
		}
	}
}
