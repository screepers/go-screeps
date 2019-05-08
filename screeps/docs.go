// Package screeps - Screeps API Library
package screeps

import (
	"log"

	"github.com/screepers/go-screeps/config"
)

// Example usage
func Example() {
	conf := config.NewConfig()
	server := conf.Servers["main"]
	c := NewClient(server)
	resp, err := c.AuthMe()
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("%v", resp)
	}
}
