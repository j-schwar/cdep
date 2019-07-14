package kernel

import (
	"cdep/cli"
	"io/ioutil"
)

// This file defines kernels which open source files and retrieve their contents.

// SourceItem is a pipeline item which stores a source file's path and contents.
type SourceItem struct {
	FilePath PathItem
	Contents string
}

// Open is a kernel which retrieves the contents of files.
//
// If unable to open a file, a fatal error is logged.
func Open(ch chan PathItem) chan SourceItem {
	out := make(chan SourceItem, cli.ChannelBufferSize)
	go func() {
		defer close(out)

		for item := range ch {
			contents, err := ioutil.ReadFile(string(item))
			if err != nil {
				fatalError(err)
			}
			out <- SourceItem{item, string(contents)}
		}
	}()
	return out
}
