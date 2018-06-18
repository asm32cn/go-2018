// go-bucket-sort-demo-1.go
package main;

const MAX int = 100;
const bn int = 5;
var C = make([]int, bn);

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

func InsertionSort(data []int, left int, right int){
	for i := left + 1; i <= right; i++ {
		var get = data[i];
		var j = i - 1;
		for j >= left && data[j] > get {
			data[j + 1] = data[j];
			j--;
		}
		data[j + 1] = get;
	}
}

func CountingSort(data []int){
	var n = len(data);
	var nFactor int = MAX / bn;

	if MAX % bn != 0 {
		nFactor++;
	}
	var MapToBucket = func(x int) int{
		return x / nFactor;
	}

	for i := 0; i < bn; i++ {
		C[i] = 0;
	}
	for i := 0; i < n; i++ {
		C[MapToBucket(data[i])]++;
	}
	for i := 1; i < bn; i++ {
		C[i] = C[i] + C[i - 1];
	}
	var B = make([]int, n);
	for i := n - 1; i >= 0; i-- {
		var b = MapToBucket(data[i]);
		C[b]--;
		B[C[b]] = data[i];
	}
	for i := 0; i < n; i++ {
		data[i] = B[i];
	}
}

func BucketSort(data []int){
	var n = len(data);
	CountingSort(data);
	for i := 0; i < bn; i++ {
		var left = C[i];
		var right int;
		if i == bn - 1 {
			right = n - 1;
		}else{
			right = C[i + 1] - 1;
		}
		if left < right {
			InsertionSort(data, left, right);
		}
	}
}

func main(){
	// var data = []int{41, 67, 34, 0, 69, 24, 78, 58, 62, 64, 5, 45, 81, 27, 61, 91, 95, 42, 27, 36};
	var data = []int{76, 11, 11, 43, 78, 35, 39, 27, 16, 55, 1, 41, 24, 19, 54, 7, 78, 69, 65, 82};

	DisplayData(data);
	BucketSort(data);
	DisplayData(data);
}