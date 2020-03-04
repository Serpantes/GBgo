package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	io.WriteString(res,
		req.URL.Query().Get("name"))
}

func main() {
	choice := 0
	for choice > 2 || choice < 1 {
		fmt.Scanln(&choice)
	}
	switch choice {
	case 1:
		fs := http.FileServer(http.Dir("static"))
		http.Handle("/", fs)
		http.HandleFunc("/hello", hello)
		http.ListenAndServe(":80", nil)
	case 2:
		black := color.RGBA{0, 0, 0, 255}
		white := color.RGBA{255, 255, 255, 255}
		size := 8
		rectImg := image.NewRGBA(image.Rect(0, 0, size, size))
		draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
		for x := 0; x <= size; x++ {
			for y := 0; y < size; y++ {
				if (x+y)%2 == 0 {
					rectImg.Set(x, y, white)
				}
			}
		}

		file, err := os.Create("pic.png")
		if err != nil {
			log.Fatalf("Failed create file: %s", err)
		}
		defer file.Close()
		png.Encode(file, rectImg)
	}
}
