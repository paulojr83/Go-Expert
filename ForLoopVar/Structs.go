package main

import (
	"fmt"
	"strconv"
)

type Car struct {
	Make string
	Year int
	Used bool
}

func (c *Car) Describe() string {
	used := ""
	if c.Used {
		used = "a used car"
	} else {
		used = "a new car"
	}
	return "This " + strconv.Itoa(c.Year) + " " + c.Make + " is " + used
}

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Classroom struct {
	Id          int       `json:"id"`
	Capacity    int       `json:"capacity"`
	Subject     string    `json:"subject"`
	StudentList []Student `json:"studentList"`
}

func main() {
	car := &Car{
		Make: "Honda",
		Year: 1985,
		Used: true,
	}
	fmt.Println(car.Describe())
	fmt.Println(*car)

	c1 := Classroom{
		Id:       1,
		Capacity: 200,
		Subject:  "Art",
		StudentList: []Student{
			{
				Id:   1,
				Name: "Paulo",
			},
			{
				Id:   2,
				Name: "John",
			},
		},
	}

	c2 := new(Classroom)
	c2.Id = 2
	c2.Capacity = 100
	c2.Subject = "Math"
	c2.StudentList = []Student{
		{
			Id:   3,
			Name: "Mary",
		},
		{
			Id:   4,
			Name: "Kell",
		},
	}

	fmt.Println(c1)
	fmt.Println(c2)
}
