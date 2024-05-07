package main

import "fmt"

type Force interface {
	getForce() int
}

type Baseball struct {
	Mass         int
	acceleration int
}

func (b *Baseball) getForce() int {
	return b.Mass * b.acceleration
}
func (f *Football) getForce() int {
	return 50
}

type Football struct {
	Point int
}

func compareForce(a, b Force) bool {
	return a.getForce() > b.getForce()
}

func (f *Football) GetScore() int {
	return f.Point
}
func main() {

	b1 := &Baseball{
		Mass:         2,
		acceleration: 20,
	}

	f1 := &Football{Point: 11}

	fmt.Println(compareForce(b1, f1))
	fmt.Println(f1.GetScore())

	rectangle := Rectangle{length: 2, width: 4}
	square := Square{side: 2}

	fmt.Println("Rectangle perimeter: ", rectangle.rectanglePerimeter())
	fmt.Println("Square perimeter: ", square.squarePerimeter())
}

type Rectangle struct {
	length, width float64
}

type Square struct {
	side float64
}

func (r Rectangle) rectanglePerimeter() float64 {
	return (2 * r.length) + (2 * r.width)
}

func (s Square) squarePerimeter() float64 {
	return s.side * 4
}
