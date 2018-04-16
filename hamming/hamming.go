package hamming

import (
  "errors"
)

// hamming distance is the number of basepairs differing in exact sequence between 2 sequences of homologous DNA
func Distance(a, b string) (int, error) {

  // quick check to see if they are equal
  if a == b {
    return 0, nil
  }

  // lengths should be the same, otherwise return the difference
  if len(a) != len(b) {

    basepairLengthDifference := len(a) - len(b)
    if basepairLengthDifference > 0 {
      basepairLengthDifference *= -1
    }
    return basepairLengthDifference, errors.New("sequences are of different lengths")
  }

  // hamming algo
  mismatches := 0
  for ii := 0; ii < len(a); ii++ {

    if a[ii] != b[ii] {
      mismatches++
    }
  }

  return mismatches, nil
}
