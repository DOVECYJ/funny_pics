package common

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const (
	LENGTH = 1024
	WIDTH  = 1024
	X_CNT  = LENGTH / 2
	Y_CNT  = WIDTH / 2
	HF     = float64(512)
)

var (
	DefaultImage Imager
)

func Sq(x float64) float64 {
	return math.Pow(x, 2)
}

func Cb(x int) float64 {
	return math.Abs(float64(x ^ 3))
}

func Cr(x int) float64 {
	return math.Cbrt(float64(x))
}

func To256(c float64) uint8 {
	return uint8(c * 255)
}

type primColor func(int, int) uint8

type rgb interface {
	r(int, int) uint8
	g(int, int) uint8
	b(int, int) uint8
}

type Imager struct {
	R primColor
	G primColor
	B primColor
}

func (m *Imager) WritePng(name string) error {
	img := image.NewRGBA(image.Rect(0, 0, LENGTH, WIDTH))
	for j := 0; j < WIDTH; j++ {
		for i := 0; i < LENGTH; i++ {
			img.Set(i, j, color.RGBA{
				m.R(i, j),
				m.G(i, j),
				m.B(i, j),
				255,
			})
		}
	}
	f, err := os.OpenFile(name, os.O_CREATE|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	return png.Encode(f, img)
}
