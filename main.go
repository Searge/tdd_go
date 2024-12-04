package main

import (
	"flag"
	"fmt"

	"github.com/Searge/tdd_go/helloworld"
)

// main is the entry point for the program which parses command-line flags
// for language and help options. It displays usage information if the help
// flag is set. Otherwise, it calls the Hello function from the helloworld
// package to print a greeting in the specified language.
func main() {
	language := flag.String("lang", "English", "Language for the greeting")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		fmt.Println("Usage: go run main.go [options] [command]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("Commands:")
		fmt.Println("  hello: Say a greeting in the specified language")
		fmt.Println("  help:  Show this help message")
		return
	}

	fmt.Println(helloworld.Hello("World", *language))
}
