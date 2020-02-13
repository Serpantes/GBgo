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

func conv(rSum float64) float64 {
	return rSum / dol
}
func gip(katO float64, katT float64) float64 {
	return math.Sqrt(math.Pow(katO, 2) + math.Pow(katT, 2)) //Теорема Пифагора с Math
}
func perim(katO float64, katT float64) float64 { //Периметр
	return katO + katT + gip(katO, katT)
}
func area(katO float64, katT float64) float64 { //Площадь
	return (katO * katT) / 2
}
func bank(rsum float64, proc float64) float64 {

	yrs := 5 //Количество лет.
	for i := 0; i < yrs; i++ {
		incr := rsum * proc / 100 //Увеличение за текущий год
		rsum = rsum + incr
	}
	return rsum
}
func main() {

	//Выбор программы
	for choice < 1 || choice > 3 {
		fmt.Println("Введите 1 - Конвертор \nВведите 2 - Треугольник \nВведите 3 - Вклад")
		fmt.Scanln(&choice)
	}
	if choice == 1 { // Конвертация рублей в доллары
		var rub float64
		fmt.Println("Введите сумму в Рублях")
		fmt.Scanln(&rub)
		fmt.Printf("У вас %.2f$\n", conv(rub)) //Не совсем понимаю где тут использовать функцию. Весь ввод-вывод в неё убрать?
	}
	if choice == 2 { //Вычисление прямоугольного треугольника
		var katOne, katTwo float64
		fmt.Print("Введите длинну первого катета: ")
		fmt.Scanln(&katOne)
		fmt.Print("Введите длинну второго катета: ")
		fmt.Scanln(&katTwo)
		fmt.Printf("Длинна гипотенузы %.2f.\n", gip(katOne, katTwo))
		fmt.Printf("Периметр треугольника %.2f.\n", perim(katOne, katTwo))
		fmt.Printf("Площаь треугольника %.2f.\n", area(katOne, katTwo))
	}
	if choice == 3 { //Вклад
		var sum, proc float64
		fmt.Println("Введите сумму вклада:")
		fmt.Scanln(&sum)
		fmt.Println("Введите годовой процент:")
		fmt.Scanln(&proc)
		fmt.Printf("Сумма вашего вклада через 5 лет составит %.2f.\n", bank(sum, proc))
	}
}
