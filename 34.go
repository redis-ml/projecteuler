package main

import (
  "fmt"
)

var buf []int 

func main() {
  n := judgeMax()
  fmt.Print("result maxN:", n)

  sum := 0
  for i:=0;i<n;i++ {
    if tryOne(i) > 0 {
      sum += i
      fmt.Println(" got one ", i)
    }
  }
  fmt.Println("result ", sum)
}

func tryOne(n int) int {
  sum := 0
  num := 0
  for nn:=n;nn>0;nn/=10 {
    num ++
    sum += buf[nn%10]
  }
  if num > 1 && n == sum {
    return n
  } else {
    return 0
  }
}
func judgeMax() int {
  ret := 1
  buf = make([]int,10)
  base := 1
  for i,_ := range buf {
    if i == 0 {
      buf[i] = 1
    } else {
      base *= i
      buf[i] = base
    }
  }

  power:=ret
  pp := buf[9]
  for ; power < pp; ret ++ {
    power *=10
    pp += buf[9] 
  }
  return power
}
