package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	defer conn.Close()

	if err := read(conn); err != nil {
		log.Println(err)
	}
}

func read(conn *websocket.Conn) error {
	for {
		_, r, err := conn.NextReader()
		if err != nil {
			return err
		}
		if _, err := io.Copy(os.Stdout, r); err != nil {
			return err
		}
	}
}
