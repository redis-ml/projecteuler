package main

import (
  "math"
  "fmt"
)

func main() {
  for n:=9;n<10;n++ {
    num, sum := tryN(math.MaxInt32)
    fmt.Println("with ", n, "-digit limit, got ", num, " with sum ", sum)
    if num == 11 {
      break
    }
  }
}


func tryN(maxN uint64) (total, sum uint64) {
  //maxF := math.Pow10(n)
  ls := NewLs()
  total = 0
  sum = 0

  maxF := float64(maxN)
  ln := int(math.Log10(maxF)) + 1
  rec := make([]bool, ln)
  for i,_ := range rec {
    rec[i] = false
  }

  marked := uint64(0)
  markStep := uint64(2000000000)
  ten := uint64(10)

  for i:=uint64(11);i<=maxN;i+=2 {
    if i >= marked {
      marked += markStep
      ls.Mark(marked)
    }
    if ls.Get(i) { continue }

    ln := int(math.Log10(float64(i)))
    if ln>1 && !rec[ln-1] {
      fmt.Println("strange ", ln)
      break }
    if ln > 0 {
      isOk := true

      for tmp := i/ten; tmp > uint64(0); tmp /= ten {
        if ls.Get(tmp) {
          isOk = false
          break
        }
      }
      if isOk {
        for base := uint64(math.Pow10(ln)); base >= ten; base /= ten {
          if ls.Get(i%base) {
            isOk = false
            break
          }
        }
      }
      rec[ln] = true
      if isOk {
        total ++
        fmt.Println("got one ", i)
        sum += i
      }
    }

  }
  return total, sum
}
