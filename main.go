package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func httpclient() {
	resp, err := http.Get("https://tubesave.vercel.app")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status", resp.Status)
	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println("hello worlds")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%\v", "%\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)

}
