package mongox

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func IsDuplicateError(err error) bool {
	var exception mongo.WriteException
	if errors.As(err, &exception) {
		for _, e := range exception.WriteErrors {
			if e.Code == 11000 {
				return true
			}
		}
	}
	return false
}
