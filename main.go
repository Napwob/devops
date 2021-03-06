package main
 
import (
"fmt"
"net/http"
)
 
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greetings, user!")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, user!)")
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "PONG!!!")
	})
	http.HandleFunc("/cat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WHERE?!?!?!?!?")
	})
	http.ListenAndServe(":8098", nil)
}
