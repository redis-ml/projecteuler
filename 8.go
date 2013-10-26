package main

import (
  "flag"
  "os"
  "fmt"
  _ "math"
  "bufio"
)


type Buf struct {
  buf []uint
  cap uint
  rPos uint
  wPos uint
  src *string
  srcPos uint
  p uint
}
func NewBuf(cap uint) *Buf {
  return &Buf{ make([]uint, cap+1), cap+1, 0, 0, nil, 0, 1}
}
func (b *Buf) feed(src *string) {
  b.src = src
  b.srcPos = 0
}
func (b *Buf) product() uint {
  return b.p
}
func (b *Buf) shift() {
  old := b.buf[b.rPos]
  b.rPos = (b.rPos +1) % b.cap
  b.p /= old
  fmt.Println("poping " , old, " newP ", b.p)
}
var cnt = 0
func (b *Buf) fill() bool {
  for (b.wPos+1)%b.cap != b.rPos {
    if b.srcPos == uint(len(*(b.src))) {
      return false
    }
    n := uint((*(b.src))[b.srcPos]- '0')
    b.srcPos ++
    cnt ++
    fmt.Println("input ", n, " cnt " , cnt)
    if n == 0 {
       b.rPos = 0
       b.wPos = 0
       b.p = 1
       fmt.Println("newP ", b.p)
       continue
    }
    b.buf[b.wPos] = n
    b.wPos = (b.wPos + 1) % b.cap
    b.p *= n
    fmt.Println("newP ", b.p)
  }
  return true
}

var data = flag.String("data", "", "input data")

func main(){
  flag.Parse()
  f, err := os.Open(*data)
  if err != nil {
    fmt.Println("error opening input data %s", *data)
    return
  }

  buf := NewBuf(5)
  largestP := uint(0)
  scanner := bufio.NewScanner(f)

  for {
    if !scanner.Scan() {
      break
    }
    line := scanner.Text()
    buf.feed(&line)
    for {
      full := buf.fill()
      if !full {
        break
      }
     
      newP := buf.product()
      if newP > largestP {
        largestP = newP
      }
      buf.shift()
    }
  }
  fmt.Println("largestP = ", largestP)  
}

