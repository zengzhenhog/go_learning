package main

import (
	"fmt"
	"regexp"
)

// const text = "my email is zzh@qq.com"
const text = `
my email is zzh@qq.com
her email is clh@ss.com
whos email is sdfw@baidu.com.cn
`

func main() {
	// re, err := regexp.Compile("zzh@qq.com")
	// re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`) // []string
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`) // [][]string
	// match := re.FindString(text) // 匹配第一个
	// match := re.FindAllString(text, -1) // 匹配所有 返回[]string
	match := re.FindAllStringSubmatch(text, -1) // 匹配所有和子匹配，正则表达式里括号包起来的内容 返回[][]string
	fmt.Println(match)
}
