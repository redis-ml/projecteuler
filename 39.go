package main


import (
  "math"
  "fmt"
)

func main() {
  max := 1000
  rec := make([]int, max+1)
  for i:=1;i<=max/3;i++ {
    i2 := i * i
    for j:=i+1;j<=(max-i)/2;j++ {
      j2 := j * j
      k := int(math.Sqrt(float64(i2+j2)))
      peri := i + j + k
      if k < j || peri > max { break }
      if i2 + j2 == k * k {
        rec[peri] ++
      }
    }
  }

  ret := 0
  l := rec[ret]
  for i,j := range rec {
    if l < j {
      ret = i
      l = j
    }
  }

  fmt.Println(" result is ", ret, " with solutions of " , l) 
}
