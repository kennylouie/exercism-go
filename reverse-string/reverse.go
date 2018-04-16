package reverse

func String(toBeReversed string) string {

  // quick check for empty string
  if toBeReversed == "" {
    return ""
  }

  toBeReversedRune := []rune(toBeReversed)

  lengthString := len(toBeReversedRune)

  for head, tail := 0, lengthString-1; head < tail; head, tail := head + 1, tail - 1 {
    tmp := toBeReversedRune[head]
    toBeReversedRune[head] = toBeReversedRune[tail]
    toBeReversedRune[tail] = tmp
  }

  return string(toBeReversedRune)

}
