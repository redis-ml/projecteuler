package main

import (
  "fmt"
)

func main() {
  dest := 200

  c := []int {1, 2, 5, 10, 20, 50, 100, 200}
  max := make([]int, len(c))
  curr := make([]int, len(c))

  for i,_ := range c {
    max[i] = dest / c[i]
    curr[i] = 0
  }

  si := len(c)-1

  cnt := 0
  for {
    sum := 0
    for i,j := range curr {
      sum += c[i] * j
    }
    if sum >= dest {
      if sum == dest {
        cnt ++
        for _,j := range curr {
          fmt.Print(" ", j)
        }
        fmt.Println()
      }
      for i:=si;i>=0&&curr[i] == 0; i-- {
        curr[i] = max[i]
      } 
    }

    // advance
    isOverflow := false
    for j:=si;j>=0;j-- {
      if curr[j] < max[j] {
        curr[j] ++
        break
      } else {
        curr[j] = 0
        if j == 0 {
          isOverflow = true
        }
      }
    }
    if isOverflow {
      break
    }
  }
  
  fmt.Println("result ", cnt)
}
