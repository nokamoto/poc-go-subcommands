package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
)

type nestedSubcommands struct {
	commander *subcommands.Commander
}

var nested *string

func (*nestedSubcommands) Name() string {
	return "subcommands"
}

func (*nestedSubcommands) Synopsis() string {
	return "nested subcommands"
}

func (c *nestedSubcommands) Usage() string {
	return c.commander.HelpCommand().Usage()
}

func (c *nestedSubcommands) SetFlags(f *flag.FlagSet) {
	c.commander = subcommands.NewCommander(f, "subcommands")
	c.commander.Register(c.commander.HelpCommand(), "")
	c.commander.Register(c.commander.FlagsCommand(), "")
	c.commander.Register(c.commander.CommandsCommand(), "")
	c.commander.Register(&nestedDescribeSubcommands{}, "")
	nested = f.String("nested", "default value", "nested flag")
	c.commander.ImportantFlag("nested")
}

func (c *nestedSubcommands) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	return c.commander.Execute(ctx, args...)
}
