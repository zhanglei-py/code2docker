package cmd

import (
	"github.com/spf13/cobra"
	"nset-cli/imp"
	"strings"
)

func init() {
	RunCmd.Flags().StringVarP(&commandName, "command", "c", "", "commands of ssh")
	RunCmd.Flags().StringVarP(&moduleName, "module", "m", "", "module name")
	RunCmd.Flags().StringVarP(&hostList, "hosts", "H", "", "host list")
	RootCmd.AddCommand(RunCmd)
}

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "-c|--command <command1>;<command2> -m|--module <module> [-H|--hosts <host1>,<addr1>]",
	Long:  `run remote command`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(commandName) == 0 {
			return
		}
		sshHosts, _ := imp.SshHosts(moduleName, hostList)
		s := strings.Split(commandName, ";")
		if len(s) == 0 {
			imp.PsshRun(commandName, sshHosts)
		} else {
			for _, v := range s {
				imp.PsshRun(v, sshHosts)
			}
		}
	},
}
