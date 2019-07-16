package cli

import "flag"

// ParseCommandLineArguments sets up and parses command line flags.
func ParseCommandLineArguments() {
	flag.StringVar(&Outfile, "o", "", "Output file, none for stdout")
	flag.BoolVar(&DisplayAll, "a", false, "Display all dependencies including std headers")
	flag.BoolVar(&UseFileRecursion, "r", false, "Recursively search directories for source files")
	flag.BoolVar(&UseMerge, "m", false, "Merge all dependencies into a single list")
	flag.BoolVar(&UseCount, "c", false, "Print out the number of references for each dependency")
	flag.BoolVar(&ShowDirectoriesOnly, "d", false, "Display directories of dependencies and not files")
	flag.BoolVar(&ListSourcesOnly, "l", false, "List only source files and no dependencies.")
	flag.BoolVar(&Verbose, "v", false, "Enable verbose logging to stderr")
	flag.IntVar(&ChannelBufferSize, "buffer-size", 16, "Size of pipeline channel buffers")
	flag.BoolVar(&ShowVersion, "version", false, "Display version information")
	flag.StringVar(&FindExp, "find", "", "A regular expression to match dependencies")
	flag.Parse()
	InputFiles = flag.Args()
}
