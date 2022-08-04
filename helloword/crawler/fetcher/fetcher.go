package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

// func determineEncoding(r io.Reader) encoding.Encoding {
// 	bytes, err := bufio.NewReader(r).Peek(1024)
// 	if err != nil {
// 		log.Printf("Fetcher error %v", err)
// 		return unicode.UTF8
// 	}
// 	e, _, _ := charset.DetermineEncoding(bytes, "")
// 	return e
// }

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code with:", resp.StatusCode)
		return nil, fmt.Errorf("error: status code with: %d", resp.StatusCode)
		// return nil, errors.New()
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	// e := determineEncoding(resp.Body) // 判断编码类型
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	// utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	// all, err := ioutil.ReadAll(resp.Body)
	return ioutil.ReadAll(utf8Reader) // utf-8 -> gbk
}
