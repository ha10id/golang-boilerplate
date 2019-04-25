package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"./db"
	"./server"

	"./config"
)

func init() {
	log.Println("WELCOME!")

}

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*enviroment)
	db.Init()
	server.Init()
}
