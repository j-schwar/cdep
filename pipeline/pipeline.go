package pipeline

import (
	"cdep/cli"
	"cdep/kernel"
)

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
	if !cli.DisplayAll {
		srcDepChan = kernel.FilterSourceDependenciesRemoveStdDeps(srcDepChan)
	}

	if cli.UseMerge || cli.UseCount {
		depChan := kernel.MergeDependencies(srcDepChan)
		if cli.UseCount {
			countChan := kernel.CountDependencies(depChan)
			kernel.PrintDependencyCounts(countChan)
		} else {
			depChan = kernel.RemoveDuplicateDependencies(depChan)
			depChan = kernel.SortDependencies(depChan)
			kernel.PrintDependencies(depChan)
		}
	} else {
		kernel.PrintSourceDependencies(srcDepChan)
	}
}
