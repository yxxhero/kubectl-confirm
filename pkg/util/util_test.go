package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGetKubeconfig tests the GetKubeconfig function
func TestGetKubeconfig(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"--kubeconfig", "test1"}, "test1"},
		{[]string{"--kubeconfig=test2"}, "test2"},
		{[]string{"a"}, ""},
	}

	for _, test := range tests {
		got := GetKubeconfig(test.args)
		require.Equal(t, test.want, got, "GetKubeconfig(%v) = %v, want %v", test.args, got, test.want)
	}

}

// TestIsHelp tests the IsHelp function
func TestIsHelp(t *testing.T) {
	tests := []struct {
		args []string
		want bool
	}{
		{[]string{"--help"}, true},
		{[]string{"-h"}, true},
		{[]string{"a"}, false},
	}

	for _, test := range tests {
		got := IsHelp(test.args)
		require.Equal(t, test.want, got, "IsHelp(%v) = %v, want %v", test.args, got, test.want)
	}

}
