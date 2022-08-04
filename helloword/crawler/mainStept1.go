// package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	// matches := re.FindAll(contents, -1)
	// for _, m := range matches {
	// 	fmt.Printf("%s\n", m)
	// }

	matches := re.FindAllSubmatch(contents, -1) // 返回  [][][]byte -> [][]string
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}

func main() {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code with:", resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body) // 判断编码类型
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	// all, err := ioutil.ReadAll(resp.Body)
	all, err := ioutil.ReadAll(utf8Reader) // utf-8 -> gbk
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", all)
	printCityList(all)
}
