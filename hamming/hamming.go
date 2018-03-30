package hamming

import (
  "errors"
)

// hamming distance is the number of basepairs differing in exact sequence between 2 sequences of homologous DNA
func Distance(a, b string) (int, error) {

  // sequence lengths should be the same, else, return the
  if len(a) != len(b) {

    basepairLengthDifference := len(a) - len(b)
    if basepairLengthDifference > 0 {
      basepairLengthDifference = -basepairLengthDifference
    }
    return basepairLengthDifference, errors.New("sequences are of different lengths")
  }

  // finding the number of different basepairs
  if a == b {
    return 0, nil
  }

  mismatches := 0
  for ii := 0; ii < len(a); ii++ {

    if a[ii] != b[ii] {
      mismatches++
    }
  }

  return mismatches, nil
}
