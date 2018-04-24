// go-file-4.go
package main

import (
	"bufio"
	"io"
	"os"
)

func main(){
	fi, err := os.Open("readme.txt");
	if err != nil { panic(err) }
	defer fi.Close();
	r := bufio.NewReader(fi);

	fo, err := os.Create("readme-2.txt");
	if err != nil { panic(err) }
	defer fo.Close();
	w := bufio.NewWriter(fo);

	buf := make([]byte, 1024);
	for{
		n, err := r.Read(buf);
		if err != nil && err != io.EOF { panic(err) }
		if n == 0 { break }

		if n2, err := w.Write(buf[:n]); err != nil {
			panic(err);
		} else if n2 != n {
			panic("err in writing");
		}
	}

	if err = w.Flush(); err != nil { panic(err) }
}
