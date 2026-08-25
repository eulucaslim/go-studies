package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, resp.Body)
		statusCode := resp.Status
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "copy error: %v", err)
		}

		fmt.Printf("Body: %v, Status Code: %v", b, statusCode)

	}
}
