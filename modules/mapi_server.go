package mapi_server
import (
     "k8s.io/client-go/tools/clientcmd"
     "metricstream.com/metricstream/modules"
     "flag"
     "path/filepath"
     "log"

)

func MapiServer() {

      dir := mdir.UserHomeDir()
      mconfig_server :=  filepath.Join(dir, ".kube", "config")
      log.Println("Using kubeconfig file: ", mconfig_server)
      config, err := clientcmd.BuildConfigFromFlags("", mconfig_server)
      if err != nil {
            log.Fatal(err)
      }
}
