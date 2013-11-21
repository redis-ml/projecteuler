package main

import (
  "math"
  "fmt"
)

func main() {

  for i:= 285+1;true;i++ {
    t := T(i)
    if solveP(t) > 0 &&solveH(t) > 0 {
      fmt.Println("result ", t)
      return
    }
  }
}

func T(n int) int {
  return n*(n+1)/2
}

func solveP(k int) int {
  n := (1+int(math.Sqrt(float64(1+24*k))))/6
  if k == n*(3*n-1)/2 {
    return n
  }
  return 0
}

func solveH(k int) int {
  n := (1+int(math.Sqrt(float64(1+8*k))))/4
  if k == n*(2*n-1) {
    return n
  }
  return 0
}

