package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ConnectionContainer struct {

    Wsc *websocket.Conn

}

func CreateServer() *ConnectionContainer {

	cont := &ConnectionContainer{}

	http.HandleFunc("/wsServer", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil{
			fmt.Println(err.Error())
		}
        cont.Wsc = conn
        fmt.Println("connection stablished")

	})

	go http.ListenAndServe(":8080", nil)

	return cont
}

func (cc ConnectionContainer) Write (str string){

	if cc.Wsc != nil {
		cc.Wsc.WriteMessage(1, []byte(str))
	} else {
		fmt.Println("conn null")
	}

}