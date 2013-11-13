package main

import (
  "fmt"
)

func main() {

  numerator := 1
  denominator := 1
  
  for i:=11;i<100;i++{
    if i%10 == 0 { continue }
    li := i % 10
    hi := i / 10
    for j:=i+1;j<100;j++ {
      if j%10 == 0 { continue }
      lj := j % 10
      hj := j / 10
      nc := 0
      np := 0
      nt := 0
      if li == lj {
        nt = li
        nc = hi
        np = hj
      } else if li == hj {
        nt = li
        nc = hi
        np = lj 
      } else if hi == lj {
        nt = hi
        nc = li
        np = hj
      } else if hi == hj {
        nt = hi
        nc = li
        np = lj
      } else {
        continue
      }
      if nt == 0 {continue}
      // i/j == nc /np
      if i * np == nc * j {
        fmt.Println("", i, "/", j, "==", nc, "/", np)
        numerator *= nc
        denominator *= np
      }
    }
  }
  c := common(numerator, denominator)
  ret := denominator / c

  fmt.Println("result ", ret)
}

func common(n,d int) int {
  l := n
  h := d
  for {
    if l > h {
      l -= h
    } else {
      h -= l
    }
    if l == 0 {
      return h
    } else if h == 0 {
      return l
    }
  }
}
