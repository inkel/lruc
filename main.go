package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var (
		code = flag.Int("code", http.StatusOK, "HTTP response code")
		ct   = flag.String("content-type", "text/plain", "Content-Type")
		body = flag.String("body", "Hello, World!", "Response body. Use `-` to read from stdin")
		addr = flag.String("addr", ":8080", "Address to listen for requests")
	)
	flag.Parse()

	bodyBytes := []byte(*body)

	if *body == "-" {
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, os.Stdin); err != nil {
			fmt.Printf("Cannot read from stdin: %v\n", err)
			os.Exit(1)
		}
		bodyBytes = buf.Bytes()
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(*code)
		w.Header().Set("Content-Type", *ct)
		w.Write(bodyBytes)
	})

	fmt.Printf("Starting lruc on %v\n", *addr)
	fmt.Println(http.ListenAndServe(*addr, nil))
}
