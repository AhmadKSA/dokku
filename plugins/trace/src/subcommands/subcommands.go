package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dokku/dokku/plugins/common"
	"github.com/dokku/dokku/plugins/trace"

	flag "github.com/spf13/pflag"
)

// main entrypoint to all subcommands
func main() {
	parts := strings.Split(os.Args[0], "/")
	subcommand := parts[len(parts)-1]

	var err error
	switch subcommand {
	case "off":
		args := flag.NewFlagSet("trace:off", flag.ExitOnError)
		args.Parse(os.Args[2:])
		err = trace.CommandOff()
	case "on":
		args := flag.NewFlagSet("trace:on", flag.ExitOnError)
		args.Parse(os.Args[2:])
		err = trace.CommandOn()
	default:
		common.LogFail(fmt.Sprintf("Invalid plugin subcommand call: %s", subcommand))
	}

	if err != nil {
		common.LogFail(err.Error())
	}
}
