package main

import (
	"img_go/common"
	"math"
)

func r(x, y int) uint8 {
	i, j, k := 0.0, 0.0, 0
	for ; k < 256; k++ {
		a := i*i - j*j + float64(x-768)/common.HF
		j = 2*i*j + float64(y-common.Y_CNT)/common.HF
		i = a
		if i*i+j*j > 4 {
			break
		}
	}
	return uint8(math.Log(float64(k)) * 47)
}

func g(x, y int) uint8 {
	return r(x, y)
}

func b(x, y int) uint8 {
	i, j, k := 0.0, 0.0, 0
	for ; k < 256; k++ {
		a := i*i - j*j + float64(x-768)/common.HF
		j = 2*i*j + float64(y-common.Y_CNT)/common.HF
		i = a
		if i*i+j*j > 4 {
			break
		}
	}
	return uint8(128 - math.Log(float64(k))*23)
}

func main() {
	m := common.Imager{r, g, b}
	m.WritePng("third.png")
}
