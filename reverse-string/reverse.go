package reverse

func String(toBeReversed string) string {

  toBeReversedRune := []rune(toBeReversed)

  lengthString := len(toBeReversedRune)
  reverseString := ""
  for ii := 1; ii <= lengthString; ii++ {
    reverseString += string(toBeReversedRune[lengthString-ii])
  }

  return reverseString

}
