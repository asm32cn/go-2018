// CopyFile.go
package main

import (
	"io"
	"os"
)

func CopyFile(dest, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}

	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return
	}

	defer destFile.Close()
	return io.Copy(destFile, srcFile)
}

func main() {
	println("copy CopyFile.exe CopyFile~.exe")
	CopyFile("CopyFile~.exe", "CopyFile.exe")
}
