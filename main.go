package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"main/pkg"
	"net/http"
)
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", project)
	router.HandleFunc("/namespaces", pkg.GetNs)
	router.HandleFunc("/pods", pkg.GetPods)
	router.HandleFunc("/nodes", pkg.GetNodes)
	router.HandleFunc("/services", pkg.GetServices)
	router.HandleFunc("/namespaces/{nsname}", pkg.GetNsName)
	log.Fatal(http.ListenAndServe(":8080", router))

	}

func project(w http.ResponseWriter, r *http.Request){
	project := "api go kubernetes iks"
	fmt.Fprint( w, project )


}


