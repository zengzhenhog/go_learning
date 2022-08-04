package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errorWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error: %s", err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

var lastOccurred = make([]int, 0xffff) // 空间换时间，类似map的效果
func lengthOfNonRepeatingSubStr(s string) int {
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	// lastOccurred[0x65] = 1
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	// http.HandleFunc("/", errorWrapper(filehandler.HandFileList))

	// err := http.ListenAndServe(":8888", nil)
	// if err != nil {
	// 	panic(err)
	// }

	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello %d\n", i) // I/O操作会切换协程
			}
		}(i)
	}
	time.Sleep(time.Millisecond)

	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { // 需要传i，否则i值为外部变量，最终i=10会超出a的长度报错
			for {
				a[i]++
				runtime.Gosched() // 手动交出控制权，会切换协程
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a) // a[i]++在写入，同时读取会冲突，使用Channel解决
}
