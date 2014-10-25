package main

import (
	"fmt"
	"log"
	"net"
)

func echo(connection net.Conn) bool {
	defer connection.Close();
	var buffer[2048] byte;
	for {
		n, err := connection.Read(buffer[0:]);
		if err != nil {
			connection.Write([]byte(err.Error()));
			return false;
		}
		packet := string(buffer[:n]);
		fmt.Printf("%s\n", packet);
		if packet == "quit" || packet == "exit" {
			connection.Write([]byte("Bye Bye!"));
			return true;
		}
		connection.Write(buffer[:n]);
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:6969");
	if err != nil {
		log.Fatal(err);
	}
	defer listener.Close();
	for {
		connection, err := listener.Accept();
		if err != nil {
			log.Fatal(err);
		}
		go echo(connection);
	}
}
