package pkg

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"main/configkub"
	"net/http"
)

func GetNodes(w http.ResponseWriter, r *http.Request) {
	api := configkub.Getconfig()
	listOptions := metav1.ListOptions{}
	nodes, err := api.Nodes().List(listOptions)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(nodes.Items)
	
}
