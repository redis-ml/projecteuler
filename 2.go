package main

import (
  "flag"
  "fmt"
)

var data = flag.Uint64("data", uint64(0), "input data")

func main() {
  flag.Parse()

  fmt.Printf("data %d\n", *data)
  sum := uint64(0)

  l := uint64(1)
  r := uint64(1)
  for {
    n := l + r
    if n > *data {
      break
    }
    if n % 2 == 0 {
      fmt.Println("sum += %d", n)
      sum += n
    }
    l = r
    r = n
  }
  fmt.Println("data: %d result: %d", *data, sum)

}
