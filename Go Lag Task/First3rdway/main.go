package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Person struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var persons = []Person{

	{Id: "1", FirstName: "Issac", LastName: "N"},
	{Id: "2", FirstName: "Albert", LastName: "E"},
	{Id: "3", FirstName: "Thomas", LastName: "E"},
}

func main() {
	log.Println("starting API server")
	router := mux.NewRouter()
	log.Println("creating routes")
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/persons", Persons).Methods("GET")
	router.HandleFunc("/persons/{id}", Personid).Methods("GET")
	router.HandleFunc("/person", Personquery).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":1000", router)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
	w.Header().Set("Content-Type", "application/json")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(persons)
	if err != nil {
		log.Fatalf("Json error")
	}
	w.Write(resp)
}
func Personid(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	for _, e := range persons {
		if e.Id == vars["id"] {
			json.NewEncoder(w).Encode(e)
			return
		}
	}
	json.NewEncoder(w).Encode("No Data found")
	return
}

func Personquery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	values := r.URL.Query()
	foo := values["Name"]
	name := strings.Join(foo, ",")
	for _, e := range persons {
		if e.FirstName == name {
			json.NewEncoder(w).Encode(e)
			return
		}
	}
	fmt.Fprintf(w, "Invaild data")
}
