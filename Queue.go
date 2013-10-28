package main

import (
  _ "fmt"
)

type Queue struct {
  cap int
  rPos int
  wPos int
  in []int
  in2[]int
}

func NewQueue(cap int) *Queue {
  return &Queue{cap, 0, 0, make([]int, cap), make([]int, cap)}
}

func (q *Queue) Offer(d int, s int) bool {
  n := (q.wPos+1)%q.cap
  if n == q.rPos {
    return true
  }
  q.in[q.wPos] = d
  q.in2[q.wPos] = s
  q.wPos = n
  return false
}

func (q *Queue) Poll() (int,int,bool) {
  if q.wPos == q.rPos {
    return 0, 0, true
  }
  d := q.in[q.rPos]
  s := q.in2[q.rPos]
  q.rPos = (q.rPos + 1) % q.cap
  return d, s, false
}

