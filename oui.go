package kickthemout

import (
	"os"

	Oui "github.com/dutchcoders/go-ouitools"
)

//Initialize oui db. Good for now but change before release.
var db = loadOui()

//Load Oui database from file. Located in main.
func loadOui() *Oui.OuiDb {
	fileLoc, _ := os.Getwd()
	fileLoc += "/oui.txt"
	//TODO: RETURN ERROR IF FILE CANNOT BE FOUND
	return Oui.New(fileLoc)
}
