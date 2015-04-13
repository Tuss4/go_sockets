package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(
		res,
		`<doctype html>
        <html>
        <head>
            <title>Hello World</title>
        </head>
        <body>
        Hello world, bruh!
        </body>
        </html>`)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9000", nil)
}
