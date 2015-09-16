package main

import (
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
)

// DrawLine draws the line represented by the Points
func DrawLine(p Points, width, height float64, name string) {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	paths := []*draw2d.Path{}

	// Draw a path
	for idx, point := range p[:len(p)-1] {
		path := new(draw2d.Path)
		path.MoveTo(float64(point.X), height-float64(point.Y))
		path.LineTo(float64(p[idx+1].X), height-float64(p[idx+1].Y))
		paths = append(paths, path)
	}
	gc.Stroke(paths...)
	gc.FillStroke()

	// Save to file
	draw2dimg.SaveToPngFile(name, dest)
}

// DrawClosedLine draws the line represented by the Points and closes it.
func DrawClosedLine(p Points, width, height float64, name string) {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	paths := []*draw2d.Path{}

	// Draw a path
	for idx, point := range p[:len(p)] {
		path := new(draw2d.Path)
		path.MoveTo(float64(point.X), height-float64(point.Y))
		if idx < len(p)-1 {
			path.LineTo(float64(p[idx+1].X), height-float64(p[idx+1].Y))
		} else {
			path.LineTo(float64(p[0].X), height-float64(p[0].Y))
		}
		paths = append(paths, path)
	}
	gc.Stroke(paths...)
	gc.FillStroke()

	// Save to file
	draw2dimg.SaveToPngFile(name, dest)
}

// DrawPolygon draws the polygon represented by Points on the canvas
func DrawPolygon(p Points, width, height float64, name string) {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	// Draw a closed shape
	gc.MoveTo(float64(p[0].X), height-float64(p[0].Y))
	for _, point := range p[1:] {
		gc.LineTo(float64(point.X), height-float64(point.Y))
		gc.MoveTo(float64(point.X), height-float64(point.Y))
	}
	gc.Close()
	gc.FillStroke()

	// Save to file
	draw2dimg.SaveToPngFile(name, dest)
}
