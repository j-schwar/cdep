package cli

// InputFiles stores the list of input files to search.
var InputFiles []string

// Outfile is the file to dump output to.
var Outfile string

// UseFileRecursion tells the search system to recursively search directories or no.
var UseFileRecursion bool

// Verbose enables extra output to stderr.
var Verbose bool

// ChannelBufferSize defines the size of pipeline channel buffers.
var ChannelBufferSize int

// DisplayAll tells the system whether it should filter out some dependencies or not.
var DisplayAll bool

// UseMerge tells the system to merge all dependencies into a single list.
var UseMerge bool

// UseCount tells the system to print out the number of references to each dependency.
var UseCount bool

// ShowVersion tells the application to print out the version number.
var ShowVersion bool
