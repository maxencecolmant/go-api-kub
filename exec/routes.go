package exec

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"main/pkg"
	"net/http"
)

func GetRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", project)
	router.HandleFunc("/api/namespaces", pkg.GetNs)
	router.HandleFunc("/api/pods", pkg.GetPods)
	router.HandleFunc("/api/nodes", pkg.GetNodes)
	router.HandleFunc("/api/services", pkg.GetServices)
	router.HandleFunc("/api/namespaces/{nsname}", pkg.GetNsName)
	router.HandleFunc("/api/namespaces/create/{nsname}", pkg.CreateNs)
	router.HandleFunc("/api/namespaces/delete/{nsname}", pkg.DeleteNs)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func project(w http.ResponseWriter, r *http.Request) {
	project := "api go kubernetes iks"
	fmt.Fprint(w, project)
}
