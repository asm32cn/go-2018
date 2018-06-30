// go-base64-demo-1.go
package main;
import(
	"strconv"
	"encoding/base64"
);

func main(){
	var strSource = "go-base64-demo-1.go\n程序中书写着所见所闻所感，编译着心中的万水千山。";
	// Z28tYmFzZTY0LWRlbW8tMS5nbwrnqIvluo/kuK3kuablhpnnnYDmiYDop4HmiYDpl7vmiYDmhJ/vvIznvJbor5HnnYDlv4PkuK3nmoTkuIfmsLTljYPlsbHjgII=
	println(strSource);
	strStdEncrypt := base64.StdEncoding.EncodeToString([]byte(strSource));
	println(strStdEncrypt);
	println(Base64Encode(strSource));
}

func Base64Encode(s string) string{
	var bytes = []byte(s);
	var (
		// i = 0;
		nCount = len(bytes)
	);
	return strconv.Itoa(len(s)) + " " + strconv.Itoa(nCount) + " " + s;
}

/*
go-base64-demo-1.go
程序中书写着所见所闻所感，编译着心中的万水千山。
Z28tYmFzZTY0LWRlbW8tMS5nbwrnqIvluo/kuK3kuablhpnnnYDmiYDop4HmiYDpl7vmiYDmhJ/vvIznvJbor5HnnYDlv4PkuK3nmoTkuIfmsLTljYPlsbHjgII=
92 go-base64-demo-1.go
程序中书写着所见所闻所感，编译着心中的万水千山。
*/