// go-rotating-hash-demo-1.go
package main;
import "fmt";

func RotatingHash(s string, nPrime int) int{
	var (
		n = len(s);
		nHash = n;
	);
	for i := 0; i < n; i++ {
		nHash = (nHash << 4 >> 28) ^ int(s[i]);
	}
	return nHash % nPrime;
}

func main(){
	var A_strKeys = []string{"C", "C++", "Java", "C#", "Python", "Go", "Scala", "vb.net", "JavaScript", "PHP", "Perl", "Ruby"};
	var n = len(A_strKeys);

	for i := 0; i < n; i++ {
		fmt.Printf("%-10d %-15s %3d\n", i, A_strKeys[i], RotatingHash(A_strKeys[i], 31));
	}
}

/*
0          C                 5
1          C++              12
2          Java              4
3          C#                4
4          Python           17
5          Go               18
6          Scala             4
7          vb.net           23
8          JavaScript       23
9          PHP              18
10         Perl             15
11         Ruby             28
*/