package core

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

type BasicImage struct {
	Path   string
	Size   int
	Suffix string
}

type WaterImage struct {
	Position    Middle
	Size        int
	Path        string
	Transparent float64
}

type CreatImage struct {
	Suffix string
	Path   string
}

const (
	TopLeftCorner    = iota //左上角
	TopRightCorner          //右上角
	Middle                  //居中
	LowerLeftCorner         //左下角
	LowerRightCorner        //右下角
)

func (b *BasicImage) GetBasicImage(path string) *BasicImage {
	img := &BasicImage{
		Path:   path,
		Size:   100,
		Suffix: "png",
	}

	return img
}
