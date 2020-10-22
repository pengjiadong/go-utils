package str

import "testing"

func TestUTF82GBK(t *testing.T) {
	word := "你好"
	x := UTF82GBK([]byte(word))
	y := GBK2UTF8([]byte(x))
	if y != word {
		t.Error(y, word)
	}
}

func TestGBK2UTF8(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"",
			args{
				data: []byte{196, 227, 186, 195},
			},
			"你好",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GBK2UTF8(tt.args.data); got != tt.want {
				t.Errorf("GBK2UTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}
