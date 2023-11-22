package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"os"
	"strings"
)

func main() {
	var (
		code = flag.Int("code", http.StatusOK, "HTTP response code")
		ct   = flag.String("content-type", "text/plain", "Content-Type")
		body = flag.String("body", "Hello, World!", "Response body. Use `-` to read from stdin")
		addr = flag.String("addr", ":8080", "Address to listen for requests")
	)

	headers := textproto.MIMEHeader{}
	flag.Func("header", "HTTP response headers. Zero, one or more are accepted", func(value string) error {
		v := strings.SplitN(value, ":", 2)
		if len(v) != 2 {
			return fmt.Errorf("header format must be key:value, got %s", value)
		}
		headers.Add(v[0], v[1])
		return nil
	})
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

		for key, values := range headers {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		w.Header().Set("Content-Type", *ct)
		w.WriteHeader(*code)
		w.Write(bodyBytes)
	})

	fmt.Printf("Starting lruc on %v\n", *addr)
	fmt.Println(http.ListenAndServe(*addr, nil))
}
