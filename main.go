package main

import (
	"fmt"
	"math"
)

var choice int

const dol = 63.33

func round(a float64) float64 { //Округление до копеек. Использовал, пока не разобрался с Printf
	a = a * 100
	ai := int(a)
	af := float64(ai)
	return af / 100
}

func main() {

	//Выбор программы
Start:
	fmt.Println("Введите 1 - Конвертор \nВведите 2 - Треугольник \nВведите 3 - Вклад")
	fmt.Scanln(&choice)

	if choice == 1 { // Конвертация рублей в доллары
		var rub float64
		fmt.Println("Введите сумму в Рублях")
		fmt.Scanln(&rub)
		fmt.Printf("У вас %.2f$\n", rub/dol)
	}
	if choice == 2 { //Вычисление прямоугольного треугольника
		var katOne, katTwo float64
		fmt.Print("Введите длинну первого катета: ")
		fmt.Scanln(&katOne)
		fmt.Print("Введите длинну второго катета: ")
		fmt.Scanln(&katTwo)
		gip := math.Sqrt(math.Pow(katOne, 2) + math.Pow(katTwo, 2)) //Теорема Пифагора с Math
		perim := katOne + katTwo + gip
		area := (katOne * katTwo) / 2
		fmt.Printf("Длинна гипотенузы %.2f.\n", gip)
		fmt.Printf("Длинна периметра: %.2f.\n", perim)
		fmt.Printf("Площадь треугольника: %.2f.\n", area)

	}
	if choice == 3 { //Вклад
		var sum, proc float64
		fmt.Println("Введите сумму вклада:")
		fmt.Scanln(&sum)
		fmt.Println("Введите годовой процент:")
		fmt.Scanln(&proc)
		yrs := 5 //Количество лет.
		for i := 0; i < yrs; i++ {
			incr := sum * proc / 100 //Увеличение за текущий год
			sum = sum + incr
		}
		fmt.Printf("Сумма вашего вклада через 5 лет составит %.2f.\n", sum)
	}
	if choice != 1 && choice != 2 && choice != 3 {
		println("Введите число от 1 до 3")
		goto Start
	}

}
