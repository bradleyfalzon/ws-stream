package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	listen := flag.String("listen", "", "Run as a server listening addr:port")
	wsURL := flag.String("url", "", "Websocket URL to connect to")
	flag.Parse()

	if *listen != "" {
		// Listen
		http.HandleFunc("/ws", wsHandler)
		log.Printf("Listening on ws://%v/ws", *listen)
		log.Fatal(http.ListenAndServe(*listen, nil))
	} else if *wsURL != "" {
		// Connect
		log.Fatal(connect(*wsURL))
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}

}
