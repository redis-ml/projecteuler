package main

import (
  "fmt"
)

// 1900 1 1
var days = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31 }

func dayByYear(year int) int {
  return (year-1900) * 365 +year/4-1900/4-(year/400-1900/400)
}
func dayInYear(year int, month int, day int) int {
  isLeap := (year % 100 != 0 && year % 4 == 0) || (year % 100 == 0 && year % 400 == 0)
  ret := 0
  for i:= 1; i<month; i++ {
    ret += days[i]
  }
  ret += day - 1
  if isLeap && month > 2 {
    ret ++
  }
  return ret
}

func calc(year int, month int, day int, v bool) int {
  ret := dayByYear(year) + dayInYear(year, month, day)
  if v {
    fmt.Printf("%04d-%02d-%02d %d\n", year,month, day, ret)
  } 
  return ret
}
func main() {
  calc(1900, 1, 1, true)
  calc(1900, 2, 28, true)
  calc(1900, 3, 1, true)
  calc(1901, 1, 1, true)
  calc(1904, 2, 28, true)
  calc(1904, 3, 1, true)
  calc(2000, 1, 1, true)
  calc(2000, 2, 1, true)
  calc(2000, 2, 29, true)
  calc(2000, 3, 1, true)

  cnt := 0
  for i:=1901;i<=2000;i++ {
    for j:=1;j<=12;j++ {
      d := calc(i, j, 1, false)
      if d % 7 == 6 {
        //fmt.Println("got ", i, j, 1)
        cnt ++
      }
    }
  }
  fmt.Println("result ", cnt)
}
