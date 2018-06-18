// go-cocktailsort-demo-1.go
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

func CocktailSort(data []int){
	var left = 0
	var right = len(data) - 1;
	for left < right {
		for i := left; i < right; i++ {
			if data[i] > data[i + 1] {
				data[i], data[i + 1] = data[i + 1], data[i];
			}
		}
		right--;
		for i := right; i > left; i-- {
			if data[i - 1] > data[i] {
				data[i - 1], data[i] = data[i], data[i - 1];
			}
		}
		left++;
	}
}

func main(){
	// var data = []int{41, 67, 34, 0, 69, 24, 78, 58, 62, 64, 5, 45, 81, 27, 61, 91, 95, 42, 27, 36};
	var data = []int{76, 11, 11, 43, 78, 35, 39, 27, 16, 55, 1, 41, 24, 19, 54, 7, 78, 69, 65, 82};

	DisplayData(data);
	CocktailSort(data);
	DisplayData(data);
}