package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	// implement me
	a := Point{
		x: int(s.start.x) + int(s.a),
		y: int(s.start.y) - int(s.a),
	}

	return a
}

func (s Square) Area() uint {
	// implement me
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	// implement me
	result := s.a * 4
	return result
}
