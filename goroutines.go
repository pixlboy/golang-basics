// Go routines, channels help go routines send and receive values

package main

import ( 
	"fmt"
	"net/http"
)

func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("error: %s\n", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// Create a response channel
	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			returnType(url, ch)
		}(url)
	}

	for range urls {
		out := <-ch
		fmt.Println(out)
	}

}