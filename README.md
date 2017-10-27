# lruc - reverse cURL
See [this Twitter thread](https://twitter.com/thorstenball/status/923890186479656960). I plan to update the docs later.

## Usage
```
Usage of lruc:
  -addr string
        Address to listen for requests (default ":8080")
  -body -
        Response body. Use - to read from a stdin (default "Hello, World!")
  -code int
        HTTP response code (default 200)
  -content-type string
        Content-Type (default "text/plain")
  -header
        HTTP response headers. Zero, one or more are accepted. Eg `-header "Access-Control-Allow-Origin: *" -header "Access-Control-Allow-Methods: POST, GET, OPTIONS"`
```

## License
See [LICENSE](LICENSE).
