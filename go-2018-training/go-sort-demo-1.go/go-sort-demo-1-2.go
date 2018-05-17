// go-sort-demo-1-2.go
package main

import (
	"fmt"
)

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	selectionSort( items )
	fmt.Println( items )
}

func selectionSort(items []int){
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}