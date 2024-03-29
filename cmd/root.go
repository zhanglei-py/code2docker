package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	moduleName  string
	moduleDir   string
	commandName string
	localFiles  string
	remoteDir   string
	hostList    string
	iniFile     string
	sectionName string
)

var RootCmd = &cobra.Command{
	Use:   "launch",
	Short: "launch help",
	Long:  `nset-launch is a bootstrap program for modules (NSET)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
