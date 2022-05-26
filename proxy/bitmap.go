package proxy

import "fmt"

type Bitmap struct {
	filename string
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image of", b.filename)
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{
		filename: filename,
	}
}
