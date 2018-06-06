package main

import (
	"fmt"
	"os"
)

var (
	// Username ...
	Username string
	// PrivateKey ...
	PrivateKey string
)

// CheckArgs ...
func CheckArgs() {
	// check if username is empty
	if Username == "" {
		msg := "Note:\n\tThis app is experimental, I'm just playing with a whisper protocol.\n"
		msg += "\tRun first go-ethereum 'geth --testnet --nodiscover --shh --rpc'\n"
		msg += "\tRun 'pchat --help' for more flags information.\n\n"
		msg += "USAGE:\n\tpchat -username [argument] [flag option] [argument]\n\n"
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}
}
