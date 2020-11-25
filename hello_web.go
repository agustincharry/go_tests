package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New Request")
	io.WriteString(w, "Hello World")
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Users []User

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{"Agustín", 25}
	json.NewEncoder(w).Encode(user)
}
func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{"Pedro", 15},
		User{"Andrea", 30},
		User{"Luisa", 21},
	}
	json.NewEncoder(w).Encode(users)
}
