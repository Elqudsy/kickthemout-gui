package kickthemout

import (
	"fmt"
)

const (
	macDownload     = "https://nmap.org/dist/nmap-7.70.dmg"
	windowsDownload = "https://nmap.org/dist/nmap-7.70-setup.exe"
	errorMessage    = `Oops! This program depends on a popular network utility called Nmap, without it this program won't run. 
If you do have Nmap and are seeing this message you may just have to put the location of Nmap on your path. Would you like
to install Nmap?`
	errorMessageLinux = `Oops! This program depends on a popular network utility called Nmap, without it this program won't run. 
If you do have Nmap and are seeing this message you may just have to put the location of Nmap on your path. To install Nmap simply run 
CentOS: yum install nmap
Debian: apt-get install nmap 
Ubuntu sudo apt-get install nmap`
)

func Test() {
	fmt.Println(macDownload)
}
