package main

import (
	"strconv"
)

// Map/Struct for a new session.
type session struct {
	SessionID int
	UserID    string
	IP        string
}

// Defining global variables.
// The sessions array isn't in a database since it's only meant to exist in the cache.
var capacity = 50
var sessions [50]session
var NextAvailableSlot = 0 // Identifies the next available index in the array.

func newSessionID(UserID string, IP string) string {
	if NextAvailableSlot != capacity { // If the server has not reached capacity.
		SessionID := NextAvailableSlot // NextAvailableSlot is just the next index, which is unique.
		sessions[NextAvailableSlot] = session{SessionID, UserID, IP}
		NextAvailableSlot++
		return strconv.Itoa(SessionID)
	} else {
		return "FULL"
	}
}
