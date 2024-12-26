package note

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/sultansyah/note-api/internal/helper"
)

type NoteService interface {
	Create(ctx context.Context, input CreateNoteRequest, userId int) (Note, error)
	Edit(ctx context.Context, inputData CreateNoteRequest, inputId GetNoteRequest, userId int) (Note, error)
	Delete(ctx context.Context, input GetNoteRequest, userId int) error
	FindById(ctx context.Context, input GetNoteRequest) (Note, error)
	FindAll(ctx context.Context, userId int) ([]Note, error)
}

type NoteServiceImpl struct {
	NoteRepository NoteRepository
	DB             *sql.DB
}

func NewNoteService(noteRepository NoteRepository, DB *sql.DB) NoteService {
	return &NoteServiceImpl{NoteRepository: noteRepository, DB: DB}
}

func (n *NoteServiceImpl) Create(ctx context.Context, input CreateNoteRequest, userId int) (Note, error) {
	tx, err := n.DB.Begin()
	if err != nil {
		return Note{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	note := Note{
		UserId:   userId,
		Note:     input.Note,
		Status:   input.Status,
		Priority: input.Priority,
		Category: input.Category,
		Tags:     input.Tags,
	}

	note, err = n.NoteRepository.Create(ctx, tx, note)
	if err != nil {
		return Note{}, err
	}

	return note, nil
}

func (n *NoteServiceImpl) Delete(ctx context.Context, input GetNoteRequest, userId int) error {
	tx, err := n.DB.Begin()
	if err != nil {
		return err
	}

	defer helper.HandleTransaction(tx, &err)

	noteId, err := strconv.Atoi(input.Id)
	if err != nil {
		return err
	}

	note, err := n.NoteRepository.FindById(ctx, tx, noteId)
	if err != nil {
		return err
	}
	if note.Id == 0 {
		return helper.ErrNotFound
	}

	if note.UserId != userId {
		return helper.ErrForbidden
	}

	err = n.NoteRepository.Delete(ctx, tx, noteId)
	if err != nil {
		return err
	}

	return nil
}

func (n *NoteServiceImpl) Edit(ctx context.Context, inputData CreateNoteRequest, inputId GetNoteRequest, userId int) (Note, error) {
	tx, err := n.DB.Begin()
	if err != nil {
		return Note{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	noteId, err := strconv.Atoi(inputId.Id)
	if err != nil {
		return Note{}, err
	}

	note, err := n.NoteRepository.FindById(ctx, tx, noteId)
	if err != nil {
		return Note{}, err
	}
	if note.Id == 0 {
		return Note{}, helper.ErrNotFound
	}
	if note.Id != noteId || note.UserId != userId {
		return Note{}, helper.ErrForbidden
	}

	note = Note{
		Id:       noteId,
		UserId:   userId,
		Note:     inputData.Note,
		Status:   inputData.Status,
		Priority: inputData.Priority,
		Category: inputData.Category,
		Tags:     inputData.Tags,
	}

	note, err = n.NoteRepository.Edit(ctx, tx, note)
	if err != nil {
		return Note{}, err
	}

	return note, nil
}

func (n *NoteServiceImpl) FindAll(ctx context.Context, userId int) ([]Note, error) {
	tx, err := n.DB.Begin()
	if err != nil {
		return []Note{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	notes, err := n.NoteRepository.FindAll(ctx, tx, userId)
	if err != nil {
		return []Note{}, err
	}
	if len(notes) == 0 {
		return []Note{}, helper.ErrNotFound
	}

	return notes, nil
}

func (n *NoteServiceImpl) FindById(ctx context.Context, input GetNoteRequest) (Note, error) {
	tx, err := n.DB.Begin()
	if err != nil {
		return Note{}, err
	}

	defer helper.HandleTransaction(tx, &err)

	noteId, err := strconv.Atoi(input.Id)
	if err != nil {
		return Note{}, err
	}

	note, err := n.NoteRepository.FindById(ctx, tx, noteId)
	if err != nil {
		return Note{}, err
	}
	if note.Id == 0 {
		return Note{}, helper.ErrNotFound
	}

	return note, nil
}
