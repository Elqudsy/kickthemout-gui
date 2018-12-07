package kickthemout

import (
	"fmt"
	"net"
	"os"
)

// Struct containing info relating to the users local network.
type LocalNetwork struct {
	MyIPs         []string
	MyMacs        map[string]string
	MyHostName    string
	AvailableIps  []string
	AvailableMacs []string
	TargetIps     []string
}

var db = loadOui("oui.txt")

//Get available local ips on the users network.
func GetAvailableIps() {
	fmt.Println("sdfsfds")
}

//Get Mac adresses and vendor info (uses oui database from a text file)
func getMacAddrs() []string {
	var macs []string
	interfaces, err := net.Interfaces()
	if err != nil {
		os.Exit(1)
	}

	for _, ifs := range interfaces {
		a := ifs.HardwareAddr.String()
		if a != "" {
			macs = append(macs, a)
		}
	}
	return macs
}

//Returns a default LocalNetwork struct with your local ips, mac address and hostname of the machine
//the program is being run on.
func DefaultLocalNetwork() *LocalNetwork {
	var (
		myIPs      []string
		myHostName string
	)

	myHostName, err := os.Hostname()
	if err != nil {
		fmt.Println("There was an error trying to get your hostname. Shutting down...")
		os.Exit(1)
	}

	addrs, err := net.LookupHost(myHostName)
	if err != nil {
		fmt.Println("There was an error trying to find your ip using the hostname. Shutting down...")
		os.Exit(1)
	}

	for _, a := range addrs {
		myIPs = append(myIPs, a)
	}

	return &LocalNetwork{
		MyIPs:      myIPs,
		MyMacs:     getMacAddr(),
		MyHostName: myHostName,
	}
}
