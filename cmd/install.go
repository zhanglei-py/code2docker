package cmd

import (
	"github.com/spf13/cobra"
	"nset-cli/imp"
	"strings"
)

func init() {
	InstallCmd.Flags().StringVarP(&moduleName, "modules", "m", "", "modules of nset")
	InstallCmd.Flags().StringVarP(&moduleDir, "workdir", "w", "", "workdir of nset")
	RootCmd.AddCommand(InstallCmd)
}

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "-m|--modules <module1>,<module2>",
	Long:  `nset-launch install <module>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(moduleName) == 0 {
			return
		}
		s := strings.Split(moduleName, ",")
		if len(s) == 0 {
			imp.RoleInstall(moduleName, moduleDir)
		} else {
			for _, v := range s {
				imp.RoleInstall(v, moduleDir)
			}
		}
	},
}
