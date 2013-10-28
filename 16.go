package main

import (
  "fmt"
  "math/big"
)

func main() {
  i := big.NewInt(int64(1))
  i10 := big.NewInt(int64(10))

  sum := uint64(0)

  i.Lsh(i, uint(1000))

  for i.BitLen() > 0 {
    m := big.NewInt(int64(0))
    i.DivMod(i, i10, m)
    sum += m.Uint64()
  }
  fmt.Println(" ret:", sum)
}
