package proxy

import "fmt"

type Image interface {
	Draw()
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}
