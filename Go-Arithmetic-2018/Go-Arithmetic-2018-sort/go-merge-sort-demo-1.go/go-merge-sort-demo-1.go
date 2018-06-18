// go-merge-sort-demo-1.go
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

func Merge(data []int, left int, mid int, right int){
	var n = right - left + 1;
	var temp = make([]int, n);
	var index = 0;
	var i = left;
	var j = mid + 1;
	for i <= mid && j <= right {
		if data[i] <= data[j] {
			temp[index] = data[i];
			i++;
		}else{
			temp[index] = data[j]
			j++;
		}
		index++;
	}
	for i <= mid {
		temp[index] = data[i];
		index++;
		i++;
	}
	for j <= right {
		temp[index] = data[j];
		index++;
		j++;
	}
	for k := 0; k < n; k++ {
		data[left] = temp[k];
		left++;
	}
}

func MergeSortRecursion(data []int, left int, right int){
	if left == right {
		return;
	}

	var mid = (left + right) / 2;
	MergeSortRecursion(data, left, mid);
	MergeSortRecursion(data, mid + 1, right);
	Merge(data, left, mid, right);
}

func MergeSortIteration(data []int){
	var n = len(data);
	var(
		left int;
		mid int;
		right int;
	);
	for i := 1; i < n; i *= 2 {
		left = 0;
		for left + i < n {
			mid = left + i - 1;
			if mid + i < n {
				right = mid + i;
			}else{
				right = n - 1;
			}
			Merge(data, left, mid, right);
			left = right + 1;
		}
	}
}

func main(){
	// var data = []int{41, 67, 34, 0, 69, 24, 78, 58, 62, 64, 5, 45, 81, 27, 61, 91, 95, 42, 27, 36};
	var data1 = []int{76, 11, 11, 43, 78, 35, 39, 27, 16, 55, 1, 41, 24, 19, 54, 7, 78, 69, 65, 82};
	var data2 = []int{76, 11, 11, 43, 78, 35, 39, 27, 16, 55, 1, 41, 24, 19, 54, 7, 78, 69, 65, 82};

	DisplayData(data1);
	MergeSortRecursion(data1, 0, len(data1) - 1);
	DisplayData(data1);

	println();
	DisplayData(data2);
	MergeSortIteration(data2);
	DisplayData(data2);
}