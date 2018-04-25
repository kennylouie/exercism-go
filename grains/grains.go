package grains

import (
  "math"
  "errors"
)

// the number of grains on a square is a power of 2
func Square(whichSquare int) (uint64, error) {

  ErrOutOfRange := errors.New("Must pick a square between 1 and 64 (inclusive).")

  if whichSquare > 64 || whichSquare < 1 {
    return 0, ErrOutOfRange
  }

  power := float64(whichSquare - 1)
  grains := math.Pow(2, power)

  return uint64(grains), nil

}

// sum of all number of grains
// 64 squares on a chessboard
func Total() uint64 {

  totalGrains := uint64(0)

  for square := 1; square <= 64; square++ {
    grains, _ := Square(square)
    totalGrains +=  grains
  }

  return totalGrains
  // return 18446744073709551615

}
