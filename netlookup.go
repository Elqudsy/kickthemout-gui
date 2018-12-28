package kickthemout

import (
	"fmt"
	"net"
	"os"
)

//LocalNetwork - Struct containing info relating to the users local network.
type LocalNetwork struct {
	MyIPs              []string
	MyMacs             map[string]string
	MyHostName         string
	AvailableIps       []string
	AvailableHostNames []string
	AvailableMacs      []string
	TargetIps          []string
}

//Get Mac adresses and vendor info (uses oui database from a text file)
func getMacAddrs() map[string]string {
	macs := make(map[string]string)
	interfaces, err := net.Interfaces()
	if err != nil {
		os.Exit(1)
	}

	for _, ifs := range interfaces {
		mac := ifs.HardwareAddr.String()
		if mac != "" {
			fmt.Println(mac)
			v, _ := db.VendorLookup(mac)
			//If vendor can't be found it returns "" which is fine for our case.
			macs[mac] = v
		}
	}
	return macs
}

//DefaultLocalNetwork - Returns a default LocalNetwork struct with your local ips, mac addresses and hostnames of the machines
//in your local network.
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
		MyMacs:     getMacAddrs(),
		MyHostName: myHostName,
	}
}
