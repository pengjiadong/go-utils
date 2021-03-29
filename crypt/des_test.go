package crypt

import "testing"

func TestEncrypt(t *testing.T) {
	d := DesCBC{
		IV:  []byte("12345678"),
		Key: []byte("test1234"),
	}
	r, err := d.Encrypt("abc")
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func TestDe(t *testing.T) {
	d := DesCBC{
		IV:  []byte("12345678"),
		Key: []byte("test1234"),
	}
	r, err := d.Decrypt("UvUjUlR/0WI=")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(r))
}
