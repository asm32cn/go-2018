// go-radix-sort-demo-1.go
package main;

func DisplayData(data []int){
	var n = len(data);
	for i := 0; i < n; i++ {
		if i > 0 {
			print(", ");
		}
		print(data[i]);
	}
	println();
}

func LsdRedixSort(data []int){
	var n = len(data);
	const dn int = 3;
	const k int = 10;
	var C = [k]int{0};
	var radix = []int{1, 1, 10, 100};

	var GetDigit = func(x int, d int) int{
		return (x / radix[d]) % 10;
	}

	// LsdRedixSort
	for d := 1; d < dn; d++ {
		// CountingSort
		for i := 0; i < k; i++ {
			C[i] = 0;
		}

		for i := 0; i < n; i++ {
			C[GetDigit(data[i], d)]++;
		}
		for i := 1; i < k; i++ {
			C[i] = C[i] + C[i - 1];
		}
		var B = make([]int, n);
		for i := n - 1; i >= 0; i-- {
			dight := GetDigit(data[i], d)
			C[dight]--;
			B[C[dight]] = data[i];
		}
		for i := 0; i < n; i++ {
			data[i] = B[i];
		}
	}
}

func main(){
	// var data = []int{41, 67, 34, 0, 69, 24, 78, 58, 62, 64, 5, 45, 81, 27, 61, 91, 95, 42, 27, 36};
	var data = []int{76, 11, 11, 43, 78, 35, 39, 27, 16, 55, 1, 41, 24, 19, 54, 7, 78, 69, 65, 82};

	DisplayData(data);
	LsdRedixSort(data);
	DisplayData(data);
}