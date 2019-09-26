package pkg

import (
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
			Kind:       "spec",
			APIVersion: "v1",
		},
		LabelSelector:       "name",
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

	fmt.Fprint( w, ns )
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
		fmt.Fprint(w,"This namespace doesn't exist")
	}


	fmt.Fprint( w, ns)

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
	var ns, _ = api.Namespaces().Create(namespace)


//rajouter execptions

	fmt.Fprint( w, "le namespace : ",ns," a bien été crée")
}

func DeleteNs(w http.ResponseWriter, r *http.Request)  {

}