package note

type CreateNoteRequest struct {
	Note     string `json:"note" example:"Buy groceries"`
	Status   string `json:"status" example:"pending"`
	Priority string `json:"priority" example:"high"`
	Category string `json:"category" example:"personal"`
	Tags     string `json:"tags" example:"shopping,home"`
}

type GetNoteRequest struct {
	Id string `uri:"id"`
}
