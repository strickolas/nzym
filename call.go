package nzym

import (
	"io/ioutil"
	"log"
	"strings"
)

// Call uses the longest sequence of tokens
// starting from the left to call the appropriate
// nzym alias.
func Call(args []string) {
	tokens := strings.Split(Clean(strings.Join(args, " ")), " ")
	nza, index, alias, command := GetConfig(), -1, "", ""
	for i := 1; i <= len(tokens); i++ {
		alias = strings.Join(args[0:i], " ")
		if val, ok := nza[alias]; ok {
			command, index = val, i
		}
	}
	if index < 0 {
		Fatal("nzym: no alias found")
	}

	tokens = append(strings.Split(command, " "), tokens[index:len(tokens)]...)
	command = "#!/bin/bash\n" + strings.Join(tokens, " ") + "\n"
	err := ioutil.WriteFile("/tmp/nzym_command.sh", []byte(command), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
