package util

import "strings"

// getKubeconfig returns the kubeconfig file path from CLI args
// --kubeconfig
func GetKubeconfig(args []string) string {
	for i, arg := range args {
		if strings.HasPrefix(arg, "--kubeconfig") {
			if strings.Contains(arg, "=") {
				return strings.Split(arg, "=")[1]
			}
			return args[i+1]
		}
	}
	return ""
}

// IsHelp returns true if the args contains --help or -h
func IsHelp(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}
