package note

import (
	"time"

	"github.com/sultansyah/note-api/internal/user"
)

type Note struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Note      string    `json:"note"`
	Status    string    `json:"status"`
	Priority  string    `json:"priority"`
	Category  string    `json:"category"`
	Tags      string    `json:"tags"`
	User      user.User `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
