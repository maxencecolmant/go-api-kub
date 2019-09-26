package exec

import (
	"github.com/gorilla/mux"
	"log"
	"main/pkg"
	"net/http"
	"encoding/json"
)

func GetRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", project)
	router.HandleFunc("/api/namespaces", pkg.GetNs).Methods("GET")
	router.HandleFunc("/api/pods", pkg.GetPods).Methods("GET")
	router.HandleFunc("/api/nodes", pkg.GetNodes).Methods("GET")
	router.HandleFunc("/api/services", pkg.GetServices).Methods("GET")
	router.HandleFunc("/api/namespaces/{nsname}", pkg.GetNsName).Methods("GET")
	router.HandleFunc("/api/namespaces/{nsname}", pkg.CreateNs).Methods("POST")
	router.HandleFunc("/api/namespaces/{nsname}", pkg.DeleteNs).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func project(w http.ResponseWriter, r *http.Request) {
	project := "api go kubernetes iks"
	json.NewEncoder(w).Encode(project)
}
