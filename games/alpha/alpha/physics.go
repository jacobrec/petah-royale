package alpha
// https://rosettacode.org/wiki/Find_the_intersection_of_two_lines#Go

import (
	"fmt"
	"errors"
)

type Point struct {
	x float64
	y float64
}

type Line struct {
	slope float64
	yint float64
}

type Rectangle struct {
	x float64
	y float64
	width float64
	height float64
}


func CreateLine (a, b Point) Line {
	slope := (b.y-a.y) / (b.x-a.x)
	yint := a.y - slope*a.x
	return Line{slope, yint}
}

func EvalX (l Line, x float64) float64 {
	return l.slope*x + l.yint
}

func Rectorsect (r Rectangle, l Line) (Point, error) {
    s1 := Line{CreateLine(Point{r.x, r.y}, Point{r.x+r.width, r.y})}
    s2 := Line{CreateLine(Point{r.x+r.width, r.y}, Point{r.x+r.width, r.y+r.height})}
    s3 := Line{CreateLine(Point{r.x+r.width, r.y+r.height}, Point{r.x, r.y+r.height})}
    s4 := Line{CreateLine(Point{r.x, r.y+r.height}, Point{r.x, r.y})}

    i1 := Intersection(l, s1)
    i2 := Intersection(l, s2)
    i3 := Intersection(l, s3)
    i4 := Intersection(l, s4)

    c := Point{r.x, r.y}
    err := errors.New("The lines do not intersect")
    p := rectorsect_getBetter(c, i1, i2)
    p = rectorsect_getBetter(c, p, i3)
    p = rectorsect_getBetter(c, p, i4)

    if(p == Point{}){
        return p, err
    }
    return p, nil

}
func rectorsect_getBetter(c, p1, p2 Point) Point {
    if (p1 == Point{}) {
        return p2
    }

    if (p2 == Point{}) {
        return p1
    }

    if(dist2(c, p1) < dist2(c, p2)){
        return p1
    }

    return p2
}

func dist2(p1, p2 Point) {
    return (p2.x - p1.x) * (p2.x - p1.x)  + (p2.y - p1.y) * (p2.y - p1.y)
}

func Intersection (l1, l2 Line) (Point, error) {
	if l1.slope == l2.slope {
		return Point{}, errors.New("The lines do not intersect")
	}
	x := (l2.yint-l1.yint) / (l1.slope-l2.slope)
	y := EvalX(l1, x)
	return Point{x, y}, nil
}
