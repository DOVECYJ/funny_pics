package main

import (
	"img_go/common"
	"math"
)

func r(x, y int) uint8 {
	return common.To256(common.Sq(math.Cos(atan2(x, y))))
}

func g(x, y int) uint8 {
	return common.To256(common.Sq(math.Cos(atan2(x, y) - acos2())))
}

func b(x, y int) uint8 {
	return common.To256(common.Sq(math.Cos(atan2(x, y) + acos2())))
}

func atan2(x, y int) float64 {
	return math.Atan2(float64(y-common.Y_CNT), float64(x-common.X_CNT)) / 2
}

func acos2() float64 {
	return 2 * math.Acos(-1) / 3
}

func main() {
	m := common.Imager{r, g, b}
	m.WritePng("first.png")
}
