package nzym

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

// Rm removes an alias from nzym.yml.
func Rm(args []string) {
	nza := GetConfig()
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
	err = WriteStringToFile("../../nzym.yml", string(d))
	if err != nil {
		Fatal("nzym: " + err.Error())
	}
}
