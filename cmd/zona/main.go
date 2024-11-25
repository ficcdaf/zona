package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ficcdaf/zona/internal/build"
)

// // validateFile checks whether a given path
// // is a valid file && matches an expected extension
// func validateFile(path, ext string) bool {
// 	return (util.CheckExtension(path, ext) == nil) && (util.PathIsValid(path, true))
// }

func main() {
	rootPath := flag.String("file", "", "Path to the markdown file.")
	flag.Parse()
	if *rootPath == "" {
		// no flag provided, check for positional argument instead
		n := flag.NArg()
		var e error
		switch n {
		case 1:
			// we read the positional arg
			arg := flag.Arg(0)
			// mdPath wants a pointer so we get arg's address
			rootPath = &arg
		case 0:
			// in case of no flag and no arg, we fail
			e = errors.New("Required argument missing!")
		default:
			// more args than expected is also fail
			e = errors.New("Unexpected arguments!")
		}
		if e != nil {
			fmt.Printf("Error: %s\n", e.Error())
			os.Exit(1)
		}

	}
	err := build.Traverse(*rootPath, "foobar")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
