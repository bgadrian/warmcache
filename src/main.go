package main

import (
	"github.com/bgadrian/hot-cache-crawler/src/scanner"
	"github.com/mkideal/cli"
)

func main() {
	cli.Run(new(scanner.CLIArguments), scanner.Scan)
}
