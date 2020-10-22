package crypt

import (
	"strings"
	"testing"
)

func TestMD5(t *testing.T) {
	x := MD5("666")
	t.Log(x)
	t.Log(strings.ToUpper(x))
}
