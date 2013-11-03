package main

import (
  "fmt"
  "math"
)

func main() {
  max := 1000

  buf := make([]int, max*10)

  longest := 0
  lnum := 0
  for i:=max-1;i>1;i-- {
     for j,_:=range buf { buf[j] = 0 }

     s := 0 
     
     len := int(math.Log10(float64(i)))
     top := int(math.Pow10(len+1))

     for buf[top] == 0 {
        buf[top] = s
        s ++
        top = (top % i)*10
     }
     ns := s - buf[top]
     if ns > longest {
       longest = ns
       lnum = i
     }
     fmt.Println(" cyclt ", i, " = ", ns)
  } 
  fmt.Println("longst cycle :" , lnum, " with " , longest)
}

