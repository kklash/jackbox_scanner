package main

import (
	"crypto/rand"
	"math/big"
	"strings"
	"sync"
)

var (
	ALPHABET           = "abcdefghijklmnopqrstuvwxyz"
	generatedRoomCodes = new(sync.Map)
)

// Generate a random unique room code. If we regenerate an old
// one, ignore it and try again.
func genRoomCode() string {
	// a - z = 97 - 122
	roomCode := ""
	for i := 0; i < 4; i++ {
		b, _ := rand.Int(rand.Reader, big.NewInt(25))
		i := byte(b.Int64() + 97)
		roomCode += string(i)
	}

	// If this one has already been generated before, try again.
	if _, ok := generatedRoomCodes.Load(roomCode); ok {
		return genRoomCode()
	}

	generatedRoomCodes.Store(roomCode, true)
	return roomCode
}

// Verify a room code is valid.
func isValidRoomCode(roomCode string) bool {
	roomCode = strings.ToLower(roomCode)
	if len(roomCode) != 4 {
		return false
	}

	for _, c := range roomCode {
		if !strings.ContainsRune(ALPHABET, c) {
			return false
		}
	}

	return true
}
