package imp

import (
	"strings"
)

func List2MapString(listString, offset string) []string {
	var mapString []string
	if len(offset) == 0 {
		offset = ","
	}
	s := strings.Split(listString, offset)
	if len(s) == 0 {
		mapString = append(mapString, listString)
	} else {
		for _, v := range s {
			mapString = append(mapString, v)
		}
	}
	return mapString
}

func SshHosts(moduleName, hostList string) ([]string, error) {
	var sshHosts []string
	sshHosts, err := ReadInventoryIni("inventory.ini", moduleName)
	if len(hostList) != 0 && len(moduleName) != 0 {
		sshHosts = append(List2MapString(hostList, ""), sshHosts...)
	} else if len(hostList) != 0 && len(moduleName) == 0 {
		sshHosts = List2MapString(hostList, "")
	}
	return sshHosts, err
}
