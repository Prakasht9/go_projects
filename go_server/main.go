package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", formHello)

	fmt.Printf("Server started at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParsedForm() err: %v", err)
	}

	fmt.Fprint(w, "Post Request SuccessFul\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprint(w, "Name: ", name, "\n")
	fmt.Fprint(w, "Address: ", address, "\n")

}

func formHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Hello")

}
