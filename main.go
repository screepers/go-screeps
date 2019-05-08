package main

import (
	"log"
	"screepers/go-screeps/config"
	"screepers/go-screeps/screeps"
)

func main() {
	conf := config.NewConfig()
	server := conf.Servers["screepsplus"]
	//server.Username = "invalid"
	c := screeps.NewClient(server)

	// resp1, _ := c.Version()
	// log.Printf("%v", resp1)

	// resp2, _ := c.Authmod()
	// log.Printf("%v", resp2)

	resp, err := c.AuthMe()
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("%v", resp)
	}
}
