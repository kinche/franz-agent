package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/kinche/franz-agent/ci"
	"github.com/kinche/franz-agent/client"
)

var (
	cmd   string // user's benchmark command
	token string // franz api key
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

	cinfo := ci.Detect()

	command := strings.Split(cmd, " ")
	args := command[1:]

	e := exec.Command(command[0], args...)
	s, _ := e.Output()

	fmt.Print(string(s))

	client.SendReport(s, cinfo.Get())

	// don't fail the CI pipeline
	os.Exit(0)
}

func usage() {
	fmt.Println("franz --cmd=\"hi\" ")
	os.Exit(0)
}
