package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	BASE_ROOM_URL = "https://ecast.jackboxgames.com/room"
)

func roomCodeUrl(code string) string {
	return BASE_ROOM_URL + "/" + code
}

type Result struct {
	roomInfo *RoomInfo
	success  bool
	err      error
}

// Determine if a room code is valid or not.
func checkRoomCode(roomCode string) (result Result) {
	var resp *http.Response
	resp, result.err = http.Get(roomCodeUrl(roomCode))
	if result.err != nil {
		return
	}

	defer resp.Body.Close()

	if result.roomInfo, result.err = parseRoomInfo(resp.Body); result.err != nil {
		return
	}

	// Not returned on bad response, ensure it is populated
	if result.roomInfo.RoomCode == "" {
		result.roomInfo.RoomCode = strings.ToUpper(roomCode)
	}

	if DEBUG {
		fmt.Printf("\nRoom code: %s\n", roomCode)
		fmt.Printf("Response Status: %s", resp.Status)
		pretty, _ := json.MarshalIndent(result.roomInfo, "", "  ")
		fmt.Println("\n" + string(pretty))
	}

	result.success = resp.StatusCode == http.StatusOK &&
		isAcceptableRoomInfo(result.roomInfo)

	return
}
