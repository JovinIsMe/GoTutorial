package main

import (
  "fmt"
  "math"
  "math/rand"
)

func randIntn(n int) int {
  return rand.Intn(n)
}

func sqrt(n int) float64 {
  float64Number := float64(n)
  return math.Sqrt(float64Number)
}

func main() {
  // var floatNumber float64 = 64
  // or
  // var floatNumber float64
  // floatNumber = 64

  // var intNumber = 64
  // or
  intNumber := 64

  fmt.Println("math.Intn: ", randIntn(intNumber))

  // Formatting with fmt package
  fmt.Println("math.Sqrt: ", sqrt(intNumber))
  fmt.Printf("math.Sqrt: %.0f\n", sqrt(intNumber))
  fmt.Printf("math.Sqrt: %v\n", sqrt(intNumber))

  fmt.Println("Pi:", math.Pi)
}
