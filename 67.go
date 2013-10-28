package main

import (
  "fmt"
  "flag"
  "os"
  "bufio"
  "strings"
  "strconv"
)

var pdata = flag.String("data", "", "input data")
var pnum = flag.Uint("num", uint(0), "input data height")

func main() {
  flag.Parse()

  bt := NewBTree(*pnum)
  f, err := os.Open(*pdata)
  if err != nil {
    fmt.Println("failed to open file ", *pdata, err)
    return
  }
  defer f.Close()

  s := bufio.NewScanner(f)
  cnt := uint(0)
  for s.Scan() {
    line := s.Text()
    if len(line) == 0 {
      continue
    }
    l := strings.Split(line, " ")
    nl := make([]int, len(l))
    for i, s := range l {
      tmp, err2 := strconv.Atoi(s)
      if err2 != nil {
        fmt.Println("err line2: ", line)
        return
      }
      nl[i] = tmp
    }
    bt.SetLine(cnt, nl)
    cnt ++
  }
  fmt.Print("input \n" , *pdata)

  last := 0
  for i:= *pnum - 1; i >= uint(0); i-- {
    lv := 0
    rv := 0
    for j:= uint(0); j <= i; j ++ {
      //fmt.Println(" process " , i, " ", j)
      max := 0
      if i != *pnum - 1 {
        if j == 0 {
          _, lv = bt.GetLChild(i,j)
        } else {
          lv = rv
        }
        _, rv = bt.GetRChild(i,j)
        max = lv
        if max < rv { max = rv }
      }
      last = max + bt.GetK(i,j)
      bt.SetV(i,j, last)
    }
    if i == 0 { 
      fmt.Println("the answer is : " , last)
      break
    }
  }
  //fmt.Print("output \n" , bt)
}
