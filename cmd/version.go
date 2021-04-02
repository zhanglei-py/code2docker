package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(VersionCmd)
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `This is Nset CLI version.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("v0.1 -- Bate")
	},
}
