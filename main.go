package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need an extra argument for the path.")
		os.Exit(1)
	}

	path := os.Args[1]
	fmt.Printf("Reading in CSV: %s\n", path)
	points, err := ReadPointsFromCSV(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Success")
	points.Print()

	fmt.Println("Drawing Line")
	DrawLine(points, 3000, 3000, "output/original.png")
	fmt.Println("Success")
	//
	// newList := Reduce(points, points.Perimeter()*0.1)
	// newList.Print()
	// fmt.Println("Drawing New Line")
	// DrawLine(newList, 3000, 3000, "output/reduced.png")
	// fmt.Println("Success")

	fmt.Println("Making convex hull with Jarvis March")
	hull := ConvexHullJM(points)
	fmt.Println(hull)
	fmt.Println("Success")
	DrawPolygon(hull, 3000, 3000, "output/hullJM.png")

	fmt.Println("Making convex hull with monotone chain algorithm")
	points.Sort()
	hull = ConvexHullMCA(points)
	fmt.Println(hull)
	fmt.Println("Success")
	DrawPolygon(hull, 3000, 3000, "output/hullMCA.png")

}
