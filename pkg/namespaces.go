package pkg

import (
	"fmt"
	"github.com/gorilla/mux"
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
	//listOptions := metav1.ListOptions{}
	ns, err := api.Namespaces().Get(nsname,getoption )
	if err != nil {
		fmt.Fprint(w,"This namespace doesn't exist")
	}


	fmt.Fprint( w, ns)

}

