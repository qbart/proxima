package bmo

func (s *Screen) position(x, y int32) Point {
	return Point{x, s.Rect.H - y - 1}
}
