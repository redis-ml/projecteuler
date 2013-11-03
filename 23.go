package main

import (
  "fmt"
  "math"
)

func d(i int) int {
  l := int(math.Sqrt(float64(i)))
  sum := 1 
  //fmt.Printf("intput: %d : 1 ", i)
  for j:=2;j<=l;j++ {
    if i % j == 0 {
      n := i / j 
      if n < j { break }
      sum += j
      //fmt.Printf("%d ", j)
      if n != j { sum += n 
      //fmt.Printf("%d ", n)
      }
    }
  }
  //fmt.Println()
  return sum
}

func main() {
  n := 28123 
  rt := make([]int, n )
  for i,_ := range rt { rt[i] = 0 }
  for i:= 1; i<n;i++ {
    rt[i] = d(i)
  }

  mark := make([]bool, n)
  for i,_ := range mark { mark[i] = false }

  for i:=12;i<n;i++ {
    if rt[i] <= i { continue }
    for j:=i;j<n;j++ {
      if rt[j] <= j { continue }
      sum := i + j
      if sum < len(mark) {
        mark[sum] = true
      }
    }
  }

  sum := 0
  for i,j := range mark{
    if ! j {
      sum += i
    }
  }
  //for i,j := range rt{
    //fmt.Printf( " d(%d)=%d",i,j )
    //if i % 15 == 0 { fmt.Println() }
  //}
  fmt.Printf( "\n result : %d", sum)
}

