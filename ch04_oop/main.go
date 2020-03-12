package main

import (
	"fmt"
)

// Mutatable 구조체
type Mutatable struct {
	a int
	b int
}

// StayTheSame ...
func (m Mutatable) StayTheSame() {
	m.a = 5
	m.b = 7
}

// Mutate ...
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

type quantity int
type costCalculator func(quantity, float64) float64
type dozen []quantity

func describe(q quantity, price float64, c costCalculator) {
	fmt.Printf("quantity: %d, price: %0.0f, cost: %0.0f\n", q, price, c(q, price))
}

type rect struct {
	width  float64
	Height float64
}

func (r rect) area() float64 {
	return r.width * r.Height
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
	case 3:
		var d dozen
		for i := quantity(1); i < 12; i++ {
			d = append(d, i)
		}
		fmt.Println(d)
	case 4:
		var offBy10Percent, offBy1000Won costCalculator
		offBy10Percent = func(q quantity, price float64) float64 {
			return float64(q) * price * 0.9
		}
		offBy1000Won = func(q quantity, price float64) float64 {
			return float64(q)*price - 1000
		}

		describe(3, 10000, offBy10Percent)
		describe(3, 1000, offBy1000Won)
	case 5:
		r := rect{3, 4}
		fmt.Println("area :", r.area())
	}
}
