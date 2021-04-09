package imp

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type RoleName struct {
	Name string `yaml:"name"`
}
type modulesRole struct {
	LocalRoles []RoleName `yaml:"roles,flow"`
}

func RoleInstall(moduleName, modulesDir string) bool {
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
		commmandString, _ := AnsibleCmd(moduleName, "_NO_ROLE_", modulesDir)
		ExecCommand(commmandString)
	} else {
		for _, v := range localVars.LocalRoles {
			commmandString, _ := AnsibleCmd(moduleName, v.Name, modulesDir)
			ExecCommand(commmandString)
		}
	}
	return true
}
