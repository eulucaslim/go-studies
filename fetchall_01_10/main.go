package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	start := time.Now()
	ch := make(chan string)

	for i, url := range os.Args[1:] {
		go fetch(url, ch, i) // inicia uma gorrotina
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // recebe do canal ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, counter int) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // envia par o canal ch
		return
	}
	filename := fmt.Sprintf("output_%d.txt", counter)

	file, err := os.Create(filename)
	check(err)
	defer file.Close()

	bodyBytes, err := io.Copy(file, resp.Body)
	defer resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %d %s", secs, bodyBytes, url)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
