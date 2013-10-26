package main

import (
  "flag"
  _ "os"
  "fmt"
  _ "math"
)

const block = uint64(1)<<25

type LS struct {
  cap uint64
  in [][]uint64
}
func x(n uint64) uint64 {
  return n / 2 / 64 / block
}
func y(n uint64) uint64 {
  return (n / 2 / 64) % block
}
func z(n uint64) uint64  {
  return (n/2) % 64
}

func NewLs(cap uint64) *LS {
  r := x(cap-1)+1
  in := make([][]uint64, r)
  return &LS{cap, in}
}

func (ls *LS) Get(i uint64) bool {
  x := x(i)
  if len(ls.in[x]) == 0 {
    return false
  }
  y := y(i)
  z := z(i)
  return (ls.in[x][y] & (uint64(1) << z)) != 0
}
func (ls *LS) Set(i uint64) {
  x := x(i)
  if len(ls.in[x]) == 0 {
    ls.in[x] = make([]uint64, block)
  }
  y := y(i)
  z := z(i)
  ls.in[x][y] |= (uint64(1) << z)
}

var data = flag.Uint64("data", uint64(0), "input data")

func main(){
  flag.Parse()
  
  fmt.Printf(" find para %d ", *data )
  ls := NewLs(*data + 1)
  sum := uint64(2)
  for i:=uint64(3); i<*data; i+=2 {
    if ls.Get(i) {
      continue
    }
    sum += i
    for j:=i*i;j<*data; j+=2*i {
      ls.Set(j)
    }
  }
  fmt.Printf("not found %d", sum)
}

