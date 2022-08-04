package parser

import (
	"helloword/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParserProfile(contents, "希恩")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Age:       32,
		WorkPlace: "阿坝金川",
		Name:      "希恩",
	}

	if profile != expected {
		t.Errorf("expected %v, but was %v", expected, profile)
	}
}
