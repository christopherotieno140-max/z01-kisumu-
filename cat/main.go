package main

import (
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	hasError := false
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			hasError = true
			break
		}
		_, err = io.Copy(os.Stdout, file)
		file.Close()
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			hasError = true
			break
		}
	}

	if hasError {
		os.Exit(1)
	}
}
