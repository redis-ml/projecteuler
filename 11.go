package main

import (
  "flag"
  "os"
  "fmt"
  _ "math"
  "bufio"
  "strings"
  "strconv"
)

const block = uint64(1)<<25

var data = flag.String("data", "nil", "input data")

func main(){
  flag.Parse()
  
  fmt.Printf(" find para %d ", *data )
  m := make([][]uint , 20)
  for i,_ := range m {
    m[i] = make([]uint, 20)
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
    ls := strings.Split(line, " ")
    if len(ls) != 20 {
      fmt.Printf("error line :", line)
      return
    }
    for i,str := range ls {
      tmp, err := strconv.Atoi(str)
      if err != nil {
        fmt.Println("err line2: ", line)
        return
      }
      m[x][i] = uint(tmp)
    }
    x ++
    
  }
  W:=20
  H:=20
  largest := uint(0)
  for i:=0;i<H;i++ {

    for j:=0;j<W-4;j++ {
      h:= m[i][j] * m[i][j+1] *m[i][j+2] *m[i][j+3]
      if h >largest {
        largest = h
      }
      if i>=3 {
        d1 := m[i][j] * m[i-1][j+1] *m[i-2][j+2] *m[i-3][j+3]
        if d1 > largest {
          largest = d1
        }
      }
      if i < H - 4 {
        d1 := m[i][j] * m[i+1][j+1] *m[i+2][j+2] *m[i+3][j+3]
        if d1 > largest {
          largest = d1
        }
      }
    }
  }
  fmt.Printf("not found %q largest %d", m, largest)
}

