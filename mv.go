package nzym

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

// Mv can be used to give an alias a different
// name. It can also be used to swap two existing
// aliases if destination alias already exists.
func Mv(args []string, nzymrcPath string) {
	usage := `
Usage: nzym mv nz mv AS nz move
		
Moves an Nzym Alias. Use single quotes to pass verbatim arguments to Nzym.
`
	if len(args) == 0 {
		Fatal(usage)
	}
	line := strings.Join(args, " ")
	var alias1, alias2 string
	nza := GetConfig(nzymrcPath)

	containsTO := strings.Contains(line, " TO ")

	if containsTO {
		// <alias> TO <command>
		args = strings.Split(line, "TO")
		alias1, alias2 = Clean(args[0]), Clean(args[1])
		if len(args) != 2 {
			Fatal(usage)
		}
	} else {
		// malformed command
		Fatal(usage)
	}

	// Check if the alias is already defined.
	if _, ok := nza[alias1]; !ok {
		Fatal("nzym: Unable to move \x1b[33m" + alias1 + " -> " + alias2 + "\x1b[m. Alias \x1b[33m" + alias1 + "\x1b[m does not exist.")
	}
	if _, ok := nza[alias2]; ok {
		Fatal("nzym: Unable to move \x1b[33m" + alias1 + " -> " + alias2 + "\x1b[m. Alias already exists.")
	}

	// Marshal the map and write it out!
	nza[alias2] = nza[alias1]
	delete(nza, alias1)
	d, err := yaml.Marshal(&nza)
	if err != nil {
		Fatal("nzym: " + err.Error())
	}
	err = WriteStringToFile(nzymrcPath, string(d))
	if err != nil {
		Fatal("nzym: " + err.Error())
	}
	fmt.Println("nzym: Alias moved:\x1b[33m", alias1, "->", alias2, "\x1b[m")
}
