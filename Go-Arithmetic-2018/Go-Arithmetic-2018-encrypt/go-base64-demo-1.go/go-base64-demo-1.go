// go-base64-demo-1.go
package main;
import(
	"bytes"
	// "strconv"
	"encoding/base64"
);

const _keyStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";

func main(){
	var strSource = "go-base64-demo-1.go\n程序中书写着所见所闻所感，编译着心中的万水千山。";
	// Z28tYmFzZTY0LWRlbW8tMS5nbwrnqIvluo/kuK3kuablhpnnnYDmiYDop4HmiYDpl7vmiYDmhJ/vvIznvJbor5HnnYDlv4PkuK3nmoTkuIfmsLTljYPlsbHjgII=
	println(strSource);
	strStdEncrypt := base64.StdEncoding.EncodeToString([]byte(strSource));
	println("strStdEncrypt:\n" + strStdEncrypt);
	strUserEncrypt := Base64Encode(strSource);
	println("strUserEncrypt:\n" + strUserEncrypt);
	byteDescript, _ := base64.StdEncoding.DecodeString(strUserEncrypt)
	println(string(byteDescript));
}

func Base64Encode(s string) string{
	var data = []byte(s);
	var ( i = 0; nCount = len(data); );
	var ( chr1 = 0; chr2 = 0; chr3 = 0; );
	var ( enc1 = 0; enc2 = 0; enc3 = 0; enc4 = 0; );
	var sb bytes.Buffer;
	for (i < nCount){
		chr1 = int(data[i]); i+= 1;
		enc1 = chr1 >> 2;
		if(i < nCount){
			chr2 = int(data[i]); i+= 1;
			enc2 = ((chr1 & 3) << 4) | ((chr2 & 0xf0) >> 4);
			if(i < nCount){
				chr3 = int(data[i]); i+= 1;
				enc3 = ((chr2 & 15) << 2) | ((chr3 & 0xc0) >> 6);
				enc4 = chr3 & 63;
			}else{
				enc3 = (chr2 & 15) << 2;
				enc4 = 64;
			}
		}else{
			enc2 = (chr1 & 3) << 4;
			enc3 = 64; enc4 = 64;
		}
		sb.WriteString(string(_keyStr[enc1]));
		sb.WriteString(string(_keyStr[enc2]));
		sb.WriteString(string(_keyStr[enc3]));
		sb.WriteString(string(_keyStr[enc4]));
	}
	return sb.String();
}

/*
go-base64-demo-1.go
程序中书写着所见所闻所感，编译着心中的万水千山。
Z28tYmFzZTY0LWRlbW8tMS5nbwrnqIvluo/kuK3kuablhpnnnYDmiYDop4HmiYDpl7vmiYDmhJ/vvIznvJbor5HnnYDlv4PkuK3nmoTkuIfmsLTljYPlsbHjgII=
92 go-base64-demo-1.go
程序中书写着所见所闻所感，编译着心中的万水千山。
*/