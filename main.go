package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Checks given room codes manually, in parallel.
// Parallelism of len(roomCodes) threads.
func checkRoomCodesManual(roomCodes []string) error {
	channel := make(chan Result, len(roomCodes))
	for i := 0; i < len(roomCodes); i++ {
		go func(i int) {
			channel <- checkRoomCode(roomCodes[i])
		}(i)
	}

	for _, roomCode := range roomCodes {
		result := <-channel
		fmt.Printf("%s: ", roomCode)
		if result.err != nil {
			return result.err
		} else if result.success {
			fmt.Println("**** FOUND ROOM ****")
		} else {
			fmt.Println("EMPTY")
		}
	}

	return nil
}

// Randomly generate room codes and test them for activity in parallel.
// Parallelism of N_WORKERS threads.
func checkRoomCodesRandom() error {
	channel := make(chan Result, N_WORKERS*10)
	done := false
	for i := 0; i < N_WORKERS; i++ {
		go func() {
			for {
				if done {
					return
				}
				channel <- checkRoomCode(genRoomCode())
			}
		}()
	}

	for codesChecked := 1; ; codesChecked++ {
		result := <-channel

		if result.err != nil {
			done = true
			return result.err
		} else if result.success {
			fmt.Printf("\nFOUND ROOM CODE: %s\n", strings.ToUpper(result.roomInfo.RoomCode))
			pretty, _ := json.MarshalIndent(result.roomInfo, "", "  ")
			fmt.Println(string(pretty))
			if FIND_FIRST {
				done = true
				return nil
			}
		} else {
			fmt.Printf("\rRoom codes checked: %d", codesChecked)
		}
	}
}

func run() error {
	// User can manually supply room codes to be checked as positional CLI arguments.
	if len(argv) > 0 {
		return checkRoomCodesManual(argv)
	}

	return checkRoomCodesRandom()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
