package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	cmd   string // user's benchmark command
	token string
)

func main() {

	flag.StringVar(&cmd, "cmd", "", "benchmark command")
	flag.StringVar(&token, "token", "", "franz api key")

	flag.Usage = usage
	flag.Parse()

	// cmd is a mandatory flag
	if cmd == "" {
		usage()
		os.Exit(0)
	}

	fKey := os.Getenv("FRANZ_API_KEY")
	if fKey == "" {
		if token == "" {
			fmt.Println("franz api key is not set")
			usage()
		} else {
			fKey = token
		}
	}
}

func usage() {
	fmt.Println("franz --cmd=\"hi\" ")
	os.Exit(0)
}
