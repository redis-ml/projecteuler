package main

import (
  _ "math"
  "fmt"
)

func main() {
  n := 7

  start := 1
  dest := 1
  nbase := 1
  ret := 1
  i := 0
  for {
    ns := (i+1)*9*nbase

    for start + ns > dest {
      delta := dest - start
      d1 := delta / (i+1)
      pos := i-delta % (i+1)
      number := nbase + d1
      for ;pos>0;number /= 10 { pos-- }
      got := number % 10

      fmt.Println(" the ", dest, "-th number is ", got)
      ret *= got

      dest *= 10
      n --
      if n==0 { break }
    }
    if n == 0 { break }
    start += ns
    nbase *= 10
    i ++
  }

  fmt.Println("result is ", ret)
}
