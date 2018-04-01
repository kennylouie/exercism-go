package luhn

import (
  "strings"
  "errors"
)

// luhn algorithm for checking validity of strings
func Valid(aString string) bool {

  // strip white space
  aString = strings.Map(isSpace, aString)

  lengthString := len(aString)

  if lengthString > 1 {

    // check for non-numbers
    numbers, err := isAllNumbers(aString)
    if err != nil {
      return false
    }

    // double every other number starting from right
    numbers = doubleEverySecondFromRight(numbers)

    sum := 0
    for _, number := range numbers {
      sum += number
    }

    if sum % 10 == 0 {
      return true
    }

  }

  return false

}

// checking white space chars
func isSpace(char rune) rune {
  if char == ' ' {
    return -1
  }
  return char
}

// check if character is a number
func isNumber(char rune) (int, error) {
  var ErrNotNum = errors.New("not a number")

  if char >= '0' && char <= '9' {
    return int(char)-'0', nil
  }

  return 0, ErrNotNum
}

// check string if it is all numbers; will return all the numbers as a slice
func isAllNumbers(aString string) ([]int, error) {

  var ErrHasNonNum = errors.New("string contains non-number characters")

  var numbers = make([]int, 0, len(aString))

  for _, character := range aString {
    number, err := isNumber(character)
    if err != nil {
      return numbers, ErrHasNonNum
    }
    numbers = append(numbers, number)
  }
  return numbers, nil

}

// double every second element of an array of numbers starting from the right
// if product is greater than 9, subtract 9
func doubleEverySecondFromRight(numbers []int) []int {

  lengthNumbers := len(numbers)

  for index := 2; index <= lengthNumbers; index += 2 {
    doubled := numbers[lengthNumbers-index]*2
    if doubled > 9 {
      doubled = doubled - 9
    }
    numbers[lengthNumbers-index] = doubled
  }

  return numbers
}
