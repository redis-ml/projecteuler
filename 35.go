package main

import (
  "fmt"
  "math"
)

func main() {
  n := 1000000
  buf := make([]bool, n)

  max := int(math.Sqrt(float64(n)))
  for i,_ := range buf {
    buf[i] = false
  }

  for i:=2;i<=max;i++ {
    if buf[i] { continue }
    for j:=i*i;j<n;j+=i {
      buf[j] = true
    }
  }

  ret := 0
  num := 0
  for i:=2;i<n;i++ {
    if buf[i] {continue}
    l := int(math.Log10(float64(i)))
    b := int(math.Pow10(l))
    nn := i

    j:=0
    for ;j<l;j++ {
      d := nn / b
      nn = (nn%b)*10 + d
      if buf[nn] {
        break
      }
    }
    isOk := j == l
    if isOk {
      fmt.Println("got one ", i)
    }

    tmpSum := 0
    tmpN := 0 
    for j=0;j<=l;j++ {
      if ! buf[nn] {
        buf[nn] = true
        tmpN ++
        tmpSum += nn
      }
      d := nn / b
      nn = (nn%b)*10 + d
    }
    if isOk {
      ret += tmpSum
      num += tmpN
    }
  }
  fmt.Println("result ", num, " sum ", ret)
}
