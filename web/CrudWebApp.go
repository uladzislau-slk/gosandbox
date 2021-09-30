package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type Product struct {
	Id      int
	Model   string
	Company string
	Price   int
}

var database *sql.DB

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec("DELETE FROM db.products WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := database.QueryRow("SELECT * FROM db.products WHERE id = ?", id)
	prod := Product{}
	err := row.Scan(&prod.Id, &prod.Model, &prod.Company, &prod.Price)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("web/templates/edit.html")
		tmpl.Execute(w, prod)
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	id := r.FormValue("id")
	model := r.FormValue("model")
	company := r.FormValue("company")
	price := r.FormValue("price")

	_, err = database.Exec("UPDATE db.products SET model = ?, company = ?, price = ? WHERE id = ?",
		model, company, price, id)

	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		model := r.FormValue("model")
		company := r.FormValue("company")
		price := r.FormValue("price")

		_, err = database.Exec("INSERT INTO db.products (model, company, price) VALUES (?, ?, ?)",
			model, company, price)

		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "web/templates/create.html")
	}
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	rows, err := database.Query("SELECT * FROM db.products")
	if err != nil {
		fmt.Println(err)
		log.Panicln(err)
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	tmpl, err := template.ParseFiles("web/templates/productsIndex.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(writer, products)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@/db")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/create", CreateHandler)
	router.HandleFunc("/edit/{id:[0-9]+}", EditPage).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", EditHandler).Methods("POST")
	router.HandleFunc("/delete/{id:[0-9]+}", DeleteHandler)

	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
