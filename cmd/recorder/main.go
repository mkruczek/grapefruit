package main

import (
	"flag"
	"grapefruit/pkg/app/recorder"
	"log"
)

func main() {

	var cfgPathFlag string
	flag.StringVar(&cfgPathFlag, "cfgpath", "test_config.json", "path to the configuration flag")
	flag.Parse()

	server, err := recorder.NewServer(cfgPathFlag)
	if err != nil {
		log.Fatalf("can't create new recorder server: %s", err)
	}
	server.Start()
}
