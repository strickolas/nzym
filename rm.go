package nzym

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

// Rm removes an alias from nzymyml.
func Rm(args []string, nzymrcPath string) {
	usage := `
Usage: nzym rm nz
	
Removes an Nzym Alias. Use single quotes to pass verbatim arguments to Nzym.
`
	if len(Clean(strings.Join(args, " "))) == 0 {
		Fatal(usage)
	}
	nza := GetConfig(nzymrcPath)
	alias := strings.Join(args, " ")
	if _, ok := nza[alias]; ok {
		delete(nza, alias)
	} else {
		fmt.Println("nzym: rm: Alias \x1b[33m" + alias + "\x1b[m does not exist.")
	}

	d, err := yaml.Marshal(&nza)
	if err != nil {
		Fatal("nzym: " + err.Error())
	}
	err = WriteStringToFile(nzymrcPath, string(d))
	if err != nil {
		Fatal("nzym: " + err.Error())
	}
}
