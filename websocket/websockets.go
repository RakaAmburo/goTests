// websockets.go
package main

import (
	"./ws"
	"fmt"

	"net/http"
	"time"
)

func main() {

	go testserver()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}

func testserver() {

	cont := ws.CreateServer()
	fmt.Println("preparando")
	for {

		time.Sleep(5 * time.Second)

		cont.Write("probanding wesocketing")

	}

}
