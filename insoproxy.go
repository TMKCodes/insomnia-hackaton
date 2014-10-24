package main

import (
	"log"
	"net"
	"time"
	"math/rand"
	"io/ioutil"
	"encoding/json"
)

type Configuration struct {
	Socket string
	Address string
	Mode string
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix());
	return rand.Intn(max - min) + min;
}

func proxy(connection net.Conn, configuration Configuration) bool {
	defer connection.Close();
	var abuffer [2048]byte;
	var bbuffer [2048]byte;
	for {
		n, err := connection.Read(abuffer[0:]);
		if err != nil {
			connection.Write([]byte(err.Error()));
			return false;
		}
		apacket := string(abuffer[:n]);
		if apacket == "quit" || apacket == "exit" {
			connection.Write([]byte("Bye, bye!"));
			return true;
		}
		apacket = strings.TrimSuffix(packet, "\n");
		split := strings.Split(packet, ":");
		if configuration.Mode == "insomnia" {
			time.Sleep(random(0, 512658235) * time.Millisecond);
		} else if configuration.Mode == "crazy" {
			
		}
		sconnection, err := net.Dial(configuration.Socket, split[0]);
		if err != nil {
			connection.Write([]byte(err.Error()));
			return false;
		}
		sconnection.Write([]byte(split[1]));
		n, err = sconnection.Read(bbuffer[0:]);
		if err != nil {
			connection.Write([]byte(err.Error()));
		}
		connection.Write(bbuffer[:n]);
	}
}

func main() {
	file, err := ioutil.ReadFile("proxy.conf");
	if err != nil {
		log.Fatal(err);
	}
	configuration := new(Configuration);
	err = json.Unmarshal(file, configuration);
	if err != nil {
		log.Fatal(err);
	}
	if configuration.Socket == "tcp" {
		listener, err := net.Listen(configuration.Socket, configuration.Address);
		if err != nil {
			log.Fatal(err);
		}
		defer listener.Close();
		for {
			connection, err := listener.Accept();
			if err != nil {
				log.Fatal(err);
			}
			go proxy(connection, configuration);
		}
	}
}
