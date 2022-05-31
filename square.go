package square

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	p := s.start
	fmt.Println(&p == &s.start)
	p.x = p.x + int(s.a)
	p.y = p.y + int(s.a)
	return p
}

func (s *Square) Area() uint {
	return s.a * s.a
}

func (s *Square) Perimeter() uint {
	return s.a * 4
}
