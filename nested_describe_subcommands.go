package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
)

type nestedDescribeSubcommands struct{}

func (*nestedDescribeSubcommands) Name() string {
	return "describe"
}

func (*nestedDescribeSubcommands) Synopsis() string {
	return "print flags"
}

func (*nestedDescribeSubcommands) Usage() string {
	return "describe\n\tPrint flags.\n"
}

func (*nestedDescribeSubcommands) SetFlags(_ *flag.FlagSet) {}

func (c *nestedDescribeSubcommands) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if len(f.Args()) != 0 {
		f.Usage()
		return subcommands.ExitUsageError
	}

	fmt.Printf("root=%v\n", *root)
	fmt.Printf("nested=%v\n", *nested)

	return subcommands.ExitSuccess
}
