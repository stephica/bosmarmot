package cmd

import (
	"fmt"
	"os"

	"github.com/monax/bosmarmot/compilers/perform"

	"github.com/monax/bosmarmot/monax/log"
	"github.com/spf13/cobra"
)

func BuildBinaryCommand() {
	CompilersCmd.AddCommand(binaryCmd)
	addBinaryFlags()
}

var (
	binaryPort   string
	binaryUrl    string
	binaryDir    string
	binlibraries string
	binarySSL    bool
	binaryLocal  bool
)

var binaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "link a binary to an address",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Errorf("Specify a contract to compile \n\n")
			CompilersCmd.Help()
			os.Exit(0)
		}
		output, err := perform.RequestBinaryLinkage(args[0], libraries)
		if err != nil {
			log.Error(err)
		}
		log.WithFields(log.Fields{
			"binary": output.Binary,
			"error":  output.Error,
		}).Warn("Output")
	},
}

func addBinaryFlags() {
	binaryCmd.Flags().StringVarP(&binaryDir, "dir", "D", setDefaultDirectoryRoute(true), "directory location to search for on the remote server")
	binaryCmd.Flags().StringVarP(&binlibraries, "libs", "L", "", "libraries string (libName:Address[, or whitespace]...)")
	binaryCmd.Flags().BoolVarP(&binaryLocal, "local", "l", setCompilerLocal(), "use local compilers to compile message (good for debugging or if server goes down)")
}
