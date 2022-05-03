package main

import (
	"os"

	"github.com/yxxhero/kubectl-confirm/cmd"
	"github.com/yxxhero/kubectl-confirm/pkg/util"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	kubeconfig := util.GetKubeconfig(os.Args[1:])
	if kubeconfig != "" {
		cmd.Kubeconfig = kubeconfig
	}
	cmd.Run(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
}
