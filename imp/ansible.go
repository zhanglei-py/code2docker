package imp

import (
	"os"
	"path/filepath"
)

func AnsbileCmd (moduleName string, roleName string) (string, error) {
	dir, _ := os.Executable()
	ExecPath := filepath.Dir(dir)
	ansibleCmd := "/opt/nset/py3/bin/ansible-playbook "
	ansibleInventory := "-i "+ ExecPath + "/inventory.ini "
	ansibleOpt := "-b "
	ansiblePlaybook := ExecPath + "/install.yaml "
	ansibleConf := "-e @config_local.yaml "
	ansibleModule := "-e module_name=" + moduleName + " "
	ansibleRole := "-e module_role_name=" + roleName
	commandString := ansibleCmd + ansibleInventory + ansibleOpt + ansiblePlaybook + ansibleConf + ansibleModule + ansibleRole
	return commandString, nil
}