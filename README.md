# nzym
A command line tool that makes it easy to digest 
scripts for Unix based systems!

# Installation
1. Download zip file from Git.
2. Unzip.
3. Navigate to /nzym/cmd/nzym
4. In your terminal, type `go build nzym.go`
5. Next, type `mv ./nzym /usr/bin/`
6. Start using Nzym.

```bash
cd Downloads
unzip nzym-master.zip
cd nzym-master/cmd/nzym
go build nzym.go
sudo mv ./nzym /usr/bin/
```


# Examples

```bash
# nzym add <alias> DOES <command>.
nzym add dc DOES docker compose

# nzym add <command> AS <alias>.
nzym add $(which docker) volume AS dc 

# nzym can use several tokens as an alias.
nzym add my py script DOES '$(which python)' $PWD/main.py

# list aliases in a columnar format>.
nzym ls

# rename aliases using the TO direction (no support for DOES keyword).
nzym mv my py script TO my py

# call an alias, adding whatever CLI args necessary.
nzym my py file1.txt file2.txt --blue

# remove an alias.
nzym rm my py
```
