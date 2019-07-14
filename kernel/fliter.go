package kernel

import "cdep/cli"

var cppStdHeaders = []string{
	"cstdlib",
	"csignal",
	"csetjmp",
	"typeinfo",
	"typeindex",
	"type_traits",
	"bitset",
	"functional",
	"utility",
	"ctime",
	"chrono",
	"cstddef",
	"initializer_list",
	"tuple",
	"any",
	"optional",
	"variant",
	"new",
	"memory",
	"scoped_allocation",
	"memory_resource",
	"climits",
	"cfloat",
	"cstdint",
	"cinttypes",
	"limits",
	"exception",
	"stdexcept",
	"cassert",
	"system_error",
	"cerrno",
	"cctype",
	"cwctype",
	"cstring",
	"cuchar",
	"string",
	"string_view",
	"charconv",
	"array",
	"vector",
	"deque",
	"list",
	"forward_list",
	"set",
	"map",
	"unordered_set",
	"unordered_map",
	"stack",
	"queue",
	"iterator",
	"algorithm",
	"cmath",
	"complex",
	"random",
	"numeric",
	"ratio",
	"cfenv",
	"iosfwd",
	"ios",
	"istream",
	"ostream",
	"sstream",
	"fstream",
	"iomanip",
	"iostream",
	"streambuf",
	"cstdio",
	"locale",
	"clocale",
	"regex",
	"atomic",
	"thread",
	"mutex",
	"shared_mutex",
	"future",
	"conditional_variable",
	"filesystem",
	"codecvt",
}

func isStdDep(dep string) bool {
	for _, stdHeader := range cppStdHeaders {
		if stdHeader == dep {
			return true
		}
	}
	return false
}

// FilterSourceDependenciesRemoveStdDeps removes any standard library input paths
// from a SourceDependenciesItem.
func FilterSourceDependenciesRemoveStdDeps(ch chan SourceDependenciesItem) chan SourceDependenciesItem {
	out := make(chan SourceDependenciesItem, cli.ChannelBufferSize)
	go func() {
		defer close(out)

		for item := range ch {
			var filteredDeps []string
			for _, dep := range item.Dependencies {
				if !isStdDep(dep) {
					filteredDeps = append(filteredDeps, dep)
				}
			}
			item.Dependencies = filteredDeps
			out <- item
		}
	}()
	return out
}

func RemoveDuplicateDependencies(ch chan DependencyItem) chan DependencyItem {
	out := make(chan DependencyItem, cli.ChannelBufferSize)
	go func() {
		defer close(out)

		memoizer := map[DependencyItem]bool{}
		for item := range ch {
			if !memoizer[item] {
				out <- item
				memoizer[item] = true
			}
		}
	}()
	return out
}
