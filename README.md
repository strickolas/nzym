# nzym
A command line tool that makes it easy to digest 
scripts!

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
