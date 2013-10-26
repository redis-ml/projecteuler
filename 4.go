package main

import (
  "flag"
  _ "os"
  "fmt"
  "math"
)

var data = flag.Uint64("data", uint64(0), "input data")

func isPalindromic(d uint) bool {
  n := uint(math.Floor(math.Log10(float64(d))))
  for i:=uint(0);i<=n/2;i++ {
    base := uint(math.Pow10(int(n-2*i)))
    rest := d % base
    l := (d - rest) / base
    r := d % 10
    fmt.Println("input: %d  base: %d rest: %d r: %d", d, base, rest, r)
    if l != r {
      return false
    }
    d = rest / 10
  }
  return true
}

func main(){
  flag.Parse()
  
  ret := uint(0)
  for i:=uint(999);i>=100;i-- {
    for j:=i;j>=100;j-- {
      p := uint(i)*j
      if p < ret {
        break
      }
      fmt.Println(" %d * %d = %d", i, j , p)
      if isPalindromic(p) {
        ret = p
        break
      }
    }
  }
  if ret != 0 {
    fmt.Println("got Result: %d", ret)
  }
}
