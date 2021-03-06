package spark

import "time"

// Message represents a spark message
type Message struct {
	RoomId        string `json:"roomId,omitempty"`
	ToPersonEmail string `json:"toPersonEmail,omitempty"`
	Text          string `json:"text,omitempty"`
	Markdown      string `json:"markdown,omitempty"`
	Files         string `json:"files,omitempty"`
}

// Room represents a spark room
type Room struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Type         string    `json:"type"`
	IsLocked     bool      `json:"isLocked"`
	TeamId       string    `json:"teamId,omitempty"`
	CreatorId    string    `json:"creatorId"`
	Created      time.Time `json:"created"`
	LastActivity time.Time `json:"lastActivity"`
}

// Rooms represents a spark room list
type Rooms struct {
	Items []Room `json:"items"`
}
