package main 

import (
  "fmt"
  "math"
)

func f(a,b,n int) int {
  return n*n+a*n+b
}

func main() {
  test := 1000000
  fib := make([]bool, test)
  for i,_ := range fib {
    fib[i] = false
  }

  fib[0] = true
  fib[1] = true
  lmt := int(math.Sqrt(float64(test)))
  for i:=2;i<lmt;i++ {
    if fib[i] { continue }
    for j:=i*i;j<test;j+=i {
      fib[j] = true
    }
  }

  largest := 0
  la := 0
  lb := 0
  for a := -999; a < 1000; a ++ {
    for b:=-999; b < 1000; b+=2 {
      n := 0
      for {
        rt := f(a,b,n)
        if rt < 0 { rt = rt * -1 }
        if rt >= test {
          fmt.Println("we can't decide ", a, b, n, rt)
          break
        } else {
          if fib[rt] {
            if n > largest {
             largest = n
             la = a
             lb = b
            }
            break
          }
        }
        n ++
      }
    }
  }
  fmt.Println("sure result : ", la, lb, largest)
}

