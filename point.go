package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gonum/plot/plotter"
	"os"
	"sort"
	"strconv"
)

// Point represents a point with an X, Y coordinate and a segment label
// that indicates which line segment to which this point belongs.
type Point struct {
	X       float64
	Y       float64
	Segment int
	Left    bool
}

// Points is a type for a slice of point structs.
type Points []Point

// ***********************************
// sorting functionality for Points
// ***********************************
func (p *Points) Len() int      { return len(*p) }
func (p *Points) Swap(i, j int) { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }
func (p *Points) Less(i, j int) bool {
	if (*p)[i].X == (*p)[j].X {
		return (*p)[i].Y < (*p)[j].Y
	}
	return (*p)[i].X < (*p)[j].X
}

// Sort mutates the slice of points by sorting them
func (p *Points) Sort() {
	sort.Sort(p)
}

// Perimeter calculates the perimeter of the lines represented by the points
func (p Points) Perimeter() float64 {
	var perimeter float64
	for idx, point := range p[:len(p)-1] {
		perimeter += makeVector(point, p[idx+1]).magnitude()
	}

	return perimeter
}

// MakePoint returns a new point from a pair of floating point values
func MakePoint(x, y float64, segment int, left bool) Point {
	return Point{
		100 * x,
		100 * y,
		segment,
		left,
	}
}

// Print pretty prints the information for a given point
func (p Point) Print() string {
	var endpoint string
	if p.Left {
		endpoint = "left ="
	} else {
		endpoint = "right ="
	}

	return fmt.Sprintf("segment %d %s (%f, %f)",
		p.Segment,
		endpoint,
		p.X,
		p.Y)
}

// Print is a method on a slice of points that pretty prints the whole slice.
func (p Points) Print() {
	for index, point := range p {
		fmt.Printf("%d: %s\n", index, point.Print())
	}
}

func (p Points) toPlotterXYs() plotter.XYs {
	xys := make(plotter.XYs, len(p))
	for idx, point := range p {
		xys[idx].X = float64(point.X)
		xys[idx].Y = float64(point.Y)
	}

	return xys
}

// GeometryError is a custom error struct
type GeometryError struct {
	Message string
}

func (ge GeometryError) Error() string {
	return ge.Message
}

// ReadPointsFromCSV reads a csv file where each line is in the format x1,y1,x2,y2
// where x1,y1 form the left point of the line segment and x2,y2 form the right point.
func ReadPointsFromCSV(path string) (Points, *GeometryError) {
	_ = "breakpoint"
	fh, err := os.Open(path)
	if err != nil {
		return make(Points, 10), &GeometryError{err.Error()}
	}

	reader := csv.NewReader(fh)
	if records, err := reader.ReadAll(); err == nil {
		points := make(Points, len(records))
		for index, record := range records {
			if len(record) != 2 {
				return make(Points, 10),
					&GeometryError{fmt.Sprintf("%s, line #%d formatted improperly. Must contain two comma separated values.\n", path, index)}
			}

			// Should return a point
			newPoint, err := getPointFromRecord(record, index)
			if err != nil {
				return points, err
			}
			points[index] = newPoint
		}
		return points, nil
	}
	return make(Points, 10), &GeometryError{fmt.Sprintf("CSV error: %s\n", err.Error())}
}

func getPointFromRecord(record []string, segment int) (Point, *GeometryError) {
	x1, err := strconv.ParseFloat(record[0], 64)
	if err != nil {
		return Point{}, &GeometryError{fmt.Sprintf("Could not convert %s to float.", record[0])}
	}

	y1, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		return Point{}, &GeometryError{fmt.Sprintf("Could not convert %s to float.", record[1])}
	}

	return MakePoint(x1, y1, segment, false), nil
}
