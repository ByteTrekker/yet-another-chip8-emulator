package main

const (
	RegisterSize = 16
)

type Register struct {
	Array [RegisterSize]int32
}

func (r *Register) Set(index int, value int32) {
	r.Array[index] = value
}

func (r *Register) Add(index int, value int32) {
	r.Array[index] += value
}

func (r *Register) Get(index int) int32 {
	return r.Array[index]
}             
