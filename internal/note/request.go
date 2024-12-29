package note

type CreateNoteRequest struct {
	Note     string `json:"note" binding:"required" example:"Buy groceries"`
	Status   string `json:"status" binding:"required" example:"pending"`
	Priority string `json:"priority" binding:"required" example:"high"`
	Category string `json:"category" binding:"required" example:"personal"`
	Tags     string `json:"tags" binding:"required" example:"shopping,home"`
}

type GetNoteRequest struct {
	Id string `uri:"id"`
}
