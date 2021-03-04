package gormx

import (
	"strings"
)

func IsUniqueError(err error) bool {
	if err != nil {
		ret := strings.HasPrefix(err.Error(), "UNIQUE constraint failed")
		return ret
	}
	return false
}
