package main

import (
	"io"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func connect(wsURL string) error {
	log.Println("Connecting to:", wsURL)

	// Connect to remote
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		w, err := conn.NextWriter(websocket.BinaryMessage)
		if err != nil {
			return err
		}
		if _, err := io.Copy(w, os.Stdin); err != nil {
			return err
		}
	}
}
