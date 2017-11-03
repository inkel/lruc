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

## As a Docker container
To install the latest (development) version:

```
docker pull inkel/lruc:latest
```

Then run it as:

```
docker run -P inkel/lruc:latest -code 200 -body "Lorem ipsum dolor sit amet"
```

By default it will expose the port `8080`. If you use the `-addr` you will need to tell Docker to publish a different port, for example:

```
docker run -p 1234:1234 inkel/lruc:latest -addr :1234
```

## License
See [LICENSE](LICENSE).
