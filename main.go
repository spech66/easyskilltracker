package main

import (
	"flag"
	"fmt"

	"github.com/spech66/easyskilltracker/route"
)

var flagBind string
var flagData string

func init() {
	flag.StringVar(&flagBind, "bind", ":8045", "ip:port to bind the server to")
	flag.StringVar(&flagData, "data", "data", "the data folder")
	flag.Parse()
}

func main() {
	fmt.Println("Easy Skill Tracker")
	fmt.Println("================")
	fmt.Println("Binding server to", flagBind)
	fmt.Println("Reading data from", flagData)

	router := route.Init(flagData)

	// Start the server
	router.Logger.Fatal(router.Start(flagBind))
}
