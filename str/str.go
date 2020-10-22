package str

import (
	"github.com/axgle/mahonia"
)

// GBK2UTF8 .
func GBK2UTF8(in []byte) string {
	dec := mahonia.NewDecoder("gbk")
	ret := dec.ConvertString(string(in))
	return ret
}

// UTF82GBK .
func UTF82GBK(in []byte) string {
	encoder := mahonia.NewEncoder("gbk")
	ret := encoder.ConvertString(string(in))
	return ret
}
