package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type stringArray []string

func (i *stringArray) String() string {
	return strings.Join(*i, ",")
}

func (i *stringArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var headerFlags stringArray
	var (
		code = flag.Int("code", http.StatusOK, "HTTP response code")
		ct   = flag.String("content-type", "text/plain", "Content-Type")
		body = flag.String("body", "Hello, World!", "Response body. Use `-` to read from stdin")
		addr = flag.String("addr", ":8080", "Address to listen for requests")
	)
	flag.Var(&headerFlags, "header", "HTTP response headers. Zero, one or more are accepted")
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

		for _, h := range headerFlags {
			ha := strings.Split(h, ": ")
			if i := w.Header().Get(ha[0]); i != "" {
				w.Header().Add(ha[0], ha[1])
			} else {
				w.Header().Set(ha[0], ha[1])
			}
		}
		w.Header().Set("Content-Type", *ct)
		w.WriteHeader(*code)
		w.Write(bodyBytes)
	})

	fmt.Printf("Starting lruc on %v\n", *addr)
	fmt.Println(http.ListenAndServe(*addr, nil))
}
