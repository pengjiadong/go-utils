package str

import (
	"strconv"
	"strings"

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

// UnescapeUnicodeByte .
// https://www.zhihu.com/question/330544039
func UnescapeUnicodeByte(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

// UnescapeUnicode .
func UnescapeUnicode(raw string) (string, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(raw), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}
