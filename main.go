package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"main/pkg"
)
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", project)
	router.HandleFunc("/namespaces", pkg.GetNs)
	router.HandleFunc("/pods", pkg.GetPods)
	log.Fatal(http.ListenAndServe(":8080", router))

	}

func project(w http.ResponseWriter, r *http.Request){
	project := "api go kubernetes iks"
	fmt.Fprint( w, project )


}
