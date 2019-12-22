package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		help()
	} else if args[0] == "add" {
		add(args[1:])
	} else if args[0] == "rm" {
		rm(args[1:])
	} else if args[0] == "mv" {
		mv(args[1:])
	} else if args[0] == "ls" {
		ls(args[1:])
	} else {
		help()
	}
}

func getConfig() map[string]string {
	// Read contents of config file.
	config, err := ioutil.ReadFile("../../nzym.yml")
	if err != nil {
		log.Fatalf("Unable to locate \"nzym.yml\".")
	}

	// Unmarshal the config file.
	var nzymAlises map[string]string
	err = yaml.Unmarshal([]byte(config), &nzymAlises)
	if err != nil {
		log.Fatalf("Unable to read \"nzym.yml\".")
	}
	return nzymAlises
}

func add(args []string) {
	usage := `
Usage:	nzym add nz a TO $(which nzym) add
        nzym add $(which nzym) AS nz a
		
		
Adds an Nzym Alias. Use single quotes to pass verbatim arguments to Nzym.`
	if len(args) == 0 {
		log.Fatalf(usage)
	}

	line := strings.Join(args, " ")
	var alias, command string
	if strings.Contains(line, "TO") {
		args = strings.Split(line, "TO")
		alias, command = args[0], args[1]
		if len(args) > 2 {
			log.Fatalf(usage)
		}
	} else if strings.Contains(line, "AS") {
		args = strings.Split(line, "AS")
		if len(args) > 2 {
			log.Fatalf(usage)
		}
		alias, command = args[1], args[0]
	} else {
		log.Fatalf(usage)
	}

	fmt.Println(alias, command)
}

func rm(args []string) {

}

func mv(args []string) {

}

func ls(args []string) {

}

func help() {
	fmt.Println(
		`
Usage:	nzym COMMAND

A command line tool that makes it easy to manage scripts!

Commands:
  add    Adds a new alias.
  rm     Removes an existing alias.
  mv     Renames an existing alias.
  ls     Lists all aliases stored in nzym.`,
	)
}
