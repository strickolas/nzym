package nzym

import (
	"os"
	"os/exec"
	"strings"
)

// Call uses the longest sequence of tokens
// starting from the left to call the appropriate
// nzym alias.
func Call(args []string, nzymrcPath string) {
	tokens := strings.Split(Clean(strings.Join(args, " ")), " ")
	nza, index, alias, command := GetConfig(nzymrcPath), -1, "", ""
	for i := 1; i <= len(tokens); i++ {
		alias = strings.Join(args[0:i], " ")
		if val, ok := nza[alias]; ok {
			command, index = val, i
		}
	}
	if index < 0 {
		Fatal("nzym: Alias not found.")
	}

	tokens = append(strings.Split(command, " "), tokens[index:len(tokens)]...)
	cmd := exec.Command(tokens[0], tokens[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()
}
