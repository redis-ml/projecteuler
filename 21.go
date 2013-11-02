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
  n := 10000
  rt := make([]int, n )
  for i,_ := range rt { rt[i] = 0 }
  for i:= 1; i<n;i++ {
    rt[i] = d(i)
  }
  sum := 0
  for i:= 1; i<n; i++ {
    if rt[i] < n && rt[i] != i && rt[rt[i]] == i {
      sum += i
      fmt.Printf("rt[%d]=%d and rt[%d]=%d\n", i, rt[i], rt[i], rt[rt[i]])
    }
  }
  //for i,j := range rt{
    //fmt.Printf( " d(%d)=%d",i,j )
    //if i % 15 == 0 { fmt.Println() }
  //}
  fmt.Printf( "\n result : %d", sum)
}

