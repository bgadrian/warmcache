package main

import (
	"os"

	"github.com/bgadrian/hot-cache-crawler/src/scanner"
	"github.com/mkideal/cli"
)

func main() {
	os.Exit(cli.Run(new(scanner.CLIArguments), scanner.Scan))
}
