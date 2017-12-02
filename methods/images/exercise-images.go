package main

import "golang.org/x/tour/pic"
import (
    "image"
    "image/color"
)

type Image struct {
    x, y, w, h int
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel    
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(i.x, i.y, i.w, i.h)    
}

func (i Image) At(x, y int) color.Color {
    //return color.RGBA{uint8(x^y), uint8(x^y), 255, 255}
    //return color.RGBA{uint8((x+y)/2), uint8((x+y)/2), 255, 255}
    return color.RGBA{uint8(x*y), uint8(x*y), 255, 255}
}

func main() {
    m := Image{0, 0, 250, 250}
    pic.ShowImage(m)
}

