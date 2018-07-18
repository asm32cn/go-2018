// go-carsar-demo-1.go
package main;

func main(){
	var strSource = "JavaCarsarDemo1.java (JAVA实现caesar凯撒加密算法)";
	var strEncrypt = CarsarEncode(strSource);
	var strDecrypt = CarsarDecode(strEncrypt);

	println("strSource:\n\t" + strSource);
	println("strEncrypt:\n\t" + strEncrypt);
	println("strDecrypt:\n\t" + strDecrypt);
}

func CarsarEncode(s string) string{
	var n = len(s);
	var data = make([]byte, n);
	for i := 0; i < n; i++ {
		var ch = int(s[i]);
		if(ch >= 65 && ch <= 90){
			ch = 65 + (ch - 65 + 3) % 26;
		}else if(ch >= 97 && ch <= 122){
			ch = 97 + (ch - 97 + 3) % 26;
		}
		data[i] = byte(ch);
	}
	return string(data);
}

func CarsarDecode(s string) string{
	var n = len(s);
	var data = make([]byte, n);
	for i := 0; i < n; i++ {
		var ch = int(s[i]);
		if(ch >= 65 && ch <= 90){
			ch = 65 + (ch - 65 - 3) % 26;
		}else if(ch >= 97 && ch <= 122){
			ch = 97 + (ch - 97 - 3) % 26;
		}
		data[i] = byte(ch);
	}
	return string(data);
}