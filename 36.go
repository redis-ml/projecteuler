package main

import (
  "math"
  "fmt"
)

func main() {
  n := 1000000
  ret := 0
  for i:=1;i<n;i++ {
    if !isPalin2(i) || ! isPalin10(i) { continue }
    ret += i
  }

  fmt.Println("585 test ", isPalin10(585), " 2base ", isPalin2(585) )
  fmt.Println("ret ", ret)
}

func isPalin2(num int)  bool {
  n := int(math.Log2(float64(num)))
  lm := 1 << uint(n)
  rm := 1

  for lm>rm {
    if (num&lm == 0) != (num&rm == 0) { return false }
    lm >>= 1
    rm <<= 1
  }
  return true
}


func isPalin10(num int)  bool {
  n := int(math.Log10(float64(num)))
  lm := int(math.Pow10(n))
  rm := 1
  for lm > rm {
    if (num/lm)%10 != (num/rm)%10 { 
      fmt.Println(" n=", n, " (n/10)%10=", (n/10)%10, " (n/rm)%10=", (n/rm)%10)
      return false }
    lm /= 10
    rm *= 10
  }
  return true
}

