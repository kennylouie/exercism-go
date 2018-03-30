package raindrops

import (
  "strconv"
  "sort"
)

// convert a number to raindrop speak
func Convert(number int) string {

  // iterate over the map in order
  orderedKeys := make([]int, 0)
  for key, _ := range factorResponsePairs {
    orderedKeys = append(orderedKeys, key)
  }
  sort.Ints(orderedKeys)

  // iterate over the map values
  rainSpeak := ""

  for _, key := range orderedKeys {
    if number % key == 0 {
      rainSpeak += factorResponsePairs[key]
    }
  }

  if len(rainSpeak) == 0 {
    return strconv.Itoa(number)
  }

  return rainSpeak

}

// raindrops factor response pairs
var factorResponsePairs = map[int]string{
  3: "Pling",
  5: "Plang",
  7: "Plong",
}
