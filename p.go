package main

import (
  "fmt"
)

func main() {

  n := 4

  buf := make([]int, n)
  cur := make([]int,n)

  for i,_ := range buf {
    buf[i] = i
    cur[i] = 0
  }

  for {
    print(buf)

    isOverflow := false
    for j:=n-1;j>=0;j-- {
      if cur[j] != 0 {
        tmp := buf[j]
        buf[j] = buf[j+cur[j]]
        buf[j+cur[j]] = tmp
      }
      if cur[j] + 1 > n-1-j {
        if j == 0 {
          isOverflow = true
        } else {
          continue
        }
      } else {
        cur[j] ++
        tmp := buf[j]
        buf[j] = buf[j+cur[j]]
        buf[j+cur[j]] = tmp
        break
      }
    }
    if isOverflow {
      break
    }
  }
}

func print(b []int) {
  for _,j := range b {
    fmt.Print(" ", j)
  }
  fmt.Println()
}


