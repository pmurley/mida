package main

import (
	"mida/cmd/build_task"
	"mida/cmd/client"
	"mida/cmd/server"
)

func main() {
	build_task.BuildTask()
	server.Server()
	client.Client()
	return
}
