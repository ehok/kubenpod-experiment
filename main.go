package main

import (
	"fmt"
	"os"

	"github.com/ehok/kubenpod/cmd/root"
	"github.com/ehok/kubenpod/version"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		v := version.Get()
		fmt.Printf("Version: %s\nGit commit: %s\nBuild date: %s\n", v.GitVersion, v.GitCommit, v.BuildDate)
		fmt.Printf("Version: %s\n", v.GitVersion)
		fmt.Printf("Git commit: %s\n", v.GitCommit)
		fmt.Printf("Build date: %s\n", v.BuildDate)
		return
	}

	root.Execute()
}
