package main

import (
  "fmt"
)

func main() {

  n := 987654321
  ls := NewLs()
  ls.Mark(uint64(n+1))

  b := make([]int, 9)
  for i,_ := range b { b[i] = 9-i }

  ret := 0
  for start:=0;start < 9;start ++ {
    if start > 0 { b[start - 1] = 0 }
    ret = p(ls, b, start, 9-1)
    if ret != 0 { break }
  }
  fmt.Println("got result ", ret)
}

func p(ls *LS, b []int, start,end int) int {
  if start == end {
    num := 0
    for _,j := range b {
      num = num * 10 + j
    }
    fmt.Println("got ", num)
    if ! ls.Get(uint64(num)) {
      return num
    }
    return 0
  }

  t := p(ls, b, start + 1, end)
  if t != 0 { return t }

  for i:=start+1;i<=end;i++ {
    tmp := b[start]
    b[start] = b[i]
    b[i] = tmp
    t = p(ls, b, start + 1, end)
    b[i] = b[start]
    b[start] = tmp

    if t > 0 { return t }
  }
  return 0
}

