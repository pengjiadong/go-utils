package str

import "encoding/base64"

func Base64Encode(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}

func Base64Decode(in string) string {
	val, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return ""
	}
	return string(val)
}

func Base64EncodeByte(in []byte) string {
	return base64.StdEncoding.EncodeToString(in)
}
