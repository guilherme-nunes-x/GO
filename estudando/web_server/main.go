package main

import (
	"fmt"
	//importe a biblioteca net/http para iniciar seu servidor web
	"net/http"
)

// Isso é um test inicial com uma função de hello world
func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World, my name is guilherme")

}

func main() {
	http.HandleFunc("/", helloWorldPage)
	http.ListenAndServe("", nil)
}
