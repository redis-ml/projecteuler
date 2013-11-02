package main

import (
  "bufio"
  "fmt"
  "flag"
  "os"
  "strings"
)

var i = flag.String("input", "", "input file")
var o = flag.String("output", "", "output file")

func scout(s string, fo *os.File) {
  for {
    l := strings.Index(s, "\"") 
    if l < 0 { break }
    s = s[l+1:]
    r := strings.Index(s, "\"")
    if r < 0 { break }
    fo.WriteString(s[0:r])
    fo.WriteString("\n")
    s = s[r+1:]
  }
} 

func main() {
  flag.Parse()
  fi, ei := os.Open(*i)
  if ei != nil {
    fmt.Println("failed to open input file", *i)
    return
  }
  defer fi.Close()

  fo, eo := os.Create(*o)
  if eo != nil {
    fmt.Println("failed to open ouput file", *o)
    return
  }
  defer fo.Close()

  s := bufio.NewScanner(fi)
  for s.Scan() {
    line := s.Text()
    scout(line, fo)
  }

}
