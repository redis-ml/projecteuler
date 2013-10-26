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
  
  ret := *data
  for i:=*data-1;i>1;i-- {
    tmp := ret
    for  ;tmp % i != 0 ; tmp += ret {
    }
    ret = tmp
  }
  if ret != 0 {
    fmt.Println("got Result: %d", ret)
  }
}
