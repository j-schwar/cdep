package kernel

import "cdep/cli"

type DependencyItem string

// MergeDependencies combines the dependencies from all dependency items.
func MergeDependencies(ch chan SourceDependenciesItem) chan DependencyItem {
	out := make(chan DependencyItem, cli.ChannelBufferSize)
	go func() {
		defer close(out)

		for item := range ch {
			for _, dep := range item.Dependencies {
				out <- DependencyItem(dep)
			}
		}
	}()
	return out
}
