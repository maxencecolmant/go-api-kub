package configkub
import(

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"

)

func Getconfig() v1.CoreV1Interface {
	kubeConfigPath := "/Users/maxence/.bluemix/plugins/container-service/clusters/bm12i98f0vornrm8jes0/kube-config-mil01-mycluster.yml"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	api := clientset.CoreV1()
	return api
}

