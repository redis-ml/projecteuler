package main

import (
  _ "fmt"
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

func NewLs() *LS {
  in := make([][]uint64, 0)
  return &LS{0, in}
}

func (ls *LS) Get(i uint64) bool {
  if i % uint64(2) == 0 {
    return i != uint64(2)
  }
  x := x(i)
  //fmt.Println("Get ", i, " x=" , x, " len(ls.in)=", len(ls.in))
  if len(ls.in[x]) == 0 {
    return false
  }
  y := y(i)
  z := z(i)
  return (ls.in[x][y] & (uint64(1) << z)) != 0
}
func (ls *LS) Set(i uint64) {
  if i % 2 == 0 {
    return
  }
  x := x(i)
  if len(ls.in[x]) == 0 {
    ls.in[x] = make([]uint64, block)
  }
  y := y(i)
  z := z(i)
  ls.in[x][y] |= (uint64(1) << z)
}

func (ls *LS) Mark(n uint64) {
  if ls.cap < n {
    // extend matrix if necessary
    nr := x(n)
    or := uint64(len(ls.in))
    if nr >= or {
      newbuf := make([][]uint64, nr+1)
      for i, a := range ls.in {
        newbuf[i] = a
      }
      ls.in = newbuf
    }

    for i := uint64(3); i < ls.cap; i+=2 {
      if ls.Get(i) {
        continue
      }
      for j := i * i; j < n;  j += 2 * i {
        if j < ls.cap {
          tmp := ls.cap - j 
          delta := tmp % (2*i)
          j = ls.cap - delta
          continue
        }
        ls.Set(i)
      }
    }
    ls.cap = n
  }
  
}

