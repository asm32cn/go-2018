// go-spider-2.go
package main

import (
	"io/ioutil"
	"net/http"
	"flag"
)

func GetHttpData(url string) (content string, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)
	return
}

func main() {
	flag.Parse()
	url := flag.Arg(0)
	data, statusCode := GetHttpData(url)

	println("url =", url)
	println("statusCode =", statusCode)
	println("data =", data)
}