package nzym

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

// Add adds a new nzym alias.
func Add(args []string) {
	usage := `
Usage:	nzym add nz a TO $(which nzym) add
        nzym add $(which nzym) AS nz a
		
Adds an Nzym Alias. Use single quotes to pass verbatim arguments to Nzym.
`
	if len(args) == 0 {
		Fatal(usage)
	}
	line := strings.Join(args, " ")
	var alias, command string
	nza := GetConfig()

	containsTO := strings.Contains(line, " TO ")
	containsAS := strings.Contains(line, " AS ")

	if containsTO && containsAS {
		// ambiguous statement
		Fatal(usage)
	} else if containsTO {
		// <alias> TO <command>
		args = strings.Split(line, "TO")
		alias, command = Clean(args[0]), Clean(args[1])
		if len(args) != 2 {
			Fatal(usage)
		}
	} else if containsAS {
		// <command> AS <alias>
		args = strings.Split(line, "AS")
		if len(args) != 2 {
			Fatal(usage)
		}
		alias, command = Clean(args[1]), Clean(args[0])
	} else {
		// no direction given
		Fatal(usage)
	}

	// Check if the alias is already defined.
	if val, ok := nza[alias]; ok {
		if nza[alias] == command {
			Fatal("nzym: Alias \x1b[33m" + alias + " -> " + val + "\x1b[m already exists.")
		}
		a := "nzym: Overwrite \x1b[33m" + alias + " -> " + val +
			" \x1b[m with \x1b[33m" + alias + " -> " + command + "\x1b[m ? (yes/no): "
		if !GetConfirmation(a, a) {
			Fatal("nzym: Did not write alias.")
		}
	}
	// Marshal the map and write it out!
	nza[alias] = command
	d, err := yaml.Marshal(&nza)
	if err != nil {
		Fatal("nzym: " + err.Error())
	}
	err = WriteStringToFile("../../nzym.yml", string(d))
	if err != nil {
		Fatal("nzym: " + err.Error())
	}
	fmt.Println("nzym: alias written:", alias, "->", command)
}
