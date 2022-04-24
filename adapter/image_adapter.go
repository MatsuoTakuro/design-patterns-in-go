package adapter

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

type vectorToRasterAdapter struct {
	points []Point
}

// Implement RasterImage interface
func (a vectorToRasterAdapter) GetPoints() []Point {
	return a.points
}

func (a *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	fmt.Println("generated", len(a.points), "points")
}

var pointCache = map[[16]byte][]Point{}

func (a *vectorToRasterAdapter) addLineCached(line Line) {
	hashing := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}

	h := hashing(line)
	if pts, ok := pointCache[h]; ok { // If pointCache[h] exists, ok is true and the process in the if-statement is executed.
		a.points = append(a.points, pts...)
		return
	}

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}
	// be sure to add these to the cache
	pointCache[h] = a.points
	fmt.Println("generated", len(a.points), "points")
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vi.Lines {
		adapter.addLineCached(line)
	}
	return adapter // as RasterImage
}
