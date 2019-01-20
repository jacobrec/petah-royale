package alpha

// https://rosettacode.org/wiki/Find_the_intersection_of_two_lines#Go

import (
	"errors"
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Line struct {
	slope float64
	yint  float64
}

type Ray struct {
	xstart float64
	isXPos bool
	slope  float64
	yint   float64
}

type LineSeg struct {
	start Point
	end   Point
}

type Rectangle struct {
	x      float64
	y      float64
	width  float64
	height float64
}

func CreateRay(a, b Point) Ray {
	slope := (b.y - a.y) / (b.x - a.x)
	yint := a.y - slope*a.x
	return Ray{a.x, a.x < b.x, slope, yint}
}

func CreateLine(a, b Point) Line {
	slope := (b.y - a.y) / (b.x - a.x)
	var yint float64
	if a.x == b.x {
		yint = a.x
	} else {
		yint = a.y - slope*a.x
	}
	return Line{slope, yint}
}

func EvalX(l Line, x float64) float64 {
	return l.slope*x + l.yint
}

func IsRectorsect(r Rectangle, l LineSeg) bool {
	_, e := Rectorsect(r, l)
	return e == nil
}

func SegSegIntersect(l1, l2 LineSeg) (Point, error) {
	if isIntersect(l1, l2) {
		ll1 := CreateLine(l1.start, l1.end)
		ll2 := CreateLine(l2.start, l2.end)
		return Intersection(ll1, ll2)
	}
	return Point{0, 0}, errors.New("Line segments do not intersect")
}

func Rectorsect(r Rectangle, l LineSeg) (Point, error) {
	s1 := LineSeg{Point{r.x, r.y}, Point{r.x + r.width, r.y}}
	s2 := LineSeg{Point{r.x + r.width, r.y}, Point{r.x + r.width, r.y + r.height}}
	s3 := LineSeg{Point{r.x + r.width, r.y + r.height}, Point{r.x, r.y + r.height}}
	s4 := LineSeg{Point{r.x, r.y + r.height}, Point{r.x, r.y}}

	goodPoints := make([]Point, 0)

	p, e := SegSegIntersect(s1, l)
	if isValid(p, e) {
		goodPoints = append(goodPoints, p)
	}
	p, e = SegSegIntersect(s2, l)
	if isValid(p, e) {
		goodPoints = append(goodPoints, p)
	}
	p, e = SegSegIntersect(s3, l)
	if isValid(p, e) {
		goodPoints = append(goodPoints, p)
	}
	p, e = SegSegIntersect(s4, l)
	if isValid(p, e) {
		goodPoints = append(goodPoints, p)
	}
	fmt.Println(goodPoints)

	if len(goodPoints) == 0 {
		return Point{0, 0}, errors.New("Rectorsect, no point")
	} else if len(goodPoints) == 1 {
		return goodPoints[0], nil
	}

	c := l.start
	pp := goodPoints[0]
	if dist2(goodPoints[1], c) < dist2(pp, c) {
		pp = goodPoints[1]
	}

	return pp, nil
}
func isValid(p Point, e error) bool {
	if math.Abs(p.x) < 0.001 && math.Abs(p.y) < 0.001 {
		return false
	}
	return e == nil && !(p.x == 0 && p.y == 0)
}

func isXinRay(x float64, l Ray) bool {
	if l.isXPos {
		return x >= l.xstart
	} else {
		return x <= l.xstart
	}
}

func dist2(p1, p2 Point) float64 {
	return (p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y)
}

// Checks infinite line intersection
func Intersection(l1, l2 Line) (Point, error) {
	if l1.slope == l2.slope {
		return Point{0, 0}, errors.New("The lines do not intersect")
	}
	if math.IsInf(l1.slope, 0) {
		x := l1.yint
		y := EvalX(l2, l1.yint)
		return Point{x, y}, nil
	}
	if math.IsInf(l2.slope, 0) {
		x := l2.yint
		y := EvalX(l1, l2.yint)
		return Point{x, y}, nil
	}
	x := (l2.yint - l1.yint) / (l1.slope - l2.slope)
	y := EvalX(l1, x)
	return Point{x, y}, nil
}

func PointInRectangle(p Point, r Rectangle) bool {
	if p.x >= r.x && p.x <= r.x+r.width &&
		p.y >= r.y && p.y <= r.y+r.height {
		return true
	}
	return false
}

func onLine(l1 LineSeg, p Point) bool { //check whether p is on the line or not
	if p.x <= math.Max(l1.start.x, l1.end.x) && p.x <= math.Min(l1.start.x, l1.end.x) &&
		(p.y <= math.Max(l1.start.y, l1.end.y) && p.y <= math.Min(l1.start.y, l1.end.y)) {
		return true
	}
	return false
}

func direction(a, b, c Point) int {
	val := (b.y-a.y)*(c.x-b.x) - (b.x-a.x)*(c.y-b.y)
	if val == 0 {
		return 0 //colinear
	} else if val < 0 {
		return 2 //anti-clockwise direction
	}
	return 1 //clockwise direction
}

func isIntersect(l1, l2 LineSeg) bool {
	//four direction for two lines and points of other line
	dir1 := direction(l1.start, l1.end, l2.start)
	dir2 := direction(l1.start, l1.end, l2.end)
	dir3 := direction(l2.start, l2.end, l1.start)
	dir4 := direction(l2.start, l2.end, l1.end)

	if dir1 != dir2 && dir3 != dir4 {
		return true //they are intersecting
	}

	if dir1 == 0 && onLine(l1, l2.start) { //when end of line2 are on the line1
		return true
	}

	if dir2 == 0 && onLine(l1, l2.end) { //when start of line2 are on the line1
		return true
	}

	if dir3 == 0 && onLine(l2, l1.start) { //when end of line1 are on the line2
		return true
	}

	if dir4 == 0 && onLine(l2, l1.end) { //when start of line1 are on the line2
		return true
	}

	return false
}

func FmtDistractor() {
	fmt.Println("Suck it go!")
	fmt.Println(math.Abs(4))

}

func TestRectorsector() {
	fmt.Println("TESTS")

	fmt.Println(SegSegIntersect(LineSeg{Point{2, 3}, Point{4, 3}}, LineSeg{Point{3, 2}, Point{3, 4}})) // hit at (3, 3)
}
