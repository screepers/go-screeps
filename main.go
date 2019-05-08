package main

import (
	"log"

	"github.com/screepers/go-screeps/config"
	"github.com/screepers/go-screeps/screeps"
)

func main() {
	conf := config.NewConfig()
	server := conf.Servers["main"]
	//server.Username = "invalid"
	c := screeps.NewClient(server)

	// resp1, _ := c.Version()
	// log.Printf("%v", resp1)

	// resp2, _ := c.Authmod()
	// log.Printf("%v", resp2)

	resp, err := c.RoomHistory("W1N1", 31768500, "")
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("%v", resp)
	}
}
