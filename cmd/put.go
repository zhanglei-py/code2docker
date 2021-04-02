package cmd

import (
	"github.com/spf13/cobra"
	"launch/imp"
	"strings"
)

func init() {
	PutCmd.Flags().StringVarP(&localFiles,"local","l","","local file path")
	PutCmd.Flags().StringVarP(&remoteDir,"remote","r","","remote dir")
	PutCmd.Flags().StringVarP(&moduleName,"module","m","","module name")
	PutCmd.Flags().StringVarP(&hostList,"hosts","H","","host list")
	RootCmd.AddCommand(PutCmd)
}

var PutCmd = &cobra.Command {
	Use:   "put",
	Short: "-l|--local <file1>,<file2> -r|--remote <dir path> -m|--module <module name> [-H|--hosts <host1>,<addr1>]",
	Long:  `nset-launch put for pscp`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(localFiles) == 0 {
			return
		}
		sshHosts, _ := imp.SshHosts(moduleName, hostList)
		s := strings.Split(localFiles, ",")
		if len(s) == 0 {
			imp.PscpCopy(localFiles, remoteDir, sshHosts)
		} else {
			for _, v := range s {
				imp.PscpCopy(v, remoteDir, sshHosts)
			}
		}
	},
}