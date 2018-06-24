// go-sort-tiobe-filename-demo-1.go
package main;
import("fmt";"regexp";"strconv");

func main(){
	var A_strFiles = []string{
		"TIOBE Index for April 2018.html",
		"TIOBE Index for February 2018.html",
		"TIOBE Index for January 2018.html",
		"TIOBE Index for June 2018.html",
		"TIOBE Index for March 2018.html",
		"TIOBE Index for May 2018.html",
		"TIOBE-exchange-matrix-data.html",
		"TIOBE-exchange-matrix-data.py",
		"TIOBE-gernate-index-py2.py",
		"TIOBE-index.html",
		"TIOBE_matrixData.txt"};
	var nCount = len(A_strFiles);
	var sorted = sort_tiobe_filename_demo(A_strFiles);
	fmt.Printf("%2s %-36s %-36s\n", "@", "Source Data", "Sorted data");
	fmt.Printf("%2s %-36s %-36s\n", "==", "==================", "==================");
	for i := 0; i < nCount; i++ {
		fmt.Printf("%2d %-36s %-36s\n", i, A_strFiles[i], sorted[i]);
		// println(A_strFiles[i]);
	}
}

func sort_tiobe_filename_demo(files []string) []string{
	var _month = map[string]int{"January": 1, "February": 2, "March": 3, "April": 4, "May": 5, "June": 6,
		"July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12};
	var _regex = regexp.MustCompile("^TIOBE Index for (\\w+) (\\d{4})\\.html$");
	var nCount = len(files);
	var buff = make([][]int, nCount);
	var sorted = make([]string, nCount);
	var n = 0;

	for i := 0; i < nCount; i++ {
		buff[i] = make([]int, 3);
		buff[i][0] = i;
		buff[i][1] = i;

		var m = _regex.FindAllSubmatch([]byte(files[i]), -1);
		if len(m) > 0 {
			var nYear, _ = strconv.Atoi(string(m[0][2]));
			buff[i][2] = nYear * 100 + _month[string(m[0][1])];
		} else {
			n++;
			buff[i][2] = n;
		};
	}

	var GetIndex = func(i int) int{	return buff[i][1]; }
	var GetData = func(i int) int{ return buff[GetIndex(i)][2]; }
	var SetIndex = func(i int, n int){ buff[i][1] = n; }

	// sort string
	for i := 0; i < nCount - 1; i++ {
		var nMin = i;
		for j := i + 1; j < nCount; j++ {
			if GetData(j) < GetData(nMin) {
				nMin = j;
			}
		}
		if i != nMin {
			var t = GetIndex(i);
			SetIndex(i, GetIndex(nMin));
			SetIndex(nMin, t);
		}
	}

	for i := 0; i < nCount; i++ {
		sorted[i] = files[GetIndex(i)];
	}
	return sorted;
}