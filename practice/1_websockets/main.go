package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// define a reader which will listen for new messages
// being sent to our Websocket endpoint

func reader(conn *websocket.Conn) {
	for {

		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return

		}
		// print out that message for clarify
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Panicln(err)
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a websocket
	// connection

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)

}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
