package main

import (
  "fmt"
  "math"
)

func main() {
  pb := make([]int, 10)
  for i,_ := range pb {
    pb[i] = int(math.Pow(float64(i), float64(5)))
  }

  up := 10000000
  total := 0

  for i:=2; i<up;i++ {
     sum := 0
     ti := i
     for ti > 0 {
       n := ti % 10
       ti /= 10
       sum += pb[n]
     }
     if sum == i {
       fmt.Println("found ", i)
       total += sum
     }
  }
  fmt.Println("result :", total)
}
