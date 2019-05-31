# poc-go-subcommands

```bash
$ poc-go-subcommands
Usage: poc-go-subcommands <flags> <subcommand> <subcommand args>

Subcommands:
	commands         list all command names
	flags            describe all known top-level flags
	help             describe subcommands and their syntax

Subcommands for subcommands:
	subcommands      nested subcommands


Top-level flags (use "poc-go-subcommands flags" for a full list):
  -root=default value: root flag

$ poc-go-subcommands subcommands
Usage: subcommands <flags> <subcommand> <subcommand args>

Subcommands:
	commands         list all command names
	describe         print flags
	flags            describe all known top-level flags
	help             describe subcommands and their syntax


Top-level flags (use "subcommands flags" for a full list):
  -nested=default value: nested flag
```

```
$ poc-go-subcommands subcommands describe foo
root=default value
nested=default value
text=foo

$ poc-go-subcommands -root x subcommands -nested y describe bar
root=x
nested=y
text=bar
```
