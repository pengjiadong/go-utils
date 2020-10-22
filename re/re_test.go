package re

import (
	"testing"
)

func TestFindAllString(t *testing.T) {
	r := FindAllString("(a)(b)", "abcabe")
	t.Log(r)
}

func TestFindAllStringSubmatch(t *testing.T) {
	r := FindAllStringSubmatch("(a)(b)", "abcabe")
	t.Log(r)
}

func TestFindString(t *testing.T) {
	r := FindString("(a)(b)", "abcabe")
	t.Log(r)
}

func TestFindStringSubmatch(t *testing.T) {
	r := FindStringSubmatch("(a)(b)", "abcabe")
	t.Log(r)
}

func TestFind(t *testing.T) {
	r1, r2, ok := Find("(a)(b)", "abcabe", 3, 1)
	t.Log(ok, r1, r2)
}

func TestFindAllStringSubmatch2(t *testing.T) {
	r := FindAllStringSubmatch("(a)1(b)", "abcabe")
	m := r.Get(0, 0)
	t.Log(r, m)
}
