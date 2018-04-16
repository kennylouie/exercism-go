package reverse

func String(toBeReversed string) string {

  // quick check for empty string
  if toBeReversed == "" {
    return ""
  }

  toBeReversedRune := []rune(toBeReversed)

  lengthString := len(toBeReversedRune)
  reverseString := ""
  for ii := 1; ii <= lengthString; ii++ {
    reverseString += string(toBeReversedRune[lengthString-ii])
  }

  return reverseString

}
