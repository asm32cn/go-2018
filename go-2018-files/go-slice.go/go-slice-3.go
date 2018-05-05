// go-slice-3.go
package main

import "fmt"

func main(){
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含)到索引4(不包含) */
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为0 */
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s) */
	fmt.Println("numbers[4:] ==", numbers[4:])
	
	number1 := make([]int, 0, 5)
	printSlice(number1)

	/* 打印子切片从索引0(包含)到索引5(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引2(包含)到索引5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

/*
len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
numbers == [0 1 2 3 4 5 6 7 8]
numbers[1:4] == [1 2 3]
numbers[:3] == [0 1 2]
numbers[4:] == [4 5 6 7 8]
len=0 cap=5 slice=[]
len=2 cap=9 slice=[0 1]
len=3 cap=7 slice=[2 3 4]
*/
