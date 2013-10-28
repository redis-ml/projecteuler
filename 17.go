package main

import (
  "fmt"
  _ "math"
  "flag"
)

var base = [...]string{
  "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", 
  "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty" ,
}
var tens = [...] string {
  "zero", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety" ,
}

func output(d int) string {
  if d >= 1000 {
    return "one thousand"
  } else if d >= 100 {
    n := d / 100
    if d % 100 == 0 {
      return base[n] + " handred"
    } else {
      return base[n] + " handred and " + output(d%100)
    }
  } else if d > 20 {
    n := d / 10
    if d % 10 == 0 {
      return tens[n]
    } else {
      return tens[n] + "-" + output(d%10)
    }
  } else {
    return base[d]
  }
}

var data = flag.Int("data", 0, "input data")

func main() {
  flag.Parse()
  s := ""
  for i:=1;i<=*data;i++ {
    o := output(i)
    s+= " " + o
  }
  cnt := 0
  for i:=0;i<len(s);i++ {
    sub := s[i:i+1]
    if sub != " " && sub != "-" {
      cnt ++
    }
  }
  fmt.Println("input ", *data, " output:" , cnt )
}

