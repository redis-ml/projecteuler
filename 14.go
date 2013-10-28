package main

import (
  "fmt"
)

const max = 1000000
var buf []uint64

func main() {
  buf = make([]uint64, max)
  l := uint64(0)
  li := uint64(0)
  v := false
  for i:= uint64(1); i<uint64(max); i++ {
    if buf[i] != 0 {
      continue
    }
    n, err := trace(i,v)
    if v {
      fmt.Println("")
    }
    if err {
      fmt.Println("got error" , i)
      return
    }
    if l < n {
      l = n
      li = i
    }
    buf[int(i)] = n
  }
  fmt.Println(" largest step:", l, ", num:", li)
  fmt.Println()
  trace(li, true)
  fmt.Println()
}

func trace(ii uint64, v bool) (uint64, bool) {
  ret := uint64(0)
  ori := ii
  if v {
    fmt.Print(" ", ii)
  }
  for i:=ii;i != 1;ret ++ {
    if i % 2 != 0 {
      ni := i * 3 + 1
      if ni < i {
        fmt.Println(" error for original:",ori, ", step:" , ret, ", n:", i)
        return 0, true 
      }
      i = ni
    } else {
      i /= 2
    }
    if v {
      fmt.Print(" ", i)
    }
    if !v && i < uint64(max) && buf[int(i)] != 0 {
      return ret + 1 + buf[int(i)], false
    } 
  }
  return ret + 1, false
}
