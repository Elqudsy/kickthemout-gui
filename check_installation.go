package kickthemout

import (
	"runtime"
	"os"
	"os/exec"
	"errors"
)

var downLoadLink = map[string]string{
	"darwin": "https://nmap.org/dist/nmap-7.70.dmg",
	"windows": "https://nmap.org/dist/nmap-7.70-setup.exe",
}

const (
	errorMessage = `Oops! This program depends on a popular network utility called Nmap, without it this program won't run. 
If you do have Nmap and are seeing this message you may just have to put the location of Nmap on your path. Would you like
to install Nmap?`
	errorMessageLinux = `Oops! This program depends on a popular network utility called Nmap, without it this program won't run. 
If you do have Nmap and are seeing this message you may just have to put the location of Nmap on your path. To install Nmap simply run 
CentOS: yum install nmap
Debian: apt-get install nmap 
Ubuntu sudo apt-get install nmap`
)

func IsRoot() (bool, error) {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		if os.Geteuid() != 0 {
			return false, errors.New("Program must be run as root. Please rerun the program as the root user or using sudo.")
		}
	} else if runtime.GOOS == "windows" {
		return false, nil
	}
	return true, nil
}

func NmapAvailable() (bool, error) {
	_, err := exec.Command("nmap").CombinedOutput()
	if err != nil {
		if runtime.GOOS == "linux"{
			return false, errors.New(errorMessageLinux)
		} else {
			return false, errors.New(errorMessage)
		}
	}
	return true, nil
}

func CheckAll() (bool, error) {
	_, rootErr := IsRoot()
	if rootErr != nil {
		return false, rootErr
	}
	_, nmapErr := NmapAvailable()
	if nmapErr != nil{
		return false, nmapErr
	}	
	return true, nil 
}
