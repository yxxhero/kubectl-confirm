package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/yxxhero/kubectl-confirm/pkg/kubectl"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/kubectl/pkg/cmd/util"
)

var (
	confirmExample = `
	use kubectl-confirm as kubectl
`
)

// confirmOptions provides information about the current KUBECONFIG
type confirmOptions struct {
	configFlags *genericclioptions.ConfigFlags
	rawConfig   clientcmdapi.Config
	args        []string

	genericclioptions.IOStreams
}

// NewconfirmOptions provides an instance of confirmOptions with default values
func NewconfirmOptions(streams genericclioptions.IOStreams) *confirmOptions {
	return &confirmOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
		IOStreams:   streams,
	}
}

func Run(streams genericclioptions.IOStreams) {

	var choice string
	if os.Getenv("KUBECTL_CONFIRM_DISABLE") == "" {
		opt := NewconfirmOptions(streams)
		if err := opt.Complete(); err != nil {
			util.CheckErr(err)
		}
		if err := opt.Render(); err != nil {
			util.CheckErr(err)
		}
		fmt.Printf("Continue (y/n):")
		fmt.Scanln(&choice)
		if strings.ToLower(choice) == "y" {
			kubectl.Run()
		}
		return
	}
	kubectl.Run()
}

// Complete sets all information required for updating the current context
func (o *confirmOptions) Complete() error {
	var err error
	o.rawConfig, err = o.configFlags.ToRawKubeConfigLoader().RawConfig()
	return err
}

func (o *confirmOptions) Render() error {
	var namespace string
	ctx := o.rawConfig.CurrentContext
	cluster := o.rawConfig.Clusters[ctx]
	apiServer := cluster.Server
	InsecureSkipTLSVerify := cluster.InsecureSkipTLSVerify
	authInfo := o.rawConfig.Contexts[ctx].AuthInfo
	clusterName := o.rawConfig.Contexts[ctx].Cluster
	if o.rawConfig.Contexts[ctx].Namespace == "" {
		namespace = "Unset"
	} else {
		namespace = o.rawConfig.Contexts[ctx].Namespace
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Context", "Cluster", "Namespace", "API Server", "InsecureSkipTLSVerify", "AuthInfo", "ClusterName"})
	t.AppendRow([]interface{}{ctx, clusterName, namespace, apiServer, InsecureSkipTLSVerify, authInfo, clusterName})
	t.Render()
	return nil
}
