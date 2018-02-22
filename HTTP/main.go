package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
	// bs := make([]byte, 9999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWritter{}

	io.Copy(lw, resp.Body)

}

func (logWritter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Println("Just wrote a bunch of bytes", len(b))
	return len(b), nil
}
