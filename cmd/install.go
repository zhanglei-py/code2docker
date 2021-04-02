package cmd

import (
	"strings"
	"launch/imp"
	"github.com/spf13/cobra"
)

func init() {
	InstallCmd.Flags().StringVarP(&moduleName,"modules","m","","modules of nset")
	RootCmd.AddCommand(InstallCmd)
}

var InstallCmd = &cobra.Command {
	Use:   "install",
	Short: "-m|--modules <module1>,<module2>",
	Long:  `nset-launch install <module>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(moduleName) == 0 {
			return
		}
		s := strings.Split(moduleName, ",")
		if len(s) == 0 {
			imp.RoleInstall(moduleName)
		} else {
			for _, v := range s {
				imp.RoleInstall(v)
			}
		}
	},
}