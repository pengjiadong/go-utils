package sys

import "testing"

func TestGetCode(t *testing.T) {
	c := GetCode()
	t.Log(c)
}
