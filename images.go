package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
	color         uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	v := i.color + uint8(x^y+x^y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{200, 200, 150}
	pic.ShowImage(&m)
}
