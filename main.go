package main

import (
	"fmt"
	"math/big"
)

var prime [99]int
var ns [540][2]int

func fillNumb() { //a
	for i := 0; i < len(ns); i++ {
		ns[i][0] = i + 2
	}
	for i := 0; i < len(ns); i++ {
		ns[i][1] = 0
	}
}

func mark(n int) {
	p := 2                                //b
	for marks := 0; marks <= n; marks++ { //e
		for p2 := p * p; p2 <= len(ns); p2 = p2 + p {
			ns[p2-2][1] = 1 //c
		}
		for i := 0; i <= 539; i++ {
			if ns[i][1] == 0 && ns[i][0] > p {
				p = ns[i][0] //d
				break
			}
		}
	}
}
func fillPrime() { //4*
	mark(100)
	i := 0
	count := 0
	for count < len(prime) {
		if ns[i][1] == 0 {
			prime[count] = ns[i][0]
			count++
		}
		i++
	}
}
func chet(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}
func chetT(a int) bool {
	if a%3 == 0 {
		return true
	}
	return false
}
func fibPrint(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(fib(i))
	}
}

/*func fibN(n int) *big.Int {
	if n < 0 {
		panic("fib n cant be negative")
	}
	a, _ := fib(n)
	return a
}*/

func fib(n int) *big.Int { //new fib func
	a := big.NewInt(1)
	b := big.NewInt(1)
	for i := 0; i < n; i++ {
		a, b = b, big.NewInt(0).Add(a, b)
	}
	return a
}

/*func fib(n int) (*big.Int, *big.Int) {
	if n == 0 {
		return big.NewInt(0), big.NewInt(1)
	}
	a, b := fib(n / 2)
	c := big.NewInt(0).Mul(a, (big.NewInt(0).Sub(big.NewInt(0).Mul(b, big.NewInt(2)), a))) // a * ((b * 2) - a)
	d := big.NewInt(0).Add(big.NewInt(0).Mul(a, a), big.NewInt(0).Mul(b, b))
	if n%2 == 0 {
		return c, d
	} else {
		return d, big.NewInt(0).Add(c, d)
	}
}*/

func main() {
	var choice int
	for choice < 1 || choice > 2 {
		fmt.Printf("1 - fib\n2 - mass\n")
		fmt.Scanln(&choice)
	}
	switch choice {
	case 1:
		fibPrint(100)
	case 2:
		fillNumb()
		fillPrime()
		fmt.Print(prime, "\n")
	}
}
