package isogram

import (
  "strings"
  "unicode"
)

func IsIsogram(word string) bool {

  lettersChecked := make(map[rune]bool)

  for _, letter := range strings.ToLower(word) {
    if lettersChecked[letter] && unicode.IsLetter(letter) {
      return false
    }

    lettersChecked[letter] = true

  }

  return true

}
