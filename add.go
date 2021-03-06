package nzym

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

// Add adds a new nzym alias.
func Add(args []string, nzymrcPath string) {
	usage := `
Usage:	nzym add nz a DOES $(which nzym) add
        nzym add $(which nzym) add AS nz a
		
Adds an Nzym Alias. Use single quotes to pass verbatim arguments to Nzym.
`
	if len(args) == 0 {
		Fatal(usage)
	}
	line := strings.Join(args, " ")
	var alias, command string
	nza := GetConfig(nzymrcPath)

	containsDOES := strings.Contains(line, " DOES ")
	containsAS := strings.Contains(line, " AS ")

	if containsDOES && containsAS {
		// ambiguous statement
		Fatal(usage)
	} else if containsDOES {
		// <alias> DOES <command>
		args = strings.Split(line, "DOES")
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
	err = WriteStringToFile(nzymrcPath, string(d))
	if err != nil {
		Fatal("nzym: " + err.Error())
	}
	fmt.Println("nzym: Alias written:\x1b[33m", alias, "->", command, "\x1b[m")
}
