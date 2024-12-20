package note

import "time"

type Note struct {
	Id        int       `json:"id"`
	Note      int       `json:"note"`
	Status    string    `json:"status"`
	Priority  string    `json:"priority"`
	Category  string    `json:"category"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
