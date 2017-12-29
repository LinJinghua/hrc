package main

import (
	"os"

	"github.com/LinJinghua/hrc/service"
	flag "github.com/spf13/pflag"
)

const (
	// PORT .
	PORT string = "8080"
)

func server()  {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}

func main() {
	server()
}
