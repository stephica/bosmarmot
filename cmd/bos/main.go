package main

import (
	"fmt"
	"os"

	"code.monax.io/platform/bosmarmot/cmd"
	"github.com/jawher/mow.cli"
)

func main() {
	bos := cli.App("bos",
		"Deep in the Burrow")
	bos.Action = func() {
		fmt.Println("BOSMARMOT")
	}
	cmd.AddVersionCommand(bos)
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
