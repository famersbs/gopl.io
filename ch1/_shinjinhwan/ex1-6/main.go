package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.Black,
							color.RGBA{0x00, 0xff, 0x00,0xff},
							color.RGBA{0xff, 0x00, 0x00,0xff},
							color.RGBA{0x00, 0xff, 0xff,0xff},
}

// const whiteIndex = 0
const blackIndex = 0
const greenIndex = 1

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)

		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycle  = 5
		res    = 0.001
		size   = 100
		nframe = 64
		delay  = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframe}
	phase := 0.0
	for i := 0; i < nframe; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8((i % 4)+1))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
