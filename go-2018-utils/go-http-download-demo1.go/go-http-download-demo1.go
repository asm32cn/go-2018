// go-http-download-demo1.go
package main;

import (
	"io/ioutil"
	"net/http"
	"flag"
	"strings"
);

// func GetHttpData(url string) (content string, statusCode int){
func GetHttpData(url string) (content []byte, statusCode int){
	resp, err1 := http.Get(url);
	if err1 != nil {
		statusCode = -100;
		return
	}
	defer resp.Body.Close();
	data, err2 := ioutil.ReadAll(resp.Body);
	if err2 != nil {
		statusCode = -200;
		return;
	}
	statusCode = resp.StatusCode;
	// content = string(data);
	content = data;
	return;
}

func main(){
	// url := "http://localhost/EapWeb-SmbPm/Resource/zh-cn/libs/ui/seeker.ui.grid_res.js";

	flag.Parse();
	url := flag.Arg(0);
	if url == "" {
		url = "http://localhost/EapWeb-SmbPm/Resource/zh-cn/libs/ui/seeker.ui.grid_res.js";
	}
	pos := strings.LastIndex(url, "/");
	file := url;
	if pos >= 0 {
		file = url[pos + 1:];
	}

	data, statusCode := GetHttpData(url);


	if file != "" && data != nil {
		err := ioutil.WriteFile(file, data, 0644);
		if err != nil { panic(err) }
	}

	println("file = ", file)

	println("url =", url);
	println("statusCode =", statusCode);
	println("data =", string(data));
	println("Done.");
}