package proxy

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func (lb *LazyBitmap) Draw() {
	if lb.bitmap == nil {
		lb.bitmap = NewBitmap(lb.filename)
	}
	lb.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{
		filename: filename,
		bitmap:   nil,
	}
}
