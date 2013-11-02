package main

import (
  "fmt"
  "math/big"
)

func main() {
  sum := uint64(0)
  i10 := big.NewInt(int64(10))
  i := big.NewInt(int64(1))
  for n:=int64(2);n<int64(100);n++ {
    bi := big.NewInt(n)
    i.Mul(i, bi)
  }

  for i.BitLen() > 0 {
    m := big.NewInt(int64(0))
    i.DivMod(i, i10, m)
    sum += m.Uint64()
  }
  fmt.Println(" ret:", sum)
}
