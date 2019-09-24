package pkg

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"main/configkub"
	"net/http"
)

func GetNs(w http.ResponseWriter, r *http.Request) {
	api := configkub.Getconfig()
	listOptions := metav1.ListOptions{}
	ns, err := api.Namespaces().List(listOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint( w, ns )
}


