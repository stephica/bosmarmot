package cmd

import (
	"fmt"
	"os"

	"github.com/monax/bosmarmot/compilers/perform"
	"github.com/monax/bosmarmot/monax/log"
	"github.com/spf13/cobra"
)

func BuildCompileCommand() {
	CompilersCmd.AddCommand(compileCmd)
	addCompileFlags()
}

var (
	compilerDir   string
	libraries     string
	compilerLocal bool
	optimizeSolc  bool
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "compile your contracts either remotely or locally",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Errorf("Specify a contract to compile \n\n")
			CompilersCmd.Help()
			os.Exit(0)
		}

		output, err := perform.RequestCompile(args[0], optimizeSolc, libraries)
		if err != nil {
			log.Error(err)
		}
		perform.PrintResponse(*output, true)
	},
}

func addCompileFlags() {
	compileCmd.Flags().StringVarP(&compilerDir, "dir", "D", setDefaultDirectoryRoute(false), "directory location to search for on the remote server")
	compileCmd.Flags().StringVarP(&libraries, "libs", "L", "", "libraries string (libName:Address[, or whitespace]...)")
	compileCmd.Flags().BoolVarP(&compilerLocal, "local", "l", setCompilerLocal(), "use local compilers to compile message (good for debugging or if server goes down)")
	compileCmd.Flags().BoolVarP(&optimizeSolc, "optimize", "o", setOptimizeSolc(), "optimize code (solidity only)")
}

func setOptimizeSolc() bool {
	return false
}

func setCompilerLocal() bool {
	return false
}

func setDefaultDirectoryRoute(binaries bool) string {
	if binaries {
		return "/binaries"
	}
	return "/"
}
