package imp

import (
	"os"
	"path/filepath"
	"gopkg.in/yaml.v2"
)
type RoleName struct {
	Name string `yaml:"name"`
}
type modulesRole struct {
	LocalRoles []RoleName `yaml:"roles,flow"`
}

func RoleInstall (moduleName string) bool {
	var localVars modulesRole
	dir, _ := os.Executable()
	ExecPath := filepath.Dir(dir)
	moduleVarsFile := ExecPath + "/modules/" + moduleName + "/vars.yaml"
	moduleVarsFileContent, err := ReadVarsFile(moduleVarsFile, moduleName)
	if err != nil {
		return false
	}
	yaml.Unmarshal([]byte(moduleVarsFileContent), &localVars)
	if len(localVars.LocalRoles) == 0 {
		commmandString, _ := AnsbileCmd(moduleName, "_NOROLE_")
		ExecCommand(commmandString)
	} else {
		for _, v := range localVars.LocalRoles {
			commmandString, _ := AnsbileCmd(moduleName, v.Name)
			ExecCommand(commmandString)
		}
	}
	return true
}