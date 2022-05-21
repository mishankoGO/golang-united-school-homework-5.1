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
	var end Point
	end.x = s.start.x + int(s.a)
	end.y = s.start.y - int(s.a)
	return end
}

func (s Square) Area() uint {
	// implement me
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	// implement me
	return s.a * 4
}
