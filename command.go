package main

import (
	"flag"
	"fmt"
	"strings"
)

var flagProcfile string

type Command struct {
	// args does not include the command name
	Run  func(cmd *Command, args []string)
	Flag flag.FlagSet

	Disabled bool
	Usage    string // first word is the command name
	Short    string // `forego help` output
	Long     string // `forego help cmd` output
}

func (c *Command) printUsage() {
	if c.Runnable() {
		fmt.Printf("Usage: forego %s\n\n", c.Usage)
	}
	fmt.Println(strings.Trim(c.Long, "\n"))
}

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Runnable() bool {
	return c.Run != nil && !c.Disabled
}

func (c *Command) List() bool {
	return c.Short != ""
}
