package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Set up testing here, if needed
	exitVal := m.Run()
	// Clean up or do additional assertions here, if needed
	os.Exit(exitVal)
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()

	w.Close()
	os.Stdout = old
	return <-outC
}

func TestHelloWorld(t *testing.T) {
	expectedOutput := "Hello, World!\n"

	actualOutput := captureOutput(main)

	if actualOutput != expectedOutput {
		t.Errorf("Expected: %s, but got: %s", expectedOutput, actualOutput)
	}
}
