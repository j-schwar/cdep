package kernel

import (
	"cdep/cli"
	"sort"
	"strings"
)

type DependencyCountItem struct {
	Dependency string
	Count      int
}

// CountDependencies counts the number of references for each dependency.
//
// This operation will block the pipeline as it requires all of its input before
// it can start producing output. To take advantage of this, this operation also
// sorts the dependencies as sorting is also a blocking operation.
func CountDependencies(ch chan DependencyItem) chan DependencyCountItem {
	out := make(chan DependencyCountItem, cli.ChannelBufferSize)
	go func() {
		defer close(out)

		counter := map[string]int{}
		for item := range ch {
			dep := string(item)
			count, ok := counter[dep]
			if ok {
				counter[dep] = count + 1
			} else {
				counter[dep] = 1
			}
		}

		var items []DependencyCountItem
		for dep, count := range counter {
			items = append(items, DependencyCountItem{dep, count})
		}
		sort.Slice(items, func(i, j int) bool {
			return strings.Compare(items[i].Dependency, items[j].Dependency) < 1
		})
		for _, item := range items {
			out <- item
		}
	}()
	return out
}
