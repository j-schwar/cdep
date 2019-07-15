package kernel

import (
	"cdep/cli"
	"regexp"
	"sort"
	"strings"
)

// SourceDependenciesItem is a pipeline item which stores dependency information about
// a single file.
type SourceDependenciesItem struct {
	FilePath     PathItem
	Dependencies []string
}

func getIncludeValue(line string) string {
	re := regexp.MustCompile(`#include\s*[<"](.*)[>"].*`)
	match := re.FindStringSubmatch(line)
	if len(match) != 2 {
		panic("did not get expected match")
	}
	return match[1]
}

// FindDependencies processes source files and extracts their dependencies.
func FindDependencies(ch chan SourceItem) chan SourceDependenciesItem {
	out := make(chan SourceDependenciesItem, cli.ChannelBufferSize)
	go func() {
		defer close(out)

		for item := range ch {
			lines := strings.Split(item.Contents, "\n")
			var dependencies []string
			for _, line := range lines {
				noWs := strings.TrimLeft(line, " \t")
				if strings.HasPrefix(noWs, "#include") {
					val := getIncludeValue(noWs)
					dependencies = append(dependencies, val)
				}
			}

			if dependencies != nil {
				sort.Strings(dependencies)
				out <- SourceDependenciesItem{item.FilePath, dependencies}
			}
		}
	}()
	return out
}
