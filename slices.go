//File_name: slices.go
//Author: Wenbin Xiao
//Description: http://tour.golang.org/#38

package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    //[x][y]int does not work.
    //Multiple dimentional arrays here are actually slice of slices and it need to be inititialzed.
    //The syntax var array [size]int requires size to be constant(which means that the size array
    //allocated by this syntax is determined during compile time instead of running time).
    ret := make([][]uint8, dy)
    for i := range ret {
        ret[i] = make([]uint8, dx)
        for j := range ret[i] {
            ret[i][j] = uint8(i) * uint8(j)
        }
    }
    return ret

}

func main() {
    pic.Show(Pic)
}