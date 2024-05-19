package main

import (
	"os"

	"github.com/ehok/kubenpod/cmd/root"
	"github.com/ehok/kubenpod/version"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		version.PrintVersion()
		return
	}

	root.Execute()
}
