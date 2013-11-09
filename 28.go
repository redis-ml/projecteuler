package main

import "fmt"

func main() {
  num := 1001
  n := (num - 1)/2
  result := 16 * n *(n+1)*(2*n+1)/6 + 4 * n * (n+1)/2 + 4*n + 1
  fmt.Println("result of ", num ," is ", result)
}
