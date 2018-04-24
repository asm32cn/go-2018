// go-file-3.go
package main

import (
	"io"
	"os"
)

func main(){
	fi, err := os.Open("readme.txt");
	if err != nil { panic(err); }
	defer fi.Close();

	fo, err := os.Create("readme-2.txt");
	if err != nil { panic(err); }
	defer fo.Close();

	buf := make([]byte, 1024);
	for{
		n, err := fi.Read(buf);
		if err != nil && err != io.EOF { panic(err) }
		if n == 0 { break }

		if n2, err := fo.Write(buf[:n]); err != nil {
			panic(err);
		}else if n2 != n {
			panic("error in writing")
		}
	}
}