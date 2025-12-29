/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/methods/25
publishedTime: undefined
---

## 练习：图像

还记得之前编写的[图片生成器](https://tour.go-zh.org/moretypes/18) 吗？

我们再来编写另外一个，不过这次它将会返回一个 `image.Image` 的实现而非一个数据切片。

定义你自己的 `Image` 类型，实现[必要的方法](https://go-zh.org/pkg/image/#Image)并调用 `pic.ShowImage`。

`Bounds` 应当返回一个 `image.Rectangle` ，例如 `image.Rect(0, 0, w, h)`。

`ColorModel` 应当返回 `color.RGBAModel`。

`At` 应当返回一个颜色。

上一个图片生成器的值 `v` 对应于此次的 `color.RGBA{v, v, 255, 255}`。
*/
package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (img Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
