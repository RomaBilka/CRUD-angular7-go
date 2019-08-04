package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type HotdogJson struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Hotdog struct {
	id          int
	name        string
	description string
}

type HotdogNew struct {
	Id          int
	Name        string
	Description string
}

var databse string = "database.db"

func indexHendler(w http.ResponseWriter, r *http.Request) {

	var HotdogsJson []HotdogJson

	db, err := sql.Open("sqlite3", databse)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	rows, err := db.Query("SELECT id, name, description FROM hotdogs")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		h := Hotdog{}
		err := rows.Scan(&h.id, &h.name, &h.description)
		if err != nil {
			fmt.Println(err)
			continue
		}
		emp := HotdogJson{Id: h.id, Name: h.name, Description: h.description}
		HotdogsJson = append(HotdogsJson, emp)

	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(HotdogsJson)

}

func createHendler(w http.ResponseWriter, r *http.Request) {

	hotdog := getHotdogRequestData(r)
	db, err := sql.Open("sqlite3", databse)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("INSERT INTO hotdogs (name, description) values ($1,$2)", hotdog.Name, hotdog.Description)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
}

func showHendler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	db, err := sql.Open("sqlite3", databse)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	row := db.QueryRow("SELECT * FROM hotdogs WHERE id=$1", params["id"])
	h := Hotdog{}
	err = row.Scan(&h.id, &h.name, &h.description)
	if err != nil {
		fmt.Println(err)
	}
	emp := HotdogJson{Id: h.id, Name: h.name, Description: h.description}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(emp)

}

func updateHendler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hotdog := getHotdogRequestData(r)
	db, err := sql.Open("sqlite3", databse)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("UPDATE hotdogs SET name = $1, description =$2 WHERE id = $3", hotdog.Name, hotdog.Description, params["id"])

	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
}
func getHotdogRequestData(r *http.Request) HotdogNew {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	data := body
	var hotdog HotdogNew
	json.Unmarshal(data, &hotdog)

	return hotdog
}
func deleteHendler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	db, err := sql.Open("sqlite3", databse)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("DELETE FROM hotdogs WHERE id = $1", params["id"])

	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
}

func main() {
	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	r.HandleFunc("/", indexHendler).Methods("GET")
	r.HandleFunc("/", createHendler).Methods("POST")
	r.HandleFunc("/{id}", showHendler).Methods("GET")
	r.HandleFunc("/{id}", updateHendler).Methods("PUT")
	r.HandleFunc("/{id}", deleteHendler).Methods("DELETE")
	http.ListenAndServe(":3000", r)
}
