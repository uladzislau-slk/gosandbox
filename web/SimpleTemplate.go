package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Title string
	Users []User
}

type User struct {
	Name string
	Age  int
}

func main() {

	data := ViewData{
		Title: "Users list",
		Users: []User{
			{Name: "Tom", Age: 21},
			{Name: "Kate", Age: 23},
			{Name: "Alice", Age: 25},
		},
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		tmpl, err := template.ParseFiles("web/templates/index.html")

		if err != nil {
			fmt.Println(err)
			return
		}

		tmpl.Execute(writer, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
