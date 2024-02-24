package main

const (
	DisplayWidth  = 64
	DisplayHeight = 32
)

type Display struct {
	Array [DisplayWidth][DisplayHeight]bool
}

func (d *Display) SetPixel(x int, y int, value bool) {
	d.Array[x][y] = value
}

func (d *Display) Clear() {
	for y := 0; y < DisplayHeight; y++ {
		for x := 0; x < DisplayWidth; x++ {
			d.Array[x][y] = false
		}
	}
}
