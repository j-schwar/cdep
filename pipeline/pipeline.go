package pipeline

import (
	"cdep/cli"
	"cdep/kernel"
	"fmt"
	"os"
)

func fatalError(msg string) {
	_, _ = fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

// Run builds and executes the pipeline.
//
// The flow of the pipeline is determined by the state of various cli args.
func Run(inputs []string) {
	var pathChan chan kernel.PathItem
	if cli.UseFileRecursion {
		pathChan = kernel.RecursiveSearch(inputs)
	} else {
		pathChan = kernel.Files(inputs)
	}
	sourceChan := kernel.Open(pathChan)
	srcDepChan := kernel.FindDependencies(sourceChan)

	if cli.FindExp != "" {
		if cli.UseMerge {
			fatalError("'find' and 'm' flags are mutually exclusive")
		}
		if cli.UseCount {
			fatalError("'find' and 'c' flags are mutually exclusive")
		}

		srcDepChan = kernel.MatchDependencies(srcDepChan)
	} else if !cli.DisplayAll {
		srcDepChan = kernel.FilterSourceDependenciesRemoveStdDeps(srcDepChan)
	}

	if cli.UseMerge || cli.UseCount {
		depChan := kernel.MergeDependencies(srcDepChan)
		if cli.ShowDirectoriesOnly {
			depChan = kernel.ToDirectory(depChan)
		}
		if cli.UseCount {
			countChan := kernel.CountDependencies(depChan)
			kernel.PrintDependencyCounts(countChan)
		} else {
			depChan = kernel.RemoveDuplicateDependencies(depChan)
			depChan = kernel.SortDependencies(depChan)
			kernel.PrintDependencies(depChan)
		}
	} else {
		if cli.ShowDirectoriesOnly {
			srcDepChan = kernel.SourceToDirectory(srcDepChan)
		}
		kernel.PrintSourceDependencies(srcDepChan)
	}
}
