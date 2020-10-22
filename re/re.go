package re

import (
	"regexp"
)

func MustCompile(p interface{}) *regexp.Regexp {
	switch p.(type) {
	case string:
		return regexp.MustCompile(p.(string))
	case *regexp.Regexp:
		return p.(*regexp.Regexp)
	default:
		panic("error p")
	}
}

// FindString .
func FindString(p interface{}, s string) (result string) {
	r := MustCompile(p)
	result = r.FindString(s)
	return
}

// FindStringSubmatch .
func FindStringSubmatch(p interface{}, s string) MatchResult {
	r := MustCompile(p)
	result := r.FindStringSubmatch(s)
	return MatchResult(result)
}

// FindAllStringSubmatch .
func FindAllStringSubmatch(p interface{}, s string) MatchsResult {
	r := MustCompile(p)
	result := r.FindAllStringSubmatch(s, -1)
	return MatchsResult(result)
}

// FindAllString .
func FindAllString(p interface{}, s string) (result MatchResult) {
	r := MustCompile(p)
	result = r.FindAllString(s, -1)
	return
}

// Find .
func Find(p interface{}, s string, indexs ...int) (m string, result [][]string, ok bool) {
	var i1, i2 int
	if len(indexs) == 2 {
		i1 = indexs[0]
		i2 = indexs[1]
	}
	r := MustCompile(p)
	result = r.FindAllStringSubmatch(s, -1)
	if result != nil {
		ok = true
	}
	if len(result) > i1 && len(result[i2]) > i2 {
		m = result[i1][i2]
	}
	return
}

// Get .
func Get(data [][]string, i1, i2 int) string {
	if data != nil {
		if len(data) > i1 && len(data[i1]) > i2 {
			return data[i1][i2]
		}
	}
	return ""
}
