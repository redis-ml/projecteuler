package main

import (
  "fmt"
)

var a [][]uint

func main() {
  col := 20 + 1
  row := 20 +1
  a = make([][]uint, row)
  for i, _ := range a {
    a[i] = make([]uint, col)
  }
  a[row-1][col-1] = 1

  h := col + row -2 
  for i:= h-1;i>=0; i-- {
    for r := 0;r<row; r++{
      c := i - r
      if c<0 || c >= col {
        continue
      }
      left := uint(0)
      if r < row - 1 {
        left = a[r+1][c]
      }
      right := uint(0)
      if c < col - 1 {
        right = a[r][c+1]
      }
      a[r][c] = left +right
    }
  }

  for _, ar := range a {
    for _,t := range ar {
      fmt.Printf(" %04d", t)
    }
    fmt.Println()
  }
}
