package re

// MatchResult .
type MatchResult []string

// Get .
func (r MatchResult) Get(i int) string {
	if r != nil {
		if len(r) > i {
			return r[i]
		}
	}
	return ""
}

// Ok .
func (r MatchResult) Ok() bool {
	return r != nil
}

// MatchsResult .
type MatchsResult [][]string

// Get .
func (r MatchsResult) Get(i1, i2 int) string {
	if r != nil {
		if len(r) > i1 && len(r[i1]) > i2 {
			return r[i1][i2]
		}
	}
	return ""
}

// Ok .
func (r MatchsResult) Ok() bool {
	return r != nil
}
