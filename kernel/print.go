package kernel

import (
	"cdep/cli"
	"fmt"
	"os"
)

func getOutputFile() *os.File {
	var out *os.File
	if cli.Outfile == "" {
		out = os.Stdout
	} else {
		var err error
		out, err = os.Create(cli.Outfile)
		if err != nil {
			fatalError(err)
		}
	}
	return out
}

func closeOutputFile(out *os.File) {
	if out != os.Stdout {
		_ = out.Close()
	}
}

// PrintSourceDependencies writes dependency information to stdout or an output file
// if desired.
func PrintSourceDependencies(ch chan SourceDependenciesItem) {
	out := getOutputFile()
	for item := range ch {
		_, err := fmt.Fprintf(out, "%s:\n", item.FilePath)
		if err != nil {
			fatalError(err)
		}
		for _, dep := range item.Dependencies {
			_, err := fmt.Fprintf(out, "    %s\n", dep)
			if err != nil {
				fatalError(err)
			}
		}
		_, _ = fmt.Fprintln(out)
	}
	closeOutputFile(out)
}

// PrintDependencies writes path information to stdout or an output file if desired.
func PrintDependencies(ch chan DependencyItem) {
	out := getOutputFile()
	for item := range ch {
		_, err := fmt.Fprintln(out, string(item))
		if err != nil {
			fatalError(err)
		}
	}
	closeOutputFile(out)
}

// PrintDependencyCounts writes dependency count information to stdout or an output file.
func PrintDependencyCounts(ch chan DependencyCountItem) {
	out := getOutputFile()
	for item := range ch {
		_, err := fmt.Fprintf(out, "%s: %d\n", item.Dependency, item.Count)
		if err != nil {
			fatalError(err)
		}
	}
	closeOutputFile(out)
}
