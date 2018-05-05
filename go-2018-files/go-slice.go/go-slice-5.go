// go-slice-5.go
package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	fmt.Println("numbers ==", numbers)
	numbers1 := numbers[1:4]
	printSlice(numbers1)
	fmt.Println("numbers[:3] ==", numbers[:3])
	numbers2 := numbers[4:]
	printSlice(numbers2)
	numbers3 := make([]int, 0, 5)
	printSlice(numbers3)
	numbers4 := numbers[:2]
	printSlice(numbers4)
	println("demo string"[2:10])
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

/*
len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
numbers == [0 1 2 3 4 5 6 7 8]
len=3 cap=8 slice=[1 2 3]
numbers[:3] == [0 1 2]
len=5 cap=5 slice=[4 5 6 7 8]
len=0 cap=5 slice=[]
len=2 cap=9 slice=[0 1]
*/