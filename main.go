package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var src = flag.String("from", "", "From")
var dst = flag.String("to", "", "To")
var file = flag.String("file", "", "File")
var myGrep = flag.String("mygrep", "", "Mygrep")

func main() {
	choice := 0
	for choice < 1 || choice > 4 {
		fmt.Scanln(&choice)
	}
	switch choice {
	case 1:
		{ //file reader
			bs, err := ioutil.ReadFile("filereadshort.go")
			if err != nil {
				fmt.Println(err) //стоит добавить вывод ошибки в консоль.
				return
			}
			fmt.Println(string(bs))
		}
	case 2:
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
	case 3:
		{ // file copy
			flag.Parse()

			f1, errR := os.OpenFile(*src, os.O_RDONLY, 0644)
			if errR != nil {
				fmt.Println("READ:", errR)
			}
			f2, errWr := os.OpenFile(*dst, os.O_WRONLY, 0644)
			if errR != nil {
				fmt.Println(errWr)
			}
			io.Copy(f2, f1)
		}
	case 4:
		{ //goGrep
			flag.Parse()

			content, err := ioutil.ReadFile(*file)
			if err != nil {
				log.Fatal(err)
			}
			fileText := string(content)
			fmt.Println(strings.Contains(fileText, *myGrep))
		}
	}
}
