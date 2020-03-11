package main

import (
	"fmt"
)

// Mutatable 구조체
type Mutatable struct {
	a int
	b int
}

func (m Mutatable) StayTheSame() {
	m.a = 5
	m.b = 7
}

func (m *Mutatable) Mutate() {
	m.a = 5
	m.b = 7
}

// Item 구조체
type Item struct {
	name     string
	price    float64
	quantity int
}

// Cost 메서드
func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func main() {
	var in int
	fmt.Print("선택 하세요: ")
	fmt.Scanln(&in)

	switch in {
	case 1:
		shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
		fmt.Println(shirt.Cost())
	case 2:
		m := &Mutatable{0, 0}
		fmt.Println(m)
		m.StayTheSame()
		fmt.Println(m)
		m.Mutate()
		fmt.Println(m)
	}
}
