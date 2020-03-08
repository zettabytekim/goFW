package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var input int
	fmt.Print("테스트 선택: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		forStmt1()
	case 2:
		forStmt2()
	case 3:
		forStmt3()
	case 4:
		switch1()
	case 5:
		findX()
	case 6:
		findX2()
	case 7:
		findX3()
	case 8:
		findX4()
	case 9:
		v := area(3, 4)
		fmt.Println(v)
	case 10:
		w, h := multiply(3, 4)
		fmt.Println(w, h)
	case 11:
		displayInt("two")
		displayInt("2")
	case 12:
		i := 10
		inc(i)
		fmt.Println(i)
	case 13:
		i := 10
		inc2(&i)
		fmt.Println(i)
	case 14:
<<<<<<< HEAD
		i := 10
		inc2(&i)
		fmt.Println(i)
	case 15:
		fmt.Println("클로져를 사용한 팩토리 함수")
		addZip := makeSuffix(".zip")
		addTgz := makeSuffix(".tar.gz")
		fmt.Println(addTgz("go1.5.1.src"))
		fmt.Println(addZip("go1.5.1.windows-amd64"))
	case 16:
		fmt.Println("함수를 매개변수로 전달")
		callback(1, add)
	case 17:
		f := func(c rune) bool {
			return unicode.Is(unicode.Hangul, c)
		}
		fmt.Println(strings.IndexFunc("hello, 월드", f))
		fmt.Println(strings.IndexFunc("hello, world", f))
=======
		f1()
	case 15:
		f()
	case 16:
		b()
>>>>>>> d8de704be910e22b7bdc7ad97fa27b1c491549b3
	}
	println("완료")
}

<<<<<<< HEAD
func callback(y int, f func(int, int)) {
	f(y, 2)
}
func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

=======
func enter(s string) { fmt.Println("entering:", s) }
func leave(s string) { fmt.Println("leaving:", s) }

func a() {
	enter("a")
	defer leave("a")
	fmt.Println("ia a")
}

func b() {
	enter("b")
	defer leave("b")
	fmt.Println("in b")
	a()
}
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func f1() {
	fmt.Println("f1 - start")
	defer f2()
	fmt.Println("f1 - end")
}

func f2() {
	fmt.Println("f2 - deferred")
}

>>>>>>> d8de704be910e22b7bdc7ad97fa27b1c491549b3
func inc2(i *int) {
	*i = *i + 1
}

func inc(i int) {
	i = i + 1
}

func displayInt(s string) {
	if v, err := strconv.Atoi(s); err != nil {
		fmt.Printf("%s는 정수가 아닙니다.\n", s)
	} else {
		fmt.Printf("정수 값은 %d입니다.\n", v)
	}
}

func multiply(w, h int) (int, int) {
	return w * 2, h * 2
}

func area(w, h int) int {
	return w * h
}

func findX4() {
	x := 5
	table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}

LOOP:
	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[row]); col++ {
			if table[row][col] == x {
				fmt.Printf("found(24) %d (row:%d, col:%d)\n", x, row, col)
				continue LOOP
			}
		}
	}
}

func findX3() {
	x := 7
	table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}

LOOP:
	for row, rowValue := range table {
		for col, colValue := range rowValue {
			if colValue == x {
				fmt.Printf("found(3) %d (row:%d, col:%d)\n", x, row, col)
				break LOOP
			}
		}
	}
}

func findX2() {
	x := 7
	table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}

LOOP:
	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[row]); col++ {
			if table[row][col] == x {
				fmt.Printf("found(2) %d (row:%d, col:%d)\n", x, row, col)
				break LOOP
			}
		}
	}
}

func findX() {
	x := 7
	table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}
	found := false

	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[row]); col++ {
			if table[row][col] == x {
				found = true
				fmt.Printf("found %d(row:%d, col:%d)\n", x, row, col)
				break
			}
		}
		if found {
			break
		}
	}
}

func switch1() {
	c := 'a'

	switch {
	case '0' <= c && c <= '9':
		fmt.Printf("%c은(는) 숫자입니다", c)
	case 'a' <= c && c <= 'z':
		fmt.Printf("%c은(는) 소문자입니다", c)
	case 'A' <= c && c <= 'Z':
		fmt.Printf("%c은(는) 숫자입니다", c)
	}
	fmt.Println()
}

func forStmt3() {
	sum, i := 0, 0
	for {
		if i >= 10 {
			break
		}
		sum += i
		i++
	}
	fmt.Println("forStmt3:", sum)
}

func forStmt2() {
	sum, i := 0, 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println("forStmt2:", sum)
}

func forStmt1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("forStmt1:", sum)
}
