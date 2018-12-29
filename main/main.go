package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"path/filepath"

	kto "github.com/cantasaurus/kickthemout-gui"
	"github.com/zserge/webview"
)

func startServer(setupResp bool) string {
	//Listen on localhost as this is how our ui is going to be sent to the webview.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	//Goroutine with only 1 http handler as that's all that's going to be used. If there
	//was an error during setup simply serve the error page instead.
	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if len(path) > 0 && path[0] == '/' {
				path = path[1:]
			}
			if path == "" && setupResp == true {
				path = "gui/index.html"
			} else if path == "" && setupResp == false {
				path = "gui/error.html"
			}
			if bs, err := Asset(path); err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.Header().Add("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
				io.Copy(w, bytes.NewBuffer(bs))
			}
		})
		log.Fatal(http.Serve(ln, nil))
	}()

	return "http://" + ln.Addr().String()
}

func startWebView() webview.WebView{
	response, _ := kto.CheckAll()
	fmt.Println("Response: ", response)
	//If response is true a different html page is served prompting the user to either install nmap or run the program as root. 
	url := startServer(response)
	w := webview.New(webview.Settings{
		Title:     "Kick Them Out",
		URL:       url,
		Width:     600,
		Height:    600,
		Resizable: true,
	})
	return w
}

func main() {
	localNet := kto.DefaultLocalNetwork()

	for i, elem := range localNet.MyIPs {
		if i == 0 {
			fmt.Println(elem)
		}
	}

	fmt.Println(localNet.MyHostName)

	for k, v := range localNet.MyMacs {
		fmt.Println("k:", k, "v:", v)
	}

	output := kto.NmapLocalNetScan("10.0.0.*")
	fmt.Println(output)

	w := startWebView()
	defer w.Exit()
	w.Run()
}
