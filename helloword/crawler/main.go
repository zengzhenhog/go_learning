package main

import (
	"helloword/crawler/engine"
	"helloword/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		URL:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
