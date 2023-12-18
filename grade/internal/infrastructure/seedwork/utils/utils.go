package utils

import (
	"fmt"
	"regexp"
)

var reQuestionPlaceholder = regexp.MustCompile(`'(?:[^']|'')*'|"(?:[^"]|"")*"|(\?)`)
var reDollarPlaceholder = regexp.MustCompile(`'(?:[^']|'')*'|"(?:[^"]|"")*"|(\$\d+)`)

func Rebind(query string) string {
	// See also another example of implementation
	// https://github.com/jmoiron/sqlx/blob/1abdd3dc2a5d5257e3714cee6e8e98835cdb1b2e/bind.go#L60
	var offset = 0
	return reQuestionPlaceholder.ReplaceAllStringFunc(query, func(s string) string {
		if s == "?" {
			offset += 1
			return fmt.Sprintf("$%d", offset)
		}
		return s
	})
}

func RebindReverse(query string) string {
	return reDollarPlaceholder.ReplaceAllStringFunc(query, func(s string) string {
		if s[:1] == "$" {
			return "?"
		}
		return s
	})
}
