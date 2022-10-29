package main

import (
	"bufio"
	"fmt"
	"net"
	"tcp-client/utils"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, err := utils.Decode(reader)
			if err != nil {
				break
			}
			fmt.Println(msg)
		}
	}()
	defer conn.Close()
	for {
		msg := ""
		fmt.Scanf("%s", &msg)
		if msg == "exit" {
			return
		}
		data := utils.Encode(msg)
		conn.Write(data)
	}
}
