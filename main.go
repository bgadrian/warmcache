package main

import (
	"github.com/bgadrian/warmcache/scanner"
	"github.com/mkideal/cli"
)

func main() {
	cli.Run(new(scanner.CLIArguments), scanner.Scan)
}
