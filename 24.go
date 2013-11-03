package main

import (
  "fmt"
)

func main() {
  n := 10
  od := 1000000

  //n = 3
  //od = 2
  p := make([]int, n)
  init := 1

  for i,_  := range p {
    p[i] = init
    init *= i+2
  }

  dest := od-1
  or := make([]int, n)
  for i,_ := range or {
    or[i] = i
  }
  for i:=n-1;i>0;i-- {
    order := dest / p[i-1]
    markAndPrint(or, order)
    dest %= p[i-1]
  }
  markAndPrint(or, 0)
}

func markAndPrint(in []int, n int) {
  cnt := 0
  for i,j := range in {
    if j < 0 { continue }
    if cnt == n {
      fmt.Printf( "%d", j)
      in[i] = -1
    }
    cnt ++
  }
}

