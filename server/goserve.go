package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

type PageVariables struct {
	PageTitle string
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")

	fmt.Println("Server deployed")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {

	Title := "SAMPLE PAGE"

	MyPageVariables := PageVariables{Title}

	parsePage("Index.html",MyPageVariables,w)

	fmt.Println("Get request Recieved")
}

func parsePage(htmlFile string,MyPageVariables PageVariables, w http.ResponseWriter) {
	t, err := template.ParseFiles(htmlFile) //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
