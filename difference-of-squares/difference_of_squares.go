package diffsquares

// sum of first n natural numbers followed by squaring
func SquareOfSums(n int) int {
  sum := 0
  for ii := 1; ii <= n; ii++ {
    sum += ii
  }
  return sum*sum
}

// squaring of the first n natural numbers followed by summation
func SumOfSquares(n int) int {
  sum := 0
  for ii := 1; ii <= n; ii++ {
    sum += ii*ii
  }
  return sum
}


// difference between squareofsums and sumofsquares
func Difference(n int) int {
  squareOfSums := SquareOfSums(n)
  sumOfSquares := SumOfSquares(n)
  return squareOfSums - sumOfSquares
}
