package note

import "github.com/gin-gonic/gin"

type NoteHandler interface {
	Create(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
}

type NoteHandlerImpl struct {
	NoteService NoteService
}

func NewNoteHandler(noteService NoteService) NoteHandler {
	return &NoteHandlerImpl{NoteService: noteService}
}

func (n *NoteHandlerImpl) Create(c *gin.Context) {
	panic("unimplemented")
}

func (n *NoteHandlerImpl) Delete(c *gin.Context) {
	panic("unimplemented")
}

func (n *NoteHandlerImpl) Edit(c *gin.Context) {
	panic("unimplemented")
}

func (n *NoteHandlerImpl) FindAll(c *gin.Context) {
	panic("unimplemented")
}

func (n *NoteHandlerImpl) FindById(c *gin.Context) {
	panic("unimplemented")
}
