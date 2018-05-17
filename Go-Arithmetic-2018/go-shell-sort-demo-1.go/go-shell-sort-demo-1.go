// go-shell-sort-demo-1.go
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

func ShellSort(data []int){
	var n = len(data);
	var h = 0;
	for h <= n {
		h = 3 * h + 1;
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			j := i - h;
			get := data[i];
			for j >= 0 && data[j] > get {
				data[j + h] = data[j];
				j = j - h;
			}
			data[j + h] = get;
		}
		h = (h - 1) / 3;
	}
}

func main(){
	// var data = []int{41, 67, 34, 0, 69, 24, 78, 58, 62, 64, 5, 45, 81, 27, 61, 91, 95, 42, 27, 36};
	var data = []int{76, 11, 11, 43, 78, 35, 39, 27, 16, 55, 1, 41, 24, 19, 54, 7, 78, 69, 65, 82};

	DisplayData(data);
	ShellSort(data);
	DisplayData(data);
}