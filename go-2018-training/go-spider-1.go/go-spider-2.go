// go-spider-2.go
// 2018-03-06 modified by PASCAL asm32 gyl
package main;

import (
	"io/ioutil"
	"net/http"
	"flag"
);

func GetHttpData(url string) (content string, statusCode int){
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
	content = string(data);
	return;
}

func main(){
	// url := "http://localhost/EapWeb-SmbPm/Resource/zh-cn/libs/ui/seeker.ui.grid_res.js";

	flag.Parse();
	url := flag.Arg(0);
	if url == "" {
		url = "http://localhost/EapWeb-SmbPm/Resource/zh-cn/libs/ui/seeker.ui.grid_res.js";
	}
	data, statusCode := GetHttpData(url);

	println("url =", url);
	println("statusCode =", statusCode);
	println("data =", data);
	println("Done.");
}
