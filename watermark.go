package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	fimg1, _ := os.Open("1.png")

	defer fimg1.Close()

	img1,_,_:=
}
