package main

import (
  "fmt"
)

func p(buf []int, level int) int {
  return pn(buf, len(buf),level)
}
func pn(buf []int, n,level int) int {
  bl := len(buf)
  if level == bl - 1 {
    //print(buf)
    return tryFind(buf, n)
  }

  sum := 0
  rt := pn(buf, n, level+1)
  if rt > 0 && level >= 4 {
    return rt
  }
  sum += rt
  for i:=level+1;i<bl;i++ {
    swap(buf, i, level)
    rt = pn(buf, n, level+1)
    swap(buf, i, level)
    if rt > 0 && level >= 4 {
      return rt
    }
    sum += rt
  }
  return sum
}

func tryFind(buf []int, n int) int {
  n1 := tryOne(buf, n, 1,4)
  n2 := tryOne(buf, n, 2,3)

  return n1+n2
}

func tryOne(buf[]int, n, a1i,a2i int) int {
  a1 := parse(buf, n-a1i, n)
  a2 := parse(buf, n-a1i-a2i, n-a1i)
  a3 := parse(buf, 0, n-a1i-a2i)

  if a1 * a2 == a3 {
    fmt.Println("found ", a1, " * ", a2, " = ", a3)
    return a3
  }

  return 0
}

func parse(buf[]int, l,r int) int {
  sum := 0
  for i:=l;i<r;i++ {
    sum = sum * 10 + buf[i]
  }
  return sum
} 

func swap(buf []int, l int, r int) {
  tmp := buf[l]
  buf[l] = buf[r]
  buf[r] = tmp
}

func print(b []int) {
  for _,j := range b {
    fmt.Print(" ", j)
  }
  fmt.Println()
}

func main() {

  n := 9
  buf := make([]int, n)

  for i,_ := range buf {
    buf[i] = i+1
  }

  ret := p(buf, 0)
  fmt.Println("result ", ret)
}


