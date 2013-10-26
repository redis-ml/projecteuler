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
  
  l := uint64(0)
  r := uint64(0)
  for i:=*data; i>=1;i-- {
    l += i * i
    r += i
  }
  ret := r*r - l
  if ret != 0 {
    fmt.Println("got Result: %d", ret)
  }
}
