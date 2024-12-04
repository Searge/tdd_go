package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	// Create a new bytes buffer to capture the output
	buf := &bytes.Buffer{}

	// Test the help command
	run("", true, buf)
	if buf.String() != "Usage: go run main.go [options] [command]\nOptions:\n  -help        Show help\n  -lang string  Language for the greeting (default \"English\")\nCommands:\n  hello: Say a greeting in the specified language\n  help:  Show this help message\n" {
		t.Errorf("expected help message to be printed, but got %q", buf.String())
	}

	// Reset the buffer
	buf.Reset()

	// Test the hello command
	run("Spanish", false, buf)
	if buf.String() != "Hola, World\n" {
		t.Errorf("expected hello message to be printed, but got %q", buf.String())
	}
}
