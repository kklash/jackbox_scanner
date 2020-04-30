package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Response returned by jackbox servers when an active room is found.
// If no room is found, only a { success, error } object is returned.
type RoomInfo struct {
	App               string `json:"apptag"`
	AudienceEnabled   bool   `json:"audienceEnabled"`
	AudienceMembers   int    `json:"numAudience"`
	Error             string `json:"error"`
	JoinAs            string `json:"joinAs"`
	PasswordProtected bool   `json:"requiresPassword"`
	RoomCode          string `json:"roomid"`
	Server            string `json:"server"`
}

// Parse a RoomInfo from a reader.
func parseRoomInfo(r io.Reader) (*RoomInfo, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	roomInfo := new(RoomInfo)
	if err = json.Unmarshal(data, roomInfo); err != nil {
		return nil, err
	}

	return roomInfo, nil
}

// Make sure the room info jives with the CLI flags.
func isAcceptableRoomInfo(roomInfo *RoomInfo) bool {
	return roomInfo != nil &&
		(PASSWORDS || !roomInfo.PasswordProtected) &&
		(!ONLY_PLAYABLE || roomInfo.JoinAs == "player")
}
