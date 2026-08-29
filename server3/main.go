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
	"strconv"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "%s %s %s", r.Method, r.URL, r.Proto)
//
//	for k, v := range r.Header {
//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
//	}
//
//	fmt.Fprintf(w, "Host = %q\n", r.Host)
//	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
//
//	if err := r.ParseForm(); err != nil {
//		log.Print(err)
//	}
//
//	for k, v := range r.Form {
//		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
//	}
//}

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0, 0xff, 0, 0xff},           // Verde
	color.RGBA{R: 255, G: 0, B: 0, A: 255}, // Index 0: Vermelho
	color.RGBA{R: 0, G: 0, B: 255, A: 255}, // Index 2: Azul
	color.RGBA{R: 255, G: 255, B: 255, A: 255},
}

type Config struct {
	cycles  int
	res     float64
	size    int
	nFrames int
	delay   int
}

func lissajous(out io.Writer, config Config) {

	freq := rand.Float64() * 3.0 // frequencia relativa do oscilador y
	anim := gif.GIF{LoopCount: config.nFrames}
	phase := 0.0 // diferença de fase

	for i := 0; i < config.nFrames; i++ {
		rect := image.Rect(0, 0, 2*config.size+1, 2*config.size+1)
		img := image.NewPaletted(rect, palette)
		var differentColor = rand.Intn(len(palette)) + 1
		for t := 0.0; t < float64(config.cycles)*2*math.Pi; t += config.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				config.size+int(
					x*float64(config.size)+0.5),
				config.size+int(
					y*float64(config.size)+0.5),
				uint8(differentColor))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, config.delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTA: ignorando erros de codificação
}

func handler(w http.ResponseWriter, r *http.Request) {

	config := Config{
		5,     // numero de revoluções completas do oscilador x
		0.001, // resolução angular
		100,   // canvas da imagem cobre de [size..+size]
		64,    // número de quadros da animação
		8,     // tempo entre quadros em unidades de 10ms
	}

	queryParams := r.URL.Query()

	if queryParams.Has("cycles") {
		config.cycles, _ = strconv.Atoi(queryParams.Get("cycles"))
	}

	if queryParams.Has("res") {
		config.res, _ = strconv.ParseFloat(queryParams.Get("res"), 64)
	}

	if queryParams.Has("size") {
		config.size, _ = strconv.Atoi(queryParams.Get("size"))
	}

	if queryParams.Has("nFrames") {
		config.nFrames, _ = strconv.Atoi(queryParams.Get("nFrames"))
	}

	if queryParams.Has("delay") {
		config.delay, _ = strconv.Atoi(queryParams.Get("delay"))
	}

	lissajous(w, config)
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9001", nil))
}
