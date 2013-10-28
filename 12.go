package main

import (
  "flag"
  _ "os"
  "fmt"
  "math"
)

var data = flag.Uint64("data", uint64(0), "input data")
var size = flag.Uint64("size", uint64(0), "input data")
var pnum = flag.Uint64("num", uint64(500), "input data")

func main(){
  flag.Parse()
  
  fmt.Printf(" find para %d %d", *data, *size)
  ls := NewLs()
  step := uint64(1) << 10
  marked := uint64(0)

  num := int(*pnum)
  rt := make([]uint64, num)
  rtn := make([]uint64, num)
  rti := 0

  for start := *data;;start ++ {
    // clear result buffer
    for i, _ := range rt {
      rtn[i] = 0
    }
    rti = 0

    l1 := start
    l2 := start + 1
    if start % 2 == 0 {
      l1 /= 2
    } else {
      l2 /= 2
    }
    
    l1m := uint64(math.Ceil(math.Sqrt(float64(l1))))
    l2m := uint64(math.Ceil(math.Sqrt(float64(l2))))

    i := uint64(2)

    for ;;i++{
      fmt.Println(" now process " , i)
      if i > marked {
        marked += step
        ls.Mark(marked)
      }
      if (i > l1m && i > l2m) || i > 10000 {
        break
      }
      if ls.Get(i) {
        continue
      }
      fmt.Println("   now process " , i)
      found := false
      if l1 % i == 0 {
        found = true
        for {
          l1 /= i
          rtn[rti] ++
          if l1 % i != 0 {
            break
          }
        }
        l1m = uint64(math.Ceil(math.Sqrt(float64(l1))))
      }
      if l2 % i == 0 {
        found = true
        for {
          l2 /= i
          rtn[rti] ++
          if l2 % i != 0 {
            break
          }
        }
        l2m = uint64(math.Ceil(math.Sqrt(float64(l2))))
      }
      if found {
        rt[rti] = i
        rti ++
      }
    }
    if i < l1 {
      rt[rti] = l1
      rtn[rti] ++
      rti ++
    } 
    if i < l2 {
      rt[rti] = l2
      rtn[rti] ++
      rti ++
    }
    fmt.Println("idx ", start, ", num ", start*(start+1)/2, ", facters:")
    p := uint64(1)
    for j:=0;j<num&&rtn[j]>0;j++{
      fmt.Println("  factor:", rt[j], ", num:", rtn[j])
      p *= rtn[j] + 1
    }
    if p > 500 {
      break
    }
  }

}

