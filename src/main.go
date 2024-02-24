package main

import (
	"fmt"
	"os"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: chip8 <filename>")
		return
	}
	filename := os.Args[1]
	fmt.Println("Loading", filename)
	memory := Memory{}
	register := Register{}
	display := Display{}
	fmt.Print(display.Array)
	fmt.Print(register.Array)
	fmt.Print(memory.Array)
}