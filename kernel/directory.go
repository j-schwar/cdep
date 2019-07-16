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
			// Since we don't have a remove duplicate kernel on this path, we will
			// remove duplicate entries on a source by source basis (i.e., any
			// duplicate directories within the same source file will be removed).
			memoizer := map[string]bool{}

			var deps []string
			for _, dep := range item.Dependencies {
				dir := filepath.Dir(dep)
				if !memoizer[dir] {
					deps = append(deps, dir)
					memoizer[dir] = true
				}
			}
			out <- SourceDependenciesItem{item.FilePath, deps}
		}
	}()
	return out
}
