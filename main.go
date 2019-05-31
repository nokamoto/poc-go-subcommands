package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"os"
)

var root = flag.String("root", "default value", "root flag")

func main() {
	flag.Parse()

	subcommands.ImportantFlag("root")

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	subcommands.Register(&nestedSubcommands{}, "subcommands")

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
