package cli

import (
	"bufio"
	"os"
	"strings"
)

var exitSig bool

func CliListen() {
	for !exitSig {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		args := strings.Split(scanner.Text(), " ")
		Execute(args)
	}
}

func Exit() {
	exitSig = true
}
