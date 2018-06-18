// go-additive-hash-demo-1.go
package main;
import "fmt";

func AdditiveHash(s string, nPrime int) int{
	var (
		n = len(s);
		nHash = n;
	);
	for i := 0; i < n; i++ {
		nHash += int(s[i]);
	}
	return nHash % nPrime;
}

func main(){
	var A_strKeys = []string{"C", "C++", "Java", "C#", "Python", "Go", "Scala", "vb.net", "JavaScript", "PHP", "Perl", "Ruby"};
	var n = len(A_strKeys);
	for i := 0; i < n; i++ {
		fmt.Printf("%-10d %-15s %3d\n", i, A_strKeys[i], AdditiveHash(A_strKeys[i], 31));
	}
}

/*
0          C                 6
1          C++               1
2          Java             18
3          C#               11
4          Python           28
5          Go               29
6          Scala            24
7          vb.net            6
8          JavaScript        2
9          PHP              18
10         Perl              4
11         Ruby             19
*/