//File_name: rot13reader.go
//Author: Wenbin Xiao
//Description: http://tour.golang.org/#63

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	size, _ := r.r.Read(p)
	for i, v := range p[:size] {
		if (v > byte('z') || v < byte('a')) && (v > byte('Z') || v < byte('A')) {
			continue
		}
		if v+13 > byte('z') {
			p[i] = v - 13
		} else {
			p[i] = v + 13
		}
	}
	return size, nil
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
