package kickthemout

import (
	Oui "github.com/dutchcoders/go-ouitools"
)

//Load Oui database from file. Located in main.
func loadOui(fileLoc string) *Oui.OuiDb {
	return Oui.New(fileLoc)
}
