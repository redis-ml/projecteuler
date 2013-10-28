package main

import (
  "fmt"
)

type BTree struct {
  h uint
  in []int
  v []int
}

func n( h uint) uint {
  return h * (h + 1)/2
}
func NewBTree(h uint) *BTree {
  return &BTree{h, make([]int, n(h)), make([]int, n(h))}
}
func (b *BTree) SetLine(h uint, l []int) bool {
  hx := h + 1
  if uint(len(l)) != hx {
    return false
  }
  start := n(h)
  for i:=uint(0); i < hx; i++ {
    b.in[start + i] = l[i]
  }
  return true
}
func (b *BTree) GetLChild( h uint, i uint) (d int, v int){
  start := n(h+1)
  return b.in[start + i], b.v[start + i]
}

func (b *BTree) GetRChild( h uint, i uint) (d int, v int){
  start := n(h+1)
  return b.in[start + i + 1], b.v[start + i + 1]
}

func (b*BTree) String() string {
  s := "origin:\n"
  for i:=uint(0); i<b.h; i ++ {
    start := n(i)
    for j:=uint(0);j<i+1;j++ {
      s = fmt.Sprintf("%s %d(%d)", s, b.in[start + j], b.v[start+j])
    }
    s = fmt.Sprintf("%s\n", s)
  }
  return s
}
func (b *BTree) SetV(i uint, j uint, v int) {
  b.v[n(i)+j] = v
}
func (b *BTree) GetK(i uint, j uint) int {
  return b.in[n(i)+j]
}
