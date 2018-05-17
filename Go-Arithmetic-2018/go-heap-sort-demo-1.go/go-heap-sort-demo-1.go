// go-heap-sort-demo-1.go
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

func Heapify(data []int, i int, size int){
	var left_child = 2 * i + 1;
	var right_child = 2 * i + 2;
	var max = i;
	if left_child < size && data[left_child] > data[max] {
		max = left_child;
	}
	if right_child < size && data[right_child] > data[max] {
		max = right_child;
	}
	if max != i {
		data[i], data[max] = data[max], data[i];
		Heapify(data, max, size);
	}
}

func BuildHeap(data []int, n int) int{
	var heap_size = n;
	for i := heap_size / 2 - 1; i >= 0; i-- {
		Heapify(data, i, heap_size);
	}
	return heap_size;
}

func HeapSort(data []int, n int){
	heap_size := BuildHeap(data, n);
	for heap_size > 1 {
		heap_size--;
		data[0], data[heap_size] = data[heap_size], data[0];
		Heapify(data, 0, heap_size);
	}
}

func main(){
	// var data = []int{41, 67, 34, 0, 69, 24, 78, 58, 62, 64, 5, 45, 81, 27, 61, 91, 95, 42, 27, 36};
	var data = []int{76, 11, 11, 43, 78, 35, 39, 27, 16, 55, 1, 41, 24, 19, 54, 7, 78, 69, 65, 82};

	DisplayData(data);
	HeapSort(data, len(data));
	DisplayData(data);
}