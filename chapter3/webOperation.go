package chapter3

import (
	"net/http"
	"fmt"
	"html"
	"io/ioutil"
)

func init() {
	fmt.Println("=== Web Operation ====")
	listen()
}

func listen() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))



	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})*/
	http.HandleFunc("/",handler)
		http.ListenAndServe("localhost:8181", nil)



}

func handler(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}