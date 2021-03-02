package kube_connect

import (
   "metricsviews.com/metricsviews/metrics/api/dir"
   "k8s.io/client-go/kubernetes"
   metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
   "k8s.io/client-go/tools/clientcmd"
   "path/filepath"
   "flag"
   "fmt"
   "context"
   "strings"
)

func MetricsPods() {
	var kubeconfig *string
	if home := dir.UserHomeDir(); home != "" {
            kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
              kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
        flag.Parse()
        // use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
           panic(err.Error())
        }
	metrics_pods := []string{}
	for _, pod := range pods.Items {
               //fmt.Printf("pod name: %s\nnamespace: %s\n", pod.ObjectMeta.GetName(), pod.ObjectMeta.GetNamespace())
	       metrics_pods = append(metrics_pods, pod.ObjectMeta.GetName())
	}
	fmt.Printf(strings.Join(metrics_pods, ", "))
}
