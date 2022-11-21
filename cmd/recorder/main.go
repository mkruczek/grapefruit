package main

import "grapefruit/pkg/app/recorder"

func main() {

	server := recorder.NewServer()
	server.Start()
}
