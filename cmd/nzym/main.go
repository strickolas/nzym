package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/strickolas/nzym"
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
	nza := nzym.GetConfig()

	if strings.Contains(line, "TO") {
		args = strings.Split(line, "TO")
		alias, command = clean(args[0]), clean(args[1])
		if len(args) > 2 {
			log.Fatalf(usage)
		}
	} else if strings.Contains(line, "AS") {
		args = strings.Split(line, "AS")
		if len(args) > 2 {
			log.Fatalf(usage)
		}
		alias, command = clean(args[1]), clean(args[0])
	} else {
		log.Fatalf(usage)
	}

	if val, ok := nza[alias]; ok {
		a := "nzym: alias already exists as \x1b[33m" + alias + " -> " + val + " \x1b[mOverwrite? (yes/no): "
		if getConfirmation(a, a) {
			nza[alias] = command

			d, err := yaml.Marshal(&nza)
			if err != nil {
				log.Fatalf("nzym: %v", err)
			}

			err = writeStringToFile("../../nzym.yml", string(d))
			if err != nil {
				log.Fatalf("nzym: %v", err)
			}

			fmt.Println("nzym: alias written:", alias, "->", command)
		}
	} else {
		fmt.Println(alias, command)
	}
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
