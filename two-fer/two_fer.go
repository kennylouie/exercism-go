package twofer

import (
  "fmt"
)

// function to produce "One for X, one for me.", where X is a name or "you"
func ShareWith(person string) string {

  if person == "" {
    person = "you"
  }

	return fmt.Sprintf("One for %s, one for me.", person)
}
