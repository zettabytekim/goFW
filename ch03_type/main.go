package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {
	var in int
	fmt.Print("선택 하세요: ")
	fmt.Scanln(&in)

	switch in {
	case 1:
		fmt.Println(intToUint8(100))
		fmt.Println(intToUint8(1000))
	case 2:
		s := "hello"
		fmt.Println(s[0])
		fmt.Println(s[len(s)-1])
	case 3:
		s1 := "hello"
		s2 := "안녕하세요"
		r1 := []rune(s1)
		r2 := []rune(s2)
		fmt.Printf("s1: %c %c %c %c %c\n", r1[0], r1[1], r1[2], r1[3], r1[4])
		fmt.Printf("s2: %c %c %c %c %c\n", r2[0], r2[1], r2[2], r2[3], r2[4])
	case 4:
		s1 := "hello"
		s2 := "안녕하세요"
		for i, c := range s1 {
			fmt.Printf("%c(%d)\t", c, i)
		}
		fmt.Println()
		for i, c := range s2 {
			fmt.Printf("%c(%d)\t", c, i)
		}
	case 5:
		s1 := "hello"
		s2 := "안녕하세요"

		fmt.Println(len(s1))
		fmt.Println(len(s2))
		fmt.Println(utf8.RuneCountInString(s2))
		fmt.Println(len([]rune(s2)))
	case 6:
		numbers := []int{3, 4, 5, 7, 8, 4, 6, 8, 32, 4}
		for idx, val := range numbers {
			fmt.Println(idx, val)
		}
	case 7:
		numbers := []int{3, 4, 5, 7, 8, 4, 6, 8, 32, 4}
		sum := 0
		for _, n := range numbers {
			sum += n
		}
		fmt.Println("sum: ", sum)
	case 8:
		numbers := []int{3, 4, 5, 7, 8, 4, 6, 8, 32, 4}
		sum := 0
		for i := range numbers {
			numbers[i] *= 2
			sum += numbers[i]
		}
		fmt.Println("numbers: ", numbers)
		fmt.Println("sum	: ", sum)
	case 9:
		s := []int{1, 2, 3, 4, 5, 6, 7}
		fmt.Println(s, "=", s[:3], s[3:5], s[5:])
	case 10:
		ns1 := []int{1, 2, 3}
		ns2 := []int{6, 7, 8}
		ns3 := []int{8, 9, 10, 11}

		ns1 = append(ns1, 4, 5)
		ns1 = append(ns1, ns2...)
		ns1 = append(ns1, ns3[1:3]...)

		fmt.Println(ns1)
	case 11:
		s := make([]int, 0, 3)
		for i := 0; i < 9; i++ {
			s = append(s, i)
			fmt.Printf("len: %d, cap: %d, %v\n", len(s), cap(s), s)
		}
	case 12:
		s := []int{1, 2, 3, 4, 5}

		s = insert(s, []int{-3, -2}, 0)
		fmt.Println(s)

		s = insert(s, []int{0}, 2)
		fmt.Println(s)

		s = insert(s, []int{6, 7}, len(s))
		fmt.Println(s)
	case 13:
		s := []int{1, 2, 3, 4, 5}

		s = insert(s, []int{-3, -2}, 0)
		fmt.Println(s)

		s = insert(s, []int{0}, 2)
		fmt.Println(s)

		s = insert(s, []int{6, 7}, len(s))
		fmt.Println(s)
	case 14:
		numberMap := map[string]int{}
		numberMap["one"] = 1
		numberMap["two"] = 2
		numberMap["three"] = 3
		fmt.Println(numberMap)
	case 15:
		numberMap := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
		}
		fmt.Println(numberMap)
	case 16:
		numberMap := make(map[string]int, 3)
		numberMap["one"] = 1
		numberMap["two"] = 2
		numberMap["three"] = 3
		fmt.Println(numberMap)
	case 17:
		numberMap := make(map[string]int)
		numberMap["one"] = 1
		numberMap["two"] = 2
		numberMap["three"] = 3

		for k, v := range numberMap {
			fmt.Println(k, v)
		}
	case 18:
		numberMap := map[int]string{}
		numberMap[1] = "one"
		numberMap[2] = "two"
		fmt.Println(numberMap)

		numberMap[3] = "three"
		fmt.Println(numberMap)

		numberMap[3] = "셋"
		fmt.Println(numberMap)

		delete(numberMap, 3)
		fmt.Println(numberMap)
	case 19:
		var p *int

		i := 1
		p = &i
		fmt.Println(i)
		fmt.Println(&i)
		fmt.Println(*p)
		fmt.Println(p)
	case 20:
		var p *int
		var pp **int

		i := 1
		p = &i
		pp = &p
		fmt.Println(i, *p, **pp)

		i++
		fmt.Println(i, *p, **pp)

		*p++
		fmt.Println(i, *p, **pp)

		**pp++
		fmt.Println(i, *p, **pp)
	case 21:
		type rect struct{ w, h float64 }

		var i int = 1
		var p *int = &i
		var s *rect = &rect{1, 2}

		fmt.Println(p)
		fmt.Println(s)
	}
	println("완료")
}

func insert2(s, new []int, index int) []int {
	result := make([]int, len(s)+len(new))
	position := copy(result, s[:index])
	position += copy(result[position:], new)
	copy(result[position:], s[index:])
	return result
}

func insert(s, new []int, index int) []int {
	return append(s[:index], append(new, s[index:]...)...)
}

func intToUint8(i int) (uint8, error) {
	if 0 <= i && i <= math.MaxUint8 {
		return uint8(i), nil
	}
	return 0, fmt.Errorf("%d cannot convert to uint8 type", i)
}
