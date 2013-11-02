package main

import (
  "os"
  "bufio"
  "fmt"
  "flag"
)


var i = flag.String("input", "", "input file")

func main() {
  flag.Parse()

  f, err := os.Open(*i)
  if err != nil {
    fmt.Println("failed to open input ", *i)
    return
  }
  defer f.Close()

  s := bufio.NewScanner(f)
  rank := uint64(0)
  sum := uint64(0)
  for s.Scan() {
    rank ++
    line := s.Text()
    ts := uint64(0)
    for _,c := range line {
      ts += uint64(c-'A'+1)
    }
    sum += rank * ts
  }

  fmt.Println("result:",sum)
}
