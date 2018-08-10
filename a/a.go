package main

import (
	"fmt"
	. "fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte
type Box struct {
	width, height, depth float64
	color                Color
}
type person struct {
	name string
	age  int
}
type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}
func (b *Box) SetColor(c Color) {
	b.color = c
}

type Skills []string
type Student struct {
	person
	speciality string
	Skills
	school string
}
type Employee struct {
	person
	speciality string
	Skills
	company string
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}
func (h *Student) SayHi() {
	fmt.Println(h.name, h.school)
}
func (e *Employee) SayHi() {
	fmt.Println(e.name, e.company)
}
func main() {
	var tom person
	tom.name, tom.age = "tom", 18
	bob := person{age: 25, name: "bob"}
	paul := person{"paul", 43}
	tb_older, tb_diff := Older(tom, bob)
	bp_older, bp_diff := Older(bob, paul)
	Println(tb_older, tb_diff)
	Println(bp_older, bp_diff)

	mar := Student{person: person{"ben", 22}, speciality: "sttt", school: "yo"}
	mar.Skills = []string{"a", "b", "c"}
	mar.SayHi()
	mar2 := Employee{person: person{"ben", 22}, company: "yo22"}
	mar2.SayHi()

	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{5, 4, 4, BLACK},
		Box{4, 6, 4, YELLOW},
		Box{4, 1, 4, BLUE},
	}
	Println(boxes)
	for k := range boxes {
		Println(boxes[k].Volume())
	}

	//Println(boxes[0].Volume())
}
