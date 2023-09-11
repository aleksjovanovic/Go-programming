package main

import (
	"fmt"
	"log"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/test" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "This is test route ...")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful\n")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	fmt.Fprintf(w, "First name: %s\n", firstName)
	fmt.Fprintf(w, "Last name: %s\n", lastName)
}

func main() {
	fileServer := http.FileServer(http.Dir("./files"))
	http.Handle("/", fileServer)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Go server is running ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
