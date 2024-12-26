package note

import (
	"context"
	"database/sql"

	"github.com/sultansyah/note-api/internal/helper"
)

type NoteRepository interface {
	Create(ctx context.Context, tx *sql.Tx, note Note) (Note, error)
	Edit(ctx context.Context, tx *sql.Tx, note Note) (Note, error)
	Delete(ctx context.Context, tx *sql.Tx, noteId int) error
	Get(ctx context.Context, tx *sql.Tx, noteId int) (Note, error)
	Gets(ctx context.Context, tx *sql.Tx) ([]Note, error)
}

type NoteRepositoryImpl struct {
}

func NewNoteRepository() NoteRepository {
	return &NoteRepositoryImpl{}
}

func (n *NoteRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, note Note) (Note, error) {
	sql := "insert into notes (user_id, note, status, priority, category, tags) VALUES (?, ?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, sql, note.UserId, note.Note, note.Status, note.Priority, note.Category, note.Tags)
	if err != nil {
		return Note{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Note{}, err
	}
	note.Id = int(id)

	return note, nil
}

func (n *NoteRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, noteId int) error {
	sql := "delete from notes where id = ?"

	_, err := tx.ExecContext(ctx, sql, noteId)
	if err != nil {
		return err
	}

	return nil
}

func (n *NoteRepositoryImpl) Edit(ctx context.Context, tx *sql.Tx, note Note) (Note, error) {
	sql := "update notes set note = ?, status = ?, priority = ?, category = ?, tags = ?"

	result, err := tx.ExecContext(ctx, sql, note.Note, note.Status, note.Priority, note.Category, note.Tags)
	if err != nil {
		return Note{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Note{}, err
	}

	note.Id = int(id)

	return note, nil
}

func (n *NoteRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, noteId int) (Note, error) {
	sql := "select id, user_id, note, status, priority, category, tags, created_at, updated_at from notes where id = ?"

	row, err := tx.QueryContext(ctx, sql, noteId)
	if err != nil {
		return Note{}, err
	}
	defer row.Close()

	var note Note

	if row.Next() {
		err := row.Scan(&note)
		if err != nil {
			return Note{}, err
		}

		return note, nil
	}

	return Note{}, helper.ErrNotFound
}

func (n *NoteRepositoryImpl) Gets(ctx context.Context, tx *sql.Tx) ([]Note, error) {
	sql := "select id, user_id, note, status, priority, category, tags, created_at, updated_at from notes"

	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		return []Note{}, err
	}
	defer rows.Close()

	var notes []Note

	if rows.Next() {
		for rows.Next() {
			var note Note

			if err := rows.Scan(&note.Id, &note.UserId, &note.Note, &note.Status, &note.Priority, &note.Category, &note.Tags, &note.CreatedAt, &note.UpdatedAt); err != nil {
				return []Note{}, err
			}

			notes = append(notes, note)
		}

		return notes, nil
	}

	return []Note{}, helper.ErrNotFound
}
