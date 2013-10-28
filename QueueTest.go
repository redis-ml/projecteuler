package main

import (
  "flag"
  _ "os"
  "fmt"
  _ "math"
)

var data = flag.Int("data", int(0), "input data")

const max = 1000000

func main(){
  flag.Parse()
  


  c := *data
  buf := make([]int, max)
  cache := make(map[uint]int)

  q := NewQueue(c+1)

  q.Offer(1, 1)
  buf[1] = 1
  cnt := 2
  
  for {
    d, s, err := q.Poll()
    if err {
      fmt.Println("finished")
      break
    }

    var d1 int
    var e1 bool
    for m:= 0;m<2;m++ {
      if m == 0 {
        d1, e1 = try1(d)
      } else {
        d1, e1 = try2(d)
      }
      if !e1 {
        if d1 < max {
          if buf[d1] == 0 {
      
            buf[d1] = s + 1
            cnt ++
            if cnt >= max {
              break
            }
            e2 := q.Offer(d1, s+1)
            if e2 {
              fmt.Println("failed to offer" ,d1)
            } else {
              //fmt.Println(" offer" ,d1, " ", buf[d1])
            }
          }
        } else {
          _, ok := cache[uint(d1)]
          if !ok {
            cache[uint(d1)] = s + 1
            e2 := q.Offer(d1, s+1)
            if e2 {
              fmt.Println("failed to offer" ,d1)
            } else {
              //fmt.Println(" offer" ,d1, " ", buf[d1])
            }
          }
          //fmt.Println("exceed ", d1, " c=", c)
        }
      }
    }
  }

  l := 0
  li := -1
  for i:=2;i<max;i++ {
    if buf[i] == 0 {
      fmt.Println("undone " , i)
    } else {
      if l < buf[i] {
        l = buf[i]
        li = i
      }
    }
  }
  fmt.Println("largest ", l, " li=", li)
}

func try1(d int) (int, bool) {
  ret := d * 2
  if ret < d {
    return ret, true
  }
  return ret, false
}

func try2 (d int) (int, bool) {
  if d < 4 {
    return d, true
  }
  m := d - 1
  if m % 3 != 0 {
    return d, true
  }
  if (m/3) % 2 != 0  {
    return m/3, false
  }
  return d, true
}

