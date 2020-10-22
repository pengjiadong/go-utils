package log

import "testing"

func Test_Init(t *testing.T) {
	opts := NewOption()
	Init(opts)
}
