package note

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sultansyah/note-api/internal/helper"
)

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
	var input CreateNoteRequest

	if !helper.BindAndValidateJSON(c, &input) {
		return
	}

	userId := c.MustGet("userId").(int)

	note, err := n.NoteService.Create(c.Request.Context(), input, userId)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success created new note",
		Data:    note,
	})
}

func (n *NoteHandlerImpl) Delete(c *gin.Context) {
	var input GetNoteRequest

	if !helper.BindAndValidateURi(c, &input) {
		return
	}

	userId := c.MustGet("userId").(int)

	err := n.NoteService.Delete(c.Request.Context(), input, userId)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success delete note",
		Data:    nil,
	})
}

func (n *NoteHandlerImpl) Edit(c *gin.Context) {
	var inputId GetNoteRequest
	if !helper.BindAndValidateURi(c, &inputId) {
		return
	}

	var inputData CreateNoteRequest
	if !helper.BindAndValidateJSON(c, &inputData) {
		return
	}

	userId := c.MustGet("userId").(int)

	note, err := n.NoteService.Edit(c.Request.Context(), inputData, inputId, userId)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success edit note",
		Data:    note,
	})
}

func (n *NoteHandlerImpl) FindAll(c *gin.Context) {
	userId := c.MustGet("userId").(int)

	notes, err := n.NoteService.FindAll(c.Request.Context(), userId)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get all data",
		Data:    notes,
	})
}

func (n *NoteHandlerImpl) FindById(c *gin.Context) {
	var input GetNoteRequest

	if !helper.BindAndValidateURi(c, &input) {
		return
	}

	note, err := n.NoteService.FindById(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get data",
		Data:    note,
	})
}
