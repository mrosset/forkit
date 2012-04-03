package main

import (
	"os"
	"testing"
)

var arg0 = os.Args[0]

func TestMain(t *testing.T) {
	os.Args = append(os.Args, "status")
	main()
}
