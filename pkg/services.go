package pkg

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"main/configkub"
	"net/http"
	_"encoding/json"

)

func GetServices(w http.ResponseWriter, r *http.Request)  {
	api := configkub.Getconfig()
	listOptions := metav1.ListOptions{
		TypeMeta:            metav1.TypeMeta{},
		LabelSelector:       "name",
		FieldSelector:       "",
		Watch:               false,
		AllowWatchBookmarks: false,
		ResourceVersion:     "",
		TimeoutSeconds:      nil,
		Limit:               0,
		Continue:            "",
	}
	services, err := api.Services("").List(listOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, services.Items)

}