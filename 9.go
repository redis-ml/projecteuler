package main

import (
  "flag"
  _ "os"
  "fmt"
  _ "math"
)

var data = flag.Uint64("data", uint64(0), "input data")

func main(){
  flag.Parse()
  
  for i:=0; i<int(*data);i++ {
    j := 1000 - 500000/(1000-i)
    k := 1000 - i - j
    if i * i +j * j == k * k  {
      fmt.Println("fount ", i, j, k)
    }
  }
}

