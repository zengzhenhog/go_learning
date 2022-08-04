package main

import (
	"fmt"
	retrieverStruct "helloword/retriever" // 包名与文件夹名不一致时
	"time"
)

// something that can Get
type Retriever interface {
	Get(string) string
}

func inspect(r Retriever) {
	fmt.Printf("%T, %v\n", r, r)
	switch v := r.(type) {
	case retrieverStruct.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	case *retrieverStruct.Retriever:
		fmt.Println("UserAgents:", v.UserAgent)
	}
}

func main() {
	r := retrieverStruct.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	x := &retrieverStruct.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(x)

	var value interface{} = "GeeksforGeeks"

	var value1, ok = value.(string)
	fmt.Println(value1, ok)

	// type assertion 会报错
	if realRetriever, ok := r.(retrieverStruct.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("realRetriever is not retrieverStruct Retriever")
	}

}
