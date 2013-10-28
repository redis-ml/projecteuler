package main

import (
  "flag"
  _ "os"
  "fmt"
  _ "math"
)

var data = flag.Uint64("data", uint64(0), "input data")
var size = flag.Uint64("size", uint64(0), "input data")

func main(){
  flag.Parse()
  
  fmt.Printf(" find para %d %d", *data, *size)
  ls := NewLs()
  cnt := uint64(1)
  step := uint64(1) << 10
  marked := uint64(0)
  for i:=uint64(3); i<=*size; i+=2 {
    if i > marked {
      marked += step
      ls.Mark(marked)
    }
    if ls.Get(i) {
      continue
    }
    cnt ++
    //fmt.Printf("data %d found %d cnt %d\n", *data, i, cnt)
    if cnt == *data {
      fmt.Printf("data %d result %d cnt %d\n", *data, i,cnt)
      return
    }
  }
  fmt.Printf("not found %d", cnt)
}

