package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	kto "github.com/cantasaurus/kickthemout"
	"github.com/zserge/webview"
)

func startServer() string {
	//Listen on localhost as this is how our ui is going to be sent to the webview.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if len(path) > 0 && path[0] == '/' {
				path = path[1:]
			}
			if path == "" {
				path = "gui/index.html"
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

func main() {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		if os.Geteuid() != 0 {
			fmt.Println("Program must be run as root. Please rerun the program as the root user or using sudo.")
			os.Exit(1)
		}
	}
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

	output := kto.NmapCall("10.0.0.*")
	fmt.Println(output)

	url := startServer()
	w := webview.New(webview.Settings{
		Title:     "Asset Test",
		URL:       url,
		Width:     600,
		Height:    600,
		Resizable: true,
	})
	defer w.Exit()
	w.Run()
}
