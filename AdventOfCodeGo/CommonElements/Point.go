package CommonElements

type Point struct {
	X, Y int
}

func (p Point) Inside(min, max Point) bool {
	return min.X <= p.X && min.Y <= p.Y && p.X <= max.X && p.Y <= max.Y
}

func (p Point) Neighbours4() PointsCollection {
	return []Point{
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
	}
}

func (p Point) Neighbours8() PointsCollection {
	var result []Point
	for offsetY := -1; offsetY <= 1; offsetY++ {
		for offsetX := -1; offsetX <= 1; offsetX++ {
			result = append(result, Point{X: p.X + offsetX, Y: p.Y + offsetY})
		}
	}
	return result
}

func (p Point) Sub(another Point) Point {
	return Point{p.X - another.X, p.Y - another.Y}
}

func (p Point) Add(another Point) Point {
	return Point{p.X + another.X, p.Y + another.Y}
}

type PointsCollection []Point

func (c PointsCollection) Inside(min, max Point) PointsCollection {
	var result []Point
	for _, n := range c {
		if n.Inside(min, max) {
			result = append(result, n)
		}
	}
	return result
}
