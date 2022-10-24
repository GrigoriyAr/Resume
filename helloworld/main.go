package main

import (

	"fmt"
	"helloworld/datafile"
	"html/template"
	"log"
	"net/http"
	"os"
)



type Guesstbook struct {       // Счетчик отзывов
	SignatureCount int 
	Signature [] string
}

func check (err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html") 
	check(err)
	err = html.Execute(writer, nil)
	check(err)

}

func createHandler(writer http.ResponseWriter, request *http.Request) {

	signature := request.FormValue("signature")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(4000))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/resume", http.StatusFound)

}
 


func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signature, err := datafile.GetStrings("signatures.txt")
	check(err)
	html, err := template.ParseFiles("view.html") 
	check(err)
	

	guestbook := Guesstbook {
		SignatureCount: len(signature),
		Signature: signature,	
	}

	err = html.Execute(writer, guestbook)
	check(err)	
}



func main() {
	
	http.HandleFunc("/resume", viewHandler)
	http.HandleFunc("/resume/new", newHandler)
	http.HandleFunc("/resume/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)	
	log.Fatal(err)
	
	

}