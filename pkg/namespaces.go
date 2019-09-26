package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"main/configkub"
	"net/http"
)

func GetNs(w http.ResponseWriter, r *http.Request) {
	api := configkub.Getconfig()
	listOptions := metav1.ListOptions{
		TypeMeta:            metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		LabelSelector:       "",
		FieldSelector:       "",
		Watch:               false,
		AllowWatchBookmarks: false,
		ResourceVersion:     "",
		TimeoutSeconds:      nil,
		Limit:               0,
		Continue:            "",
	}
	ns, err := api.Namespaces().List(listOptions)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(ns)
}


func GetNsName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nsname := vars["nsname"]
	api := configkub.Getconfig()
	getoption := metav1.GetOptions{
		TypeMeta:        metav1.TypeMeta{},
		ResourceVersion: "",
	}
	//listOptions := m
	ns, err := api.Namespaces().Get(nsname,getoption )
	if err != nil {
		json.NewEncoder(w).Encode("This Namespace doesn't exist")
	}


	json.NewEncoder(w).Encode(ns)

}


func CreateNs(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	nsname := vars["nsname"]


	api := configkub.Getconfig()
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: nsname,
		},
		Status: v1.NamespaceStatus{
			Phase: v1.NamespaceActive,
		},
	}
	var ns,err = api.Namespaces().Create(namespace)
	if err != nil {
		fmt.Fprint(w,"The namespace ", nsname," already exist")
	}else {
		fmt.Fprint( w, "The namespace : ",ns.Name," have been created")
		json.NewEncoder(w).Encode(ns)
	}
}

func DeleteNs(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	nsname := vars["nsname"]

	api := configkub.Getconfig()
	deletePolicy := metav1.DeletePropagationForeground
	deleteOption := &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}
	var namespaceDel= api.Namespaces().Delete(nsname,deleteOption)
	if namespaceDel != nil {
		fmt.Fprint(w,"The namespace ", nsname," dont exist")
	}else {
		fmt.Fprint(w ,"The namespace ", nsname, " have been deleted")
	}

}