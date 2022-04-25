package main

import (
	"os"

	"github.com/yxxhero/kubectl-confirm/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	cmd.Run(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
}
