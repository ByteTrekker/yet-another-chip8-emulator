package main

const (
	MemorySize = 4096
)

type Memory struct {
	Array [MemorySize]int32
}

func (m *Memory) LoadProgram(filename string) {
}

func (m *Memory) Dump() {
}
