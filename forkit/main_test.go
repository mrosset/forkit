package main

import (
	"os"
	"testing"
)

var arg0 = os.Args[0]

func resetArgs() {
	os.Args = append([]string{}, arg0)
}

func TestClone(t *testing.T) {
	resetArgs()
	os.Args = append(os.Args, "clone")
	main()
}

func TestStatus(t *testing.T) {
	resetArgs()
	os.Args = append(os.Args, "status")
	main()
}
