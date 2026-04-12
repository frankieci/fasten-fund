package main

/**
 * Created by frankieci on 2022/1/7 8:58 pm
 */

import (
	"bytes"
	"io"
	"strings"

	// "strings"

	// "context"
	"flag"
	"fmt"
	"os"

	// "path/filepath"
	"text/template"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	// "k8s.io/apimachinery/pkg/api/errors"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/util/homedir"
	// "k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/tools/clientcmd"
)

func main() {
	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }

	var configfile string
	flag.StringVar(&configfile, "config", "", "Config file wtih custom parameters")

	flag.Parse()

	// // use the current context in kubeconfig
	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// // create the clientset
	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// namespace := "default"
	// s, err := clientset.CoreV1().Secrets("default").Get(context.TODO(), "env", metav1.GetOptions{})
	// _ = s
	// if errors.IsNotFound(err) {
	// 	fmt.Printf("ConfigMap %s in namespace %s not found\n", "env", namespace)
	// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 	fmt.Printf("Error getting ConfigMap %s in namespace %s: %v\n",
	// 		"env", namespace, statusError.ErrStatus.Message)
	// } else if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Printf("Found ConfigMap %s in namespace %s\n", "env", namespace)
	// }

	// cm, _ := clientset.CoreV1().ConfigMaps("default").Get(context.TODO(), "env", metav1.GetOptions{})
	// _ = cm

	valuesMap := map[string]interface{}{
		"kubeVersion":    "v1.20.4",
		"etcdVersion":    "v3.4.0",
		"pauseVersion":   "v1.1.13",
		"coreDNSVersion": "v1.18.0",
		"base":           "YWRtaW4=",
		"node":           "127.0.0.1",
	}

	customMap := map[string]interface{}{}
	valuesMapping := map[string]interface{}{
		"Values": valuesMap,
		"Custom": customMap,
	}

	configfile = os.Args[1]
	if configfile != "" {
		// f, _ := os.OpenFile(configfile, os.O_CREATE|os.O_RDWR, 0644)
		f, _ := os.OpenFile(configfile, os.O_RDONLY, 0644)
		configm, _ := io.ReadAll(f)

		var output2 bytes.Buffer
		t2, _ := template.New("Custom").Parse(string(configm))
		t2.Execute(&output2, valuesMapping)
		logrus.Infoln(fmt.Sprintf("output2: %s", output2.String()))

		v := viper.New()
		v.SetConfigType("toml")
		if err := v.ReadConfig(bytes.NewBuffer(output2.Bytes())); err != nil {
			logrus.Errorln(err)
		}

		if err := v.Unmarshal(&customMap); err != nil {
			logrus.Errorln(err)
		}

	}

	if (len(os.Args)) > 1 {

		for i, v := range os.Args[1:] {
			logrus.Infoln("args ", i, " :  ", v)

			var output bytes.Buffer
			t, err := template.ParseFiles(v)
			if err != nil {
				logrus.Errorln(err)
			}
			// t.Funcs()

			if err := t.Execute(&output, valuesMapping); err != nil {
				logrus.Errorln(err)
			}
			logrus.Infoln(fmt.Sprintf("output %s: %s", t.Name(), output.String()))
			logrus.Infoln("-----", strings.TrimSuffix(t.Name(), ".j2"))
			os.WriteFile(strings.TrimSuffix(t.Name(), ".j2"), output.Bytes(), 0644)
		}
	}
	//Couldn't start dlv dap: Error: timed out while waiting for DAP server to start

}
