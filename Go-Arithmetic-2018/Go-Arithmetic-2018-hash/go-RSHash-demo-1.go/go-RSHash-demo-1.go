// go-RSHash-demo-1.go
package main;

import ("fmt");

func RSHash(s string) int{
	var(
		n = len(s);
		b = 378551;
		a = 63689;
		nHash = 0;
	);
	for i := 0; i < n; i++ {
		nHash = nHash * a + int(s[i]);
		a = a * b;
	}
	return nHash & 0x7FFFFFFF;
}

func main(){
	var A_strKeys = []string{"C", "C++", "Java", "C#", "Python", "Go", "Scala", "vb.net", "JavaScript", "PHP", "Perl", "Ruby"};
	var n = len(A_strKeys);

	for i := 0; i < n; i++ {
		fmt.Printf("%-10d %-15s %3d\n", i, A_strKeys[i], RSHash(A_strKeys[i]) % 31);
	}
}

/*
0          C                 5
1          C++               7
2          Java              2
3          C#               25
4          Python           19
5          Go               22
6          Scala            14
7          vb.net           18
8          JavaScript       17
9          PHP              22
10         Perl              4
11         Ruby              9
*/