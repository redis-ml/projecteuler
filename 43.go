package main

import (
  "fmt"
  "math"
)

type Data struct {
  pb []bool
  p []int
  digit int
  substart int
  sublen int
  base int
}

func NewData(digit,substart,sublen uint) *Data {
  maxL := int(sublen)
  maxNum := int(math.Pow10(maxL)) - 1
  maxDiv := int(math.Sqrt(float64(maxNum)))

  pb := make([]bool, maxDiv+1)
  for i,_ := range pb { pb[i] = false }
  pb[0] = true
  pb[1] = true
  maxPb := int(math.Sqrt(float64(maxDiv+1)))
  for i:=2;i<=maxPb;i++ {
    if pb[i] { continue }
    for j:=i*i;j<maxDiv+1;j+=i {
      pb[j] = true
    }
  }
  fmt.Println("prime buffer")
  for _,j := range pb {
    fmt.Print(" ", j)
  }
  fmt.Println()

  b := make([]int, digit)
  for i,_ := range b {
    b[i] = i
  }

  base := int(math.Pow10(int(sublen)-1))
  return &Data{pb, b, int(digit), int(substart), int(sublen), base}
}

func main() {
  d := NewData(10, 1, 3)
  sum := d.per(0)
  fmt.Println("result ", sum)
}

func (d *Data) per(start int) uint64 {
  if (start == d.digit-1) {
    isOk := true
    tmp := 0
    pbi := 0
    for i:=d.substart;i<=d.digit-d.sublen;i++ {
      if i == d.substart {
        for j:=0;j<d.sublen-1;j++ {
          tmp = tmp * 10 + d.p[i+j]
        }
      }
      tmp = (tmp % d.base)*10 + d.p[i+d.sublen-1]
      for d.pb[pbi] { pbi ++ }
      //fmt.Println("   got one ", tmp, " prod ", pbi)
      if tmp % pbi != 0 {
        isOk = false
        break
      }
      pbi ++
    }
    if isOk {
      num := uint64(0)
      for _,j := range d.p { num = num * uint64(10) + uint64(j) }
      fmt.Println("found one ", num)
      return num
    }
    return 0
  }

  ret := uint64(0)
  if start > 0 {
    ret += d.per(start + 1)
  }

  for i:=start + 1; i<d.digit; i++ {
    tmp := d.p[start]
    d.p[start] = d.p[i]
    d.p[i] = tmp
    ret += d.per(start + 1)
    d.p[i] = d.p[start]
    d.p[start] = tmp
  }
  return ret
}

