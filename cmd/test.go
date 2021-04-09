package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nset-cli/imp"
)

func init() {
	TestCmd.Flags().StringVarP(&iniFile, "ini", "i", "", "ini filename")
	TestCmd.Flags().StringVarP(&sectionName, "section", "s", "", "section of ini")
	RootCmd.AddCommand(TestCmd)
}

var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "-i|--ini <file.ini>",
	Long:  `nset-launch parse ini file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(iniFile) == 0 {
			return
		} else {
			/*
				sectionTag := iniFile[0:1] + iniFile[len(iniFile)-1:]
				fmt.Println(sectionTag)
			*/
			section, _ := imp.ReadInventoryIni(iniFile, sectionName)
			for _, v := range section {
				fmt.Println(v)
			}
		}
	},
}
