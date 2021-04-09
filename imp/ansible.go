package imp

import (
	"fmt"
	"os"
	"path/filepath"
)

func AnsibleCmd(moduleName, roleName, workDir string) (string, error) {
	var moduleDir string
	dir, _ := os.Executable()
	if len(workDir) > 0 {
		moduleDir = workDir
	} else {
		moduleDir = filepath.Dir(dir)
	}
	ansibleCmd := "/opt/nset/py3/bin/ansible-playbook"
	ansibleInventory := "-i " + moduleDir + "/inventory.ini"
	ansibleOpt := "-b"
	ansiblePlaybook := moduleDir + "/install.yaml"
	ansibleConf := "-e @config_local.yaml"
	ansibleModule := "-e module_name=" + moduleName
	ansibleRole := "-e module_role_name=" + roleName
	ansibleWorkDir := "-e WORKDIR=" + moduleDir
	commandString := fmt.Sprintf("%s %s %s %s %s %s %s %s", ansibleCmd, ansibleInventory, ansibleOpt, ansiblePlaybook, ansibleConf, ansibleModule, ansibleRole, ansibleWorkDir)
	return commandString, nil
}
