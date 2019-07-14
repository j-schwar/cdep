package kernel

import (
	"cdep/cli"
	"sort"
)

// SortDependencies sorts inbound dependencies.
//
// Since this operation requires the whole set of items, it blocks the pipeline
// until it has completed. As such it should be used sparingly.
func SortDependencies(ch chan DependencyItem) chan DependencyItem {
	out := make(chan DependencyItem, cli.ChannelBufferSize)
	go func() {
		defer close(out)

		var list []string
		for item := range ch {
			list = append(list, string(item))
		}
		sort.Strings(list)
		for _, item := range list {
			out <- DependencyItem(item)
		}
	}()
	return out
}
