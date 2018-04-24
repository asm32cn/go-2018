// go-spider-1.go
package main

import (
	"io/ioutil"
	"net/http"
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
	url := "http://10.52.10.211/asm32.article/?id=29"
	data, statusCode := GetHttpData(url)

	println("url =", url)
	println("statusCode =", statusCode)
	println("data =", data)
}