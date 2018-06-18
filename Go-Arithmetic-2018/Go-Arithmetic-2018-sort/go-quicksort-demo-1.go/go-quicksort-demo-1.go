// go-quicksort-demo-1.go
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

func QuickSort(data []int, left int, right int){
	if left < right {
		key := data[left];
		low := left;
		high := right;
		for low < high {
			for low < high && data[high] >= key {
				high--;
			}
			data[low] = data[high];
			for low < high && data[low] <= key {
				low++;
			}
			data[high] = data[low];
		}
		data[low] = key;

		QuickSort(data, left, low - 1);
		QuickSort(data, low + 1, right);
	}
}

func main(){
	// var data = []int{41, 67, 34, 0, 69, 24, 78, 58, 62, 64, 5, 45, 81, 27, 61, 91, 95, 42, 27, 36};
	var data = []int{76, 11, 11, 43, 78, 35, 39, 27, 16, 55, 1, 41, 24, 19, 54, 7, 78, 69, 65, 82};

	DisplayData(data);
	QuickSort(data, 0, len(data) - 1);
	DisplayData(data);
}