package kickthemout

import (
	"os/exec"
	"runtime"
)

//NmapCall - Get available local ips on the users network.
func NmapLocalNetScan(ip string) string {
	args := []string{"nmap", "-sP", ip}
	data, err := exec.Command("sudo", args...).CombinedOutput()
	if err != nil {
		if runtime.GOOS == "linux" {
			return errorMessageLinux
		}
		return errorMessage
	}
	return string(data)
}


