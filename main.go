package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	{ //file reader
		bs, err := ioutil.ReadFile("filereadshort.go")
		if err != nil {
			fmt.Println(err) //стоит добавить вывод ошибки в консоль.
			return
		}
		fmt.Println(string(bs))
	}

	{ //csv encoding
		in := `name,Sname,age
Vasya,Lojkin,12
Lojka,Vasin,21
Harry,Potter,31`
		r := csv.NewReader(strings.NewReader(in))

		for {
			record, err := r.Read()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(record)
		}
	}
}
