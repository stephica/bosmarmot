package main

import (
	"fmt"
	"os"

	"code.monax.io/platform/bosmarmot/consensus/tendermint"
	"code.monax.io/platform/bosmarmot/logging/lifecycle"
	"github.com/jawher/mow.cli"
)

func main() {
	bos := cli.App("bos",
		"Deep in the Burrow")
	bos.Action = func() {
		//logger, _ := lifecycle.NewStdErrLogger()
		//tendermint.LaunchGenesisValidator(logger)
	}
	bos.Run(os.Args)
}

// Print informational output to Stderr
func printf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
