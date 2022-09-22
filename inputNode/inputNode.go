//input node takes messagenfrom user and sends it to random node
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// json obj for http messages
// define ports of layer two nodes
// ask for message from user
// marshall into json
// send to next available in list using HTTP post

var PortsRange = struct {
	min int
	max int
}{300, 400} //TODO figure out ideal range

func main() {

	go func() {
		f := http.HandlerFunc(func(w http.ResponseWriter, resp *http.Request) {
			clientMessage := make([]byte, 100)
			resp.Body.Read(clientMessage) //TODO: handle error
			fmt.Println("response: " + string(clientMessage))
			w.Write([]byte("response: " + string(clientMessage))) //TODO: handle error
		})
		http.ListenAndServe(":3000", f) //TODO: handle error
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("enter command: ")
		line, err := reader.ReadString('\n') //TODO: handle error
		m := strings.NewReader(line)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(line)
		http.Post("http://localhost:3000", "test/plain", m) //TODO: handle error and Resp
	}

}
