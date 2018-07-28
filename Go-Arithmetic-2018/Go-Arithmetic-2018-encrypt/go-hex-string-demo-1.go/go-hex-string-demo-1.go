// go-hex-string-demo-1.go
package main;

const strHexDigit = "0123456789ABCDEF";

func main(){
	var strSource = "go-base64-demo-1.go\n程序中书写着所见所闻所感，编译着心中的万水千山。";
	var strEncrypt = HexStringEncode(strSource);
	var strDecrypt = HexStringDecode(strEncrypt);
	println("strSource:\n\t" + strSource);
	println("strEncrypt:\n\t" + strEncrypt);
	println("strDecrypt:\n\t" + strDecrypt);
}

func HexStringEncode(s string) string{
	var _map = make(map[int]byte, 16);
	for i := 0; i < 16; i++ { _map[i] = strHexDigit[i]; };
	var nCount = len(s);
	var cache = make([]byte, nCount * 2);
	for i := 0; i < nCount; i++ {
		var n = i + i;
		var nCode = int(s[i]);
		cache[n] = _map[ (nCode & 0xf0) >> 4 ];
		cache[n + 1] = _map[ nCode & 15 ];
	}
	return string(cache);
}

func HexStringDecode(s string) string{
	var _map = make(map[byte]int, 16);
	for i := 0; i < 16; i++ { _map[strHexDigit[i]] = i; };
	var nCount = len(s) / 2;
	var cache = make([]byte, nCount);
	for i := 0; i < nCount; i++ {
		var n = i + i;
		var nCode = _map[ s[n] ] << 4 + _map[ s[n + 1] ];
		cache[i] = byte(nCode);
	}
	return string(cache);
}