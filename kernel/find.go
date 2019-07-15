package kernel

import (
	"cdep/cli"
	"regexp"
)

// MatchDependencies applies a regular expression filter to a source item's
// dependency list. It emits only source items who's dependencies match and
// those items will only contain the matching dependencies.
func MatchDependencies(ch chan SourceDependenciesItem) chan SourceDependenciesItem {
	out := make(chan SourceDependenciesItem, cli.ChannelBufferSize)
	go func() {
		defer close(out)

		for item := range ch {
			var matchedDeps []string
			for _, dep := range item.Dependencies {
				matched, err := regexp.MatchString(cli.FindExp, dep)
				if err != nil {
					fatalError(err)
				}
				if matched {
					matchedDeps = append(matchedDeps, dep)
				}
			}

			if matchedDeps != nil {
				out <- SourceDependenciesItem{item.FilePath, matchedDeps}
			}
		}
	}()
	return out
}
