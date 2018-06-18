// go-insertion-sort-dichotomy-demo-1.go
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

func InsertionSortDichotomy(data []int){
	var n = len(data);
	for i := 1; i < n; i++ {
		get := data[i];
		left := 0;
		right := i - 1;
		for left <= right {
			mid := (left + right) / 2;
			if data[mid] > get {
				right = mid - 1;
			}else{
				left = mid + 1;
			}
		}
		for j := i - 1; j >= left; j-- {
			data[j + 1] = data[j];
		}
		data[left] = get;
	}
}

func main(){
	// var data = []int{41, 67, 34, 0, 69, 24, 78, 58, 62, 64, 5, 45, 81, 27, 61, 91, 95, 42, 27, 36};
	var data = []int{76, 11, 11, 43, 78, 35, 39, 27, 16, 55, 1, 41, 24, 19, 54, 7, 78, 69, 65, 82};

	DisplayData(data);
	InsertionSortDichotomy(data);
	DisplayData(data);
}