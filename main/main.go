package main

import (
	"fmt"

	kto "github.com/cantasaurus/kickthemout"
	"github.com/zserge/webview"
)

func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("Minimal webview example",
		"https://en.m.wikipedia.org/wiki/Main_Page", 800, 600, true)
	localNet := kto.DefaultLocalNetwork()
	for i, elem := range localNet.MyIPs {
		if i == 0 {
			fmt.Println(elem)
		}
	}
	fmt.Println(localNet.MyHostName)
	for _, elem := range localNet.MyMac {
		fmt.Println(elem)
	}
}
