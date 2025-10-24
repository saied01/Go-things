package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// upgrader used to upgrade HTTP connections to websocket connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade http conn to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error upgrading:", err)
		return
	}

	defer conn.Close()

	// listen for incoming messages
	for {
		// read msg from client
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error reading message:", err)
			break
		}
		fmt.Println("received: %s\\n", message)
		// echo msg back to client
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("error writing message:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("websocket server started on: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error starting server:", err)
	}
}
