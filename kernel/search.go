package kernel

import (
	"cdep/cli"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// PathItem is a pipeline item which denotes a filepath.
type PathItem string

// Files selects only the files from an input set.
//
// To allow wildcards to work easily, Files will ignore any directory present
// within the input set.
func Files(inputs []string) chan PathItem {
	out := make(chan PathItem, cli.ChannelBufferSize)
	go func() {
		defer close(out)

		for _, path := range inputs {
			stat, err := os.Stat(path)
			if err != nil {
				if os.IsNotExist(err) {
					fatalError(fmt.Sprintf("%s: No such file or directory\n", path))
				} else {
					fatalError(err)
				}
			}
			if stat.IsDir() && cli.Verbose {
				_, _ = fmt.Fprintf(os.Stderr, "%s: Is a directory\n", path)
			} else if stat.Mode().IsRegular() {
				out <- PathItem(path)
			} else if cli.Verbose {
				_, _ = fmt.Fprintf(os.Stderr, "%s: Skipping non-regular file\n", path)
			}
		}
	}()
	return out
}

func recursiveSearch(paths []string, out chan PathItem) {
	for _, path := range paths {
		stat, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				fatalError(fmt.Sprintf("%s: No such file or directory\n", path))
			} else {
				fatalError(err)
			}
		}
		if stat.IsDir() {
			dirContents, err := ioutil.ReadDir(path)
			if err != nil {
				fatalError(err)
			}
			var recPaths []string
			for _, e := range dirContents {
				recPaths = append(recPaths, filepath.Join(path, e.Name()))
			}
			recursiveSearch(recPaths, out)
		} else if stat.Mode().IsRegular() {
			out <- PathItem(path)
		} else if cli.Verbose {
			_, _ = fmt.Fprintf(os.Stderr, "%s: Skipping non-regular file\n", path)
		}
	}
}

// RecursiveSearch performs a recursive search on any directories in the
// input set.
//
// All actual files in the input set and ones found via the recursive search
// are sent to the output channel.
func RecursiveSearch(inputs []string) chan PathItem {
	out := make(chan PathItem, cli.ChannelBufferSize)
	go func() {
		defer close(out)
		recursiveSearch(inputs, out)
	}()
	return out
}
