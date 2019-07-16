package kernel

import "path/filepath"

// AsDirectory strips off the ending filename from DependencyItems.
func ToDirectory(ch chan DependencyItem) chan DependencyItem {
	out := make(chan DependencyItem)
	go func() {
		defer close(out)

		for item := range ch {
			out <- DependencyItem(filepath.Dir(string(item)))
		}
	}()
	return out
}

// SourceToDirectory strips off the ending filenames from a source's dependencies.
func SourceToDirectory(ch chan SourceDependenciesItem) chan SourceDependenciesItem {
	out := make(chan SourceDependenciesItem)
	go func() {
		defer close(out)

		for item := range ch {
			var deps []string
			for _, dep := range item.Dependencies {
				deps = append(deps, filepath.Dir(dep))
			}
			out <- SourceDependenciesItem{item.FilePath, deps}
		}
	}()
	return out
}
