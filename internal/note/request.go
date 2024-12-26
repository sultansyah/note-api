package note

type CreateNoteRequest struct {
	Note     string `json:"note"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
	Category string `json:"category"`
	Tags     string `json:"tags"`
}

type GetNoteRequest struct {
	Id string `uri:"id"`
}
