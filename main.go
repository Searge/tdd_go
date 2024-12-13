package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Searge/tdd_go/di"
	"github.com/Searge/tdd_go/helloworld"
	"github.com/Searge/tdd_go/mocking"
)

// Repo: https://github.com/Searge/tdd_go
// License: https://github.com/Searge/tdd_go/blob/main/LICENSE

func run(language string, help bool, out io.Writer) {
	if help {
		fmt.Fprintln(out, "Usage: go run main.go [options] [command]")
		fmt.Fprintln(out, "Options:")
		fmt.Fprintln(out, "  -help        Show help")
		fmt.Fprintln(out, "  -lang string  Language for the greeting (default \"English\")")
		fmt.Fprintln(out, "Commands:")
		fmt.Fprintln(out, "  hello: Say a greeting in the specified language")
		fmt.Fprintln(out, "  help:  Show this help message")
		return
	}

	fmt.Fprintln(out, helloworld.Hello("World", language))
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "Searge")
}
func main() {
	language := flag.String("lang", "English", "Language for the greeting")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	run(*language, *help, os.Stdout)

	// Dependency injection code
	di.Greet(os.Stdout, "Elodie")
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))

	fmt.Printf("\n")

	// Mocking code
	mocking.Countdown(os.Stdout)
}
