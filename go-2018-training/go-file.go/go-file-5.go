// go-file-5.go
package main

import (
	"io/ioutil"
)

func main(){
	println("go-file-5.go");
	b, err := ioutil.ReadFile("readme.txt");
	if err != nil {
		panic(err);
	}else{
		println(string(b));
		err = ioutil.WriteFile("readme-2.txt", b, 0644);
		if err != nil { panic(err) }
	}
}
