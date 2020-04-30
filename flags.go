package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	DEBUG         bool
	FIND_FIRST    bool
	N_WORKERS     int
	ONLY_PLAYABLE bool
	PASSWORDS     bool

	argv []string
)

func init() {
	flag.BoolVar(&DEBUG, "verbose", false, "Enable verbose logging of each request.")
	flag.BoolVar(&FIND_FIRST, "first", false, "Terminate upon finding first valid room code.")
	flag.BoolVar(&ONLY_PLAYABLE, "playable", false, "Only count games open to joining as players.")
	flag.BoolVar(&PASSWORDS, "passwords", false, "Show games which are password-protected.")
	flag.IntVar(&N_WORKERS, "workers", 5, "Number of parallel workers checking rooms.")

	flag.Parse()

	argv = flag.Args()
	if len(argv) > 0 {
		for _, roomCode := range argv {
			if !isValidRoomCode(roomCode) {
				fmt.Fprintf(os.Stderr, "ERROR: Invalid room code '%s'\n", roomCode)
				os.Exit(1)
			}
		}
	}
}
