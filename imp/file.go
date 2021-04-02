package imp

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadVarsFile (fileName string, moduleName string) (string, error){
	var fd string
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
		return "",err
	} else {
		buff := bufio.NewReader(f)
		for {
			line, err := buff.ReadString('\n')
			//lineString := strings.Replace(line, "\n", "", -1)
			lineString := strings.Split(line,":")[0]
			if lineString != moduleName + "_vars" {
				fd = fd + line
			}
			if err == io.EOF {
				break
			}

		}
	}
	return string(fd),nil
}

func ReadInventoryIni (fileName, sectionName string)([]string, error) {
	var fd []string
	if len(sectionName) != 0 {
		sectionName = "[" +  sectionName + "_local]"
	} else {
		sectionName = "[all_ipaddrs]"
	}
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
		return fd, err
	} else {
		buff := bufio.NewReader(f)
		record := 0
		for {
			line, err := buff.ReadString('\n')
			if err != io.EOF {
				//line = strings.Replace(line, " ", "", -1)
				line := strings.Replace(line, "\n", "", -1)
				if len(line) > 0 {
					sectionTag := line[0:1] + line[len(line)-1:]
					if sectionName == line && record == 0 {
						record = 1
					}
					if sectionName != line && record == 1 && sectionTag == "[]" {
						record = 2
						break
					}
					if sectionName != line && record == 1 {
						fd = append(fd, strings.Split(line, " ")[0])
					}
				}
			} else {
				break
			}
		}
	}
	return fd,nil
}
