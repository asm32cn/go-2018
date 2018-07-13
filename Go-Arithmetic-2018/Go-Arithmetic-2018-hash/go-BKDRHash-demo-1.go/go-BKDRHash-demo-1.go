// go-BKDRHash-demo-1.go
package main;

import ("fmt"; );

func BKDRHash(s string) int {
	var (
		nSeed = 131; // 31 131 1313 13131 131313 etc
		nHash = 0;
		n = len(s);
	);
	for i := 0; i < n; i++ {
		nHash = nHash * nSeed + int(s[i]);
	}
	return nHash & 0x7fffffff;
}

func main(){
	var A_strKeys = []string{"C", "C++", "Java", "C#", "Python", "Go", "Scala", "vb.net", "JavaScript", "PHP", "Perl", "Ruby"};
	var n = len(A_strKeys);

	for i := 0; i < n; i ++ {
		var nHash = BKDRHash(A_strKeys[i]);
		fmt.Printf("%-3d %-15s %10d %3d\n", i, A_strKeys[i], nHash, nHash % 31);
	}
}

/*
0   C                       67   5
1   C++                1155463   0
2   Java             168038906  27
3   C#                    8812   8
4   Python          1962499928   9
5   Go                    9412  19
6   Scala           1045413186   0
7   vb.net           763463135   2
8   JavaScript       557701633   8
9   PHP                1382392   9
10  Perl             181595583   1
11  Ruby             186364258   8
*/