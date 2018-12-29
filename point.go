package main

// Point Point
type Point struct {
	X int
	Y int
}

func digistSum(n int) int {
	ret := 0
	if n < 0 {
		n = n * (-1)
	}
	for ; n > 0; n = n / 10 {
		ret += n % 10

	}
	return ret
}

// GetDigistSum GetDigistSum
func (p *Point) GetDigistSum() int {

	return digistSum(p.X) + digistSum(p.Y)
}

// Add Add
func (p *Point) Add(delta *Point) {
	p.X = p.X + delta.X
	p.Y = p.Y + delta.Y
}

// IsEqual IsEqual
func (p *Point) IsEqual(pa *Point) bool {
	return p.X == pa.X && p.Y == pa.Y
}

// SymmetryGroup  SymmetryGroup
func (p *Point) SymmetryGroup() []Point {
	var ret []Point
	ret = append(ret, *p) // 坐标本身
	if p.X == 0 && p.Y == 0 {
		return ret
	}

	if p.X == 0 {
		// exchage
		ret = append(ret, Point{p.Y, p.X})
		ret = append(ret, Point{p.Y * (-1), p.X})
		ret = append(ret, Point{p.X, p.Y * (-1)})
		return ret

	}

	if p.Y == 0 {
		ret = append(ret, Point{p.Y, p.X})

		ret = append(ret, Point{p.Y, p.X * (-1)})
		ret = append(ret, Point{p.X * (-1), p.Y})
		return ret
	}

	if p.X == p.Y {
		ret = append(ret, Point{p.X * (-1), p.Y})
		ret = append(ret, Point{p.X, p.Y * (-1)})
		ret = append(ret, Point{p.X * (-1), p.Y * (-1)})

		return ret
	}

	ret = append(ret, Point{p.X * (-1), p.Y})
	ret = append(ret, Point{p.X, p.Y * (-1)})
	ret = append(ret, Point{p.X * (-1), p.Y * (-1)})

	ret = append(ret, Point{p.Y, p.X})
	ret = append(ret, Point{p.Y, p.X * (-1)})
	ret = append(ret, Point{p.Y * (-1), p.X})
	ret = append(ret, Point{p.Y * (-1), p.X * (-1)})

	return ret
}
