package main

import (
  "fmt"
  "math"
)

func main() {
  n := 10
  c := math.Sqrt(float64(5))
  a := (float64(1)+c)/float64(2)
  b := (float64(1)-c)/float64(2)

  for i:=0; i<=n;i++ {
    rt := (math.Pow(a, float64(i)) - math.Pow(b,float64(i)))/c
    fmt.Printf("F(%d) = %d\n", i, int(rt))
  }
}
