package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 .
func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
