// go-pointers.go
package main

import "fmt"

var ip *int	/* 指向整形 */
var fp *float32 /* 指向浮点型 */

func main(){
	var a int = 10

	fmt.Printf("变量的地址: %x\n", &a)
}