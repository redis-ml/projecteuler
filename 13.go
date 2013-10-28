package main

import (
  "flag"
  "os"
  "fmt"
  "math"
  "bufio"
  _ "strings"
  "strconv"
)

var data = flag.String("data", "nil", "input data")
var prow = flag.Uint("row", 100, "input data")
var pcol = flag.Uint("col", 50, "input data")

const t9 = uint64(1000000000)

func main(){
  flag.Parse()
  
  row := *prow
  col := *pcol
  fmt.Printf(" find para %d ", *data )
  m := make([][]uint8 , int(row))
  for i,_ := range m {
    m[i] = make([]uint8, int(col))
  }
  f, err := os.Open(*data)
  if err != nil {
    fmt.Println("failed to open file ", *data)
    return
  }

  s := bufio.NewScanner(f)
  x := 0
  for s.Scan() {
    line := s.Text()
    if len(line) == 0 {
      continue
    }
    if len(line) != int(col) {
      fmt.Printf("error line :", line)
      return
    }
    for i:=0;i<int(col);i++ {
      tmp,err := strconv.Atoi(line[i:i+1])
      if err != nil {
        fmt.Println("err line2: ", line)
        return
      }
      m[x][i] = uint8(tmp)
    }
    x ++
    
  }
  if x != int(row) {
    fmt.Println("wrong row number ", x)
    return
  }

  last := uint64(0)
  remain := uint64(0)
  i := 0
  for ;i<int(col);i++ {

    sum := uint64(0)
    for j:=0;j<int(row);j++ {
      sum += uint64(m[j][int(col)-1-i])
    }
    nLast := last + sum * 10
    oldLast := last
    last = nLast / 10
    n := nLast % 10
    if i >= 10 {
      remain += t9 * 10 * n
      remain /= 10
    } else {
      remain += uint64(math.Pow10(i)) * n
    }
    fmt.Println("sum:" , sum, ", oldLast:", oldLast, ", newLast:", last, ", remain:", remain)
  }
  fn := int(math.Floor(math.Log10(float64(last))))
  fmt.Print("last sum " , last)
  for i:=10;i>fn;i-- {
    base := t9
    fmt.Print(remain/base)
    remain %= t9 
    remain *= 10
  }
}

