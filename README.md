# funny_pics

之前有一个帖子，用3段140字符以内的代码生成一张1024×1024的图片，结果出现了很多有意思的图片。当时代码是用C语言写的，的确很巧妙。

Go语言页有`image`库，可以方便的处理2D图像，于是我复现了其中几张图。

![first](https://raw.githubusercontent.com/DOVECYJ/funny_pics/master/pictures/first.png)

![second](https://raw.githubusercontent.com/DOVECYJ/funny_pics/master/pictures/second.png)

![third](https://raw.githubusercontent.com/DOVECYJ/funny_pics/master/pictures/third.png)

你需要做的是实现三个返回RGB值的函数，然后就可画出图像了。

```go
import "github.com/dovecyj/funy_pic/common"

func r(x, y int) uint8 {
    // your code
}

func g(x, y int) uint8 {
    // your code
}

func b(x, y int) uint8 {
    // your code
}

func main() {
    img := common.Imager{r, g, b}
    img.WritePng("1.png")
}
```

Have fun!
