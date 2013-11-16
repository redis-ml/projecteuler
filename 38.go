package main

import (
  "fmt"
)

func main() {
  b := make([]bool, 9)
  max  := 9999

  ret := uint64(0)
  for i:=1;i<=max;i++ {
    n := tryOne(i, b)
    if n > 0 { fmt.Println(" found one :", n) }
    if n > ret { ret = n }
  }
  fmt.Println(" ret ", ret )

}

func clear(b []bool) {
  for i,_ := range b { b[i] = false }
}

func tryOne(input int, b []bool) uint64 {
  j := uint64(1)
  i := uint64(input)
  ten := uint64(10)
  clear(b) 
  wc := 0
  rt := uint64(0)
  for {
    p := i * j

    for tmp:=p;tmp>uint64(0);tmp/=ten {
      wc ++
      rt *= ten
      d := int(tmp % ten)
      if d == 0 || b[d-1] { return 0 }
      b[d-1] = true
    }
    rt += p
    if wc == len(b) { return rt }
    j ++
  }
}

