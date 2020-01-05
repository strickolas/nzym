package main

import (
	"fmt"
	"os"

	"github.com/strickolas/nzym"
)

func main() {
	args := os.Args[1:]
	path := os.ExpandEnv("/home/$USER/.nzymrc")
	if !nzym.FileExists(path) {
		_, err := os.Create(path)
		if err != nil {
			nzym.Fatal("nzym: " + err.Error())
		}
	}
	if len(args) == 0 {
		help()
	} else if args[0] == "add" {
		nzym.Add(args[1:], path)
	} else if args[0] == "rm" {
		nzym.Rm(args[1:], path)
	} else if args[0] == "mv" {
		nzym.Mv(args[1:], path)
	} else if args[0] == "ls" {
		nzym.Ls(args[1:], path)
	} else {
		nzym.Call(args, path)
	}
}

func help() {
	fmt.Println(
		`
Usage:	nzym COMMAND

A command line tool that makes it easy to digest scripts!

Commands:
  add    Adds a new alias.
  rm     Removes an existing alias.
  mv     Renames an existing alias.
  ls     Lists all aliases stored in nzym.`,
	)
}
