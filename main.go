package main

import (
	_ "context"
	"fmt"
	_ "github.com/bndr/gopencils"
	_ "github.com/ericchiang/k8s"
	_ "github.com/ericchiang/k8s/apis/core/v1"
	"github.com/gorilla/mux"
	_ "io/ioutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"

	//"k8s.io/client-go/1.4/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/rest"
	"log"
	"net/http"
)
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", GetNs)
	log.Fatal(http.ListenAndServe(":8080", router))

	}

func GetNs(w http.ResponseWriter, r *http.Request) {
	//response, _ := http.Get("http://localhost:8001/api/v1/namespaces")
	//responseData, _ := ioutil.ReadAll(response.Body)
	//var ns =string(responseData)
	//fmt.Fprint( w,ns )
	//log.Printf("GET /")

	//api := gopencils.Api("http://localhost:8001/api/v1/namespaces")
	//api.Res("namespaces").Get("name")

	// uses the current context in kubeconfig
	kubeconfig := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)
	config, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)
	//config, _ := rest.InClusterConfig()
	// creates the clientset
	clientset, _:= kubernetes.NewForConfig(config)
	// access the API to list pods
	pods, _:= clientset.CoreV1().Pods("").List(v1.ListOptions{})

		fmt.Fprint(w, pods)
	}
