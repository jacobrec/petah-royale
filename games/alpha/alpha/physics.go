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

type Rectangle struct {
	x      float64
	y      float64
	width  float64
	height float64
}

func CreateLine(a, b Point) Line {
	slope := (b.y - a.y) / (b.x - a.x)
	yint := a.y - slope*a.x
	return Line{slope, yint}
}

func EvalX(l Line, x float64) float64 {
	return l.slope*x + l.yint
}

func IsRectorsect(r Rectangle, l Line) bool {
	_, e := Rectorsect(r, l)
	return e == nil
}

func Rectorsect(r Rectangle, l Line) (Point, error) {
	s1 := CreateLine(Point{r.x, r.y}, Point{r.x + r.width, r.y})
	s2 := CreateLine(Point{r.x + r.width, r.y}, Point{r.x + r.width, r.y + r.height})
	s3 := CreateLine(Point{r.x + r.width, r.y + r.height}, Point{r.x, r.y + r.height})
	s4 := CreateLine(Point{r.x, r.y + r.height}, Point{r.x, r.y})

	p1, _ := Rectorsector(l, s1, r)
	p2, _ := Rectorsector(l, s2, r)
	p3, _ := Rectorsector(l, s3, r)
	p4, _ := Rectorsector(l, s4, r)

	c := Point{r.x, r.y}
	err := errors.New("The lines do not intersect")
	p := rectorsect_getBetter(c, p1, p2, r)
	p = rectorsect_getBetter(c, p, p3, r)
	p = rectorsect_getBetter(c, p, p4, r)

	if (p == Point{}) {
		return p, err
	}
	return p, nil
}

func Rectorsector(l, s Line, r Rectangle) (Point, error){
	i1, e := Intersection(l, s)
    if e != nil {
        return Point{}, e
    }

	err := errors.New("The lines do not intersect in the bounds")
    if !PointInRectangle(i1, r){
        return Point{}, err
    }

    return i1, nil


}

func rectorsect_getBetter(c, p1, p2 Point, r Rectangle) Point {
	if (p1 == Point{} || math.IsNaN(p1.x) || math.IsNaN(p1.y)) {
        if PointInRectangle(p2, r) {
            return p2
        }
	}

	if (p2 == Point{} || math.IsNaN(p2.x) || math.IsNaN(p2.y)) {
        if PointInRectangle(p1, r) {
            return p1
        }
	}

	if dist2(c, p1) < dist2(c, p2) {
		return p1
	}

	return p2
}

func dist2(p1, p2 Point) float64 {
	return (p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y)
}

// Checks infinite line intersection
func Intersection(l1, l2 Line) (Point, error) {
	if l1.slope == l2.slope {
		return Point{}, errors.New("The lines do not intersect")
	}
	x := (l2.yint - l1.yint) / (l1.slope - l2.slope)
	y := EvalX(l1, x)
	return Point{x, y}, nil
}

func PointInRectangle(p Point, r Rectangle)  bool{
	if p.x >= r.x && p.x <= r.x+r.width &&
		p.y >= r.y && p.y <= r.y+r.height {
            return true
	}
    return false
}

func FmtDistractor() {
	fmt.Println("Suck it go!")
}

func TestRectorsector() {
	fmt.Println("TESTS")

	r := Rectangle{0, 0, 10, 10}
	l := CreateLine(Point{20, 20}, Point{5, 5})
	fmt.Println(Rectorsect(r, l)) // hit

	r = Rectangle{10, 10, 10, 10}
	l = CreateLine(Point{0, 0}, Point{20, 20})
	fmt.Println(Rectorsect(r, l)) // hit

	r = Rectangle{10, 10, 10, 10}
	l = CreateLine(Point{11, 11}, Point{9, 9})
	fmt.Println(Rectorsect(r, l)) // hit

	r = Rectangle{10, 10, 10, 10}
	l = CreateLine(Point{10, 11}, Point{9, 9})
	fmt.Println(Rectorsect(r, l)) // hit

	r = Rectangle{25, 19, 5, 2}
	l = CreateLine(Point{0, 0}, Point{25, 22})
	fmt.Println(Rectorsect(r, l)) // miss

	r = Rectangle{48, 0, 2, 35}
	l = CreateLine(Point{0, 0}, Point{2, 7})
	fmt.Println(Rectorsect(r, l)) // miss

	r = Rectangle{25, 19, 5, 2}
	l = CreateLine(Point{29, 17}, Point{29, 22})
	fmt.Println(Rectorsect(r, l)) // hit

}
