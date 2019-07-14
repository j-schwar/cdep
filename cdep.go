package main

import (
	"cdep/cli"
	"cdep/pipeline"
	"fmt"
	"os"
)

const version = "0.1.0"

func main() {
	cli.ParseCommandLineArguments()
	if cli.ShowVersion {
		fmt.Println(version)
		return
	}
	if len(cli.InputFiles) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "cdep: at least one input file is required")
		os.Exit(1)
	}

	pipeline.Run(cli.InputFiles)
}
