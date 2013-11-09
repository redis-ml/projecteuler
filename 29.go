package main

import (
  "fmt"
  "math"
)

func main() {

  a1 := 2
  a2 := 100 
  b1 := 2
  b2 := 100 

  abuf := make([]int, a2+1)
  for i,_ := range abuf { abuf[i] = 0 }

  cnt := 0

  for a:=a1;a<=a2;a++ {
    max := int(math.Log(float64(a2))/math.Log(float64(a)))

    tbuf := make([]int,max*b2 + 1)
    for i,_ := range tbuf { tbuf[i] = 0 }

    for i:=1;i<=max;i++ {
      nn := int(math.Pow(float64(a), float64(i)))
      jstart := b1
      if abuf[nn] > 0 {
        jstart = abuf[nn] + 1
      }
      for j:=jstart;j<=b2;j++ {
        if tbuf[i*j] == 0 {
          tbuf[i*j] = 1
          abuf[nn] = j
          cnt ++
          //fmt.Println("a=", nn, " b=", j, " a^b=", int(math.Pow(float64(nn), float64(j))))
        } else {
          if j < b2 {
            abuf[nn] = j + 1
          }
        }
      }
    }
  }

  fmt.Println("result ", cnt)
}
