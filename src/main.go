package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"./listas"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func uploader(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2000)
	file, fileinfo, err := r.FormFile("archivo")
	f, err := os.OpenFile("./files/"+fileinfo.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintf(w, fileinfo.Filename)

}
func agregar(w http.ResponseWriter, r *http.Request) {
	var mens listas.Mensaje
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar mensaje")
		return
	}
	json.Unmarshal(reqBody, &mens)
	var cadena string
	for i := 0; i < len(mens.Msg); i++ {
		cadena = cadena + mens.Msg[i].To_String() + "\n"

	}
	fmt.Fprintln(w, mens.To_String())
	fmt.Fprintf(w, cadena)
	fmt.Println(mens.To_String())
	fmt.Println(cadena)

}

func main() {
	fmt.Println("un server papu")
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/agregar", agregar).Methods("POST")
	router.HandleFunc("/files", uploader)
	log.Fatal(http.ListenAndServe(":3000", router))
}
