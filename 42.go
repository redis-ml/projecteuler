package main

import (
  "flag"
  "fmt"
  "os"
  "bufio"
  "math"
)

var input = flag.String("input", "", "input data file")

func main() {
  flag.Parse()

  f, err := os.Open(*input)
  if err != nil {
    fmt.Println("failed to open input file ", *input)
    return
  }
  defer f.Close()

  found := 0
  s := bufio.NewScanner(f)
  for s.Scan() {
    l := s.Text()
    if len(l) == 0 { continue }
    sum := 0
    for _,j := range l {
      n := int(j-'A'+1)
      sum += n
    }
    num := sum * 2
    n := int(math.Sqrt(float64(num)))
    if n*(n+1) == num {
      found ++
      fmt.Println("found one ", l, " : ", (num/2) , " = ", n , "*", (n+1), "/2")
    }
  }
  fmt.Println("result ", found)
}
