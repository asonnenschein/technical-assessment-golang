package models

import (
	"image"
	"image/color"
)

type GrayScale struct {
	Image image.Image
}

func (c *GrayScale) ColorModel() color.Model {
	return color.GrayModel
}

func (c *GrayScale) Bounds() image.Rectangle {
	return c.Image.Bounds()
}

func (c *GrayScale) At(x int, y int) color.Color {
	return color.GrayModel.Convert(c.Image.At(x, y))
}
