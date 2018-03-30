package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
  stringSplits := strings.FieldsFunc(s, delim)
  acronym := ""
  for _, word := range stringSplits {
    acronym += strings.ToUpper(string(word[0]))
  }
  return acronym
}

func delim(r rune) bool {
  return r == ' ' || r == '-'
}
