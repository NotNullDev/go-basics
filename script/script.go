package main

import (
	"os"
	"os/exec"
	"strings"
)

// go run script.go
func main() {
	cmd := exec.Command("bash", "./test.sh")

	input := `jacek
23
`

	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = strings.NewReader(input)

	out, err := cmd.Output()

	if err != nil {
		print("err: ", string(err.Error()))
	} else {
		print("out: ", string(out))
	}

}
