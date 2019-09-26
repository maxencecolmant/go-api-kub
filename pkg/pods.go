package pkg

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"main/configkub"
	"net/http"
)

func GetPods(w http.ResponseWriter, r *http.Request) {
	api := configkub.Getconfig()
	listOptions := metav1.ListOptions{}
	pods, err := api.Pods("").List(listOptions)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(pods.Items)



}


