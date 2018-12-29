package main

import "fmt"

// can be optimized from recursive to iterative
func path(start Point) int {
	ret := 0
	upret := 0
	rightret := 0
	downret := 0
	uppoint := start
	uppoint.Y = uppoint.Y + 1
	ok := isPointOK(&uppoint)
	if ok {
		fmt.Printf("%+v\n", uppoint)
		sumpoint <- uppoint
		pathrecordpoint <- uppoint
		upret = path(uppoint)

		ret = ret + 1
	}
	rightpoint := start
	rightpoint.X = rightpoint.X + 1
	ok = isPointOK(&rightpoint)
	if ok {
		fmt.Printf("%+v\n", rightpoint)
		sumpoint <- rightpoint
		pathrecordpoint <- rightpoint
		rightret = path(rightpoint)
		ret = ret + 1

	}

	return ret + upret + downret + rightret
}

func isPointOK(p *Point) bool {
	_, ok := pointVisit[*p]
	if ok {
		return false
	}

	if p.X > xylimit {
		return false
	}

	if p.Y > xylimit {
		return false
	}

	if p.X < 0 {
		return false
	}

	if p.Y < 0 {
		return false
	}
	if p.X < p.Y {
		return false
	}

	if sumlimit < p.GetDigistSum() {
		return false
	}
	pointVisit[*p] = true
	return true
}
