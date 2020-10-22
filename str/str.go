package str

import "github.com/axgle/mahonia"

// GBK2UTF8 .
func GBK2UTF8(data []byte) string {
	dec := mahonia.NewDecoder("gbk")
	ret := dec.ConvertString(string(data))
	return ret
}

// UTF82GBK .
func UTF82GBK(data []byte) string {
	encoder := mahonia.NewEncoder("gbk")
	ret := encoder.ConvertString(string(data))
	return ret
}
