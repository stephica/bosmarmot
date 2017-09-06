package cmd

import (
	"fmt"

	"github.com/jawher/mow.cli"
	"code.monax.io/platform/bosmarmot/release"
)

func AddVersionCommand(cmd *cli.Cli) {
	cmd.Command("version", "Get version number",
		func(versionCmd *cli.Cmd) {
			versionCmd.Action = func() {
				fmt.Println(release.Version())
			}

			versionCmd.Command("notes", "Get release notes for this version",
				func(changesCmd *cli.Cmd) {
					changesCmd.Action = func() {
						fmt.Println(release.Notes())
					}
				})

			versionCmd.Command("changelog", "Get changelog",
				func(changesCmd *cli.Cmd) {
					changesCmd.Action = func() {
						fmt.Println(release.Changelog())
					}
				})

		})
}
