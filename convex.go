package main

// ConvexHullMCA returns the set of points that enclose the
// convex hull of the polyline.
// The algorithm used is Andrew's Monotone Chain Algorithm.
// O(n log(n)).
// I chose this algorithm over Graham's Scan because it was easier
// to sort the points by x-coordinate instead of polar angle, as that
// sort was already implemented.
func ConvexHullMCA(p Points) Points {
	upperHull := make(Points, 0)
	lowerHull := make(Points, 0)

	for _, point := range p {
		for len(lowerHull) >= 2 && ccw(lowerHull.SecondToLast(), lowerHull.Last(), point) <= 0 {
			lowerHull = lowerHull[:len(lowerHull)-1]
		}
		lowerHull = append(lowerHull, point)
	}

	idx := len(p) - 1
	for idx >= 0 {
		for len(upperHull) >= 2 && ccw(upperHull.SecondToLast(), upperHull.Last(), p[idx]) <= 0 {
			upperHull = upperHull[:len(upperHull)-1]
		}
		upperHull = append(upperHull, p[idx])
		idx--
	}

	lowerHull = lowerHull[:len(lowerHull)]
	upperHull = upperHull[:len(upperHull)]

	return append(lowerHull, upperHull...)
}

// ConvexHullJM computes the convex hull of the set of
// points through the Jarvis' March algorithm.
func ConvexHullJM(p Points) Points {
	_ = "breakpoint"
	currPoint := 0
	maxPoint := 0
	minPoint := 0
	maxAngle := 0
	minAngle := 0
	usedPoints := make([]bool, len(p))

	// find the point with minimum y-value
	// and the one with maximum y-value
	for idx, point := range p {
		if point.Y < p[minPoint].Y {
			minPoint = idx
		}

		if point.Y > p[maxPoint].Y {
			maxPoint = idx
		}
	}

	usedPoints[minPoint] = true
	currPoint = minPoint

	// build the left-hand side of the hull
	for currPoint != maxPoint {
		maxAngle = currPoint
		for idx := range p {
			currAngle := angleBetweenPoints(p[currPoint], p[maxAngle])
			nextAngle := angleBetweenPoints(p[currPoint], p[idx])
			if nextAngle >= currAngle && (!usedPoints[idx] || idx == maxPoint) && nextAngle <= 270.0 {
				maxAngle = idx
			}
		}

		usedPoints[maxAngle] = true
		currPoint = maxAngle
	}

	// build the right-hand side of the hull
	currPoint = minPoint
	for currPoint != maxPoint {
		minAngle = maxPoint
		for idx := range p {
			currAngle := angleBetweenPoints(p[currPoint], p[minAngle])
			nextAngle := angleBetweenPoints(p[currPoint], p[idx])
			if nextAngle <= currAngle && (!usedPoints[idx] || idx == maxPoint) && nextAngle >= 90.0 {
				minAngle = idx
			}
		}

		usedPoints[minAngle] = true
		currPoint = minAngle
	}

	return selectPointsInHull(usedPoints, p)
}

func ccw(p1, p2, p3 Point) float64 {
	return (p2.X-p1.X)*(p3.Y-p1.Y) - (p2.Y-p1.Y)*(p3.X-p1.X)
}

func selectPointsInHull(pointMap []bool, p Points) Points {
	var finalPoints []Point
	for idx, val := range pointMap {
		if val {
			finalPoints = append(finalPoints, p[idx])
		}
	}

	return finalPoints
}
