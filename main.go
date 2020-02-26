package main

import (
	"fmt"

	calculator "./pkg"
)

//1
type iStruct interface {
	printStruct()
}
type myStruct struct {
	stuff     int
	moreStuff int
}
type anotherMyStruct struct {
	stuff        int
	moreStuff    int
	anotherStuff int
}

func (s myStruct) printStruct() {
	fmt.Println("stuff=", s.stuff)
	fmt.Println("moreStuff=", s.moreStuff)
}
func (s anotherMyStruct) printStruct() {
	fmt.Println("stuff=", s.stuff)
	fmt.Println("moreStuff=", s.moreStuff)
	fmt.Println("anotherStuff=", s.anotherStuff)
}
func doStuff(s iStruct) {
	s.printStruct()
}

//2

type person struct {
	name  string
	phone int
}
type phoneBook []person

func (a phoneBook) Len() int           { return len(a) }
func (a phoneBook) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a phoneBook) Less(i, j int) bool { return a[i].name < a[j].name }

//4
var coord = map[int]string{
	1: "a",
	2: "b",
	3: "c",
	4: "d",
	5: "e",
	6: "f",
	7: "g",
	8: "h",
}

type point struct {
	x int
	y int
}

func (p *point) getX() int {
	return p.x
}
func (p *point) getY() int {
	return p.y
}

func (p point) string() string {
	return fmt.Sprintf("%s%d", coord[p.x], p.y)
}
func newPoint(x, y int) point {
	return point{
		x: x,
		y: y,
	}
}

var availableMoves = []point{
	newPoint(1, 2),
	newPoint(-1, 2),
	newPoint(2, 1),
	newPoint(-2, 1),
	newPoint(1, -2),
	newPoint(-1, -2),
	newPoint(-2, -1),
	newPoint(2, -1),
}

type fig struct {
	position point
}

func (k *fig) getPos() *point {
	return &k.position
}

func (k *fig) calcAvail() []point {
	res := make([]point, 0, 4)
	for _, v := range availableMoves {
		newX := k.getPos().getX() + v.getX()
		newY := k.getPos().getY() + v.getY()
		if isAvail(newX, newY) {
			res = append(res, newPoint(newX, newY))
		}
	}

	return res
}

func isAvail(x, y int) bool {
	return x >= 1 && x <= 8 && y >= 1 && y <= 8
}

func newFig(x, y int) fig {
	return fig{
		position: point{
			x: x,
			y: y,
		},
	}
}

func main() {
	{ //1
		mStr := &myStruct{
			stuff:     4,
			moreStuff: 42,
		}
		anotherMStr := &anotherMyStruct{
			stuff:     4,
			moreStuff: 42,
		}
		doStuff(mStr)
		doStuff(anotherMStr)
	}
	{ //2
		var pBook phoneBook
		pBook = append(pBook, person{
			name:  "Vasya",
			phone: 89999843177,
		})
		pBook = append(pBook, person{
			name:  "Kolya",
			phone: 89171234365,
		})
		pBook = append(pBook, person{
			name:  "Nikita",
			phone: 89255423145,
		})
	}
	{ //3
		calcHelp := "Spravka, ToDo" //2bDone
		input := ""
		for {
			fmt.Print("> ")
			if _, err := fmt.Scanln(&input); err != nil {
				fmt.Println(err)
				continue
			}
			if input == "help" {
				fmt.Println(calcHelp)
				continue
			}
			if input == "exit" {
				break
			}

			if res, err := calculator.Calculate(input); err == nil {
				fmt.Printf("Результат: %v\n", res)
			} else {
				fmt.Println("Не удалось произвести вычисление")
			}
		}
	}
	{ //4
		f := newFig(2, 4)
		fmt.Println(f.calcAvail())
	}
}
