// go-pointers-3.go
package main

import "fmt"

func main(){
	var ptr *int

	fmt.Printf("ptr的值为: %x\n", ptr)
	println( "ptr == nil ", ptr == nil )
}