package main

import (
	"img_go/common"
	"math/rand"
)

var (
	cr [1024][1024]uint8
	cg [1024][1024]uint8
	cb [1024][1024]uint8
)

func r(x, y int) uint8 {
	if cr[x][y] == 0 {
		if h(999) == 0 {
			cr[x][y] = uint8(h(256))
		} else {
			cr[x][y] = r((x+h(2))%common.LENGTH, (y+h(2))%common.WIDTH)
		}
	}
	return uint8(cr[x][y])
}

func g(x, y int) uint8 {
	if cg[x][y] == 0 {
		if h(999) == 0 {
			cg[x][y] = uint8(h(256))
		} else {
			cg[x][y] = g((x+h(2))%common.LENGTH, (y+h(2))%common.WIDTH)
		}
	}
	return uint8(cg[x][y])
}

func b(x, y int) uint8 {
	if cb[x][y] == 0 {
		if h(999) == 0 {
			cb[x][y] = uint8(h(256))
		} else {
			cb[x][y] = b((x+h(2))%common.LENGTH, (y+h(2))%common.WIDTH)
		}
	}
	return uint8(cb[x][y])
}

func h(n int) int {
	return rand.Int() % n
}

func main() {
	m := common.Imager{r, g, b}
	m.WritePng("second.png")
}
