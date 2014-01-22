//File_name: images.go
//Author: Wenbin Xiao
//Description: http://tour.golang.org/#62

package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

// type Image interface {
//         // ColorModel returns the Image's color model.
//         ColorModel() color.Model
//         // Bounds returns the domain for which At can return non-zero color.
//         // The bounds do not necessarily contain the point (0, 0).
//         Bounds() Rectangle
//         // At returns the color of the pixel at (x, y).
//         // At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
//         // At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
//         At(x, y int) color.Color
// }

type Image struct{
    W int
    H int
}

//Methods to implement interface Image
func (img *Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.W, img.H)
}

func (img *Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x), uint8(y), 255, 255}
}
func main() {
    m := &Image{100, 100}
    pic.ShowImage(m)
}