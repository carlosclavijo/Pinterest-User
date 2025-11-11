package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/board"
	"github.com/google/uuid"
	"time"
)

const (
	QueryGetAllBoards = `SELECT id, user_id, name, description, visibility, pin_count, portrait, created_at, updated_at, deleted_at
						 FROM boards`
	QueryGetListBoards = `SELECT id, user_id, name, description, visibility, pin_count, portrait, created_at, updated_at, deleted_at
						  FROM boards
						  WHERE deleted_at IS NULL`
	QueryGetListBoardsByUserId = `SELECT id, name, description, visibility, pin_count, portrait, created_at, updated_at, deleted_at
								  FROM boards
								  WHERE user_id = $1 AND deleted_at IS NULL`
	QueryGetListBoardsByName = `SELECT id, user_id, name, description, visibility, pin_count, portrait, created_at, updated_at, deleted_at
								FROM boards
								WHERE name ILIKE '%' || $1 || '%' AND deleted_at IS NULL`
	QueryGetBoardById = `SELECT id, user_id, name, description, visibility, pin_count, portrait, created_at, updated_at, deleted_at
						 FROM boards
						 WHERE id = $1`
	QueryExistBoardById = `SELECT EXISTS(
								SELECT 1
								FROM boards
								WHERE id = $1 AND deleted_at IS NULL)`
	QueryCreateBoard = `INSERT INTO boards (id, user_id, name, description, visibility, pin_count, portrait, created_at, updated_at)
						   VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
						   RETURNING id, user_id, name, description, visibility, pin_count, portrait, created_at, updated_at, deleted_at`
	QueryUpdateBoard = `UPDATE boards
						SET name = $2, description = $3, visibility = $4, pin_count = $5, portrait = $6, updated_at = $7
						WHERE id = $1 AND deleted_at IS NULL`
	QueryDeleteBoard = `UPDATE boards
						SET deleted_at = $2
						WHERE id = $1 AND deleted_at IS NULL`
)

type boardRepository struct {
	DB *sql.DB
}

func NewBoardRepository(db *sql.DB) boards.BoardRepository {
	return &boardRepository{
		DB: db,
	}
}

func (r boardRepository) GetAll(ctx context.Context) ([]*boards.Board, error) {
	var (
		boardsList           []*boards.Board
		boardId, userId      uuid.UUID
		name                 string
		description          *string
		visibility           bool
		pinCount             int
		portrait             *string
		createdAt, updatedAt time.Time
		deletedAt            *time.Time
	)

	rows, err := r.DB.QueryContext(ctx, QueryGetAllBoards)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		err = rows.Scan(&boardId, &userId, &name, &description, &visibility, &pinCount, &portrait, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrScan, err)
		}

		board := boards.NewBoardFromDB(boardId, userId, name, description, visibility, pinCount, portrait, createdAt, updatedAt, deletedAt)
		boardsList = append(boardsList, board)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(got, ErrIterationRows, err)
	}

	return boardsList, nil
}

func (r boardRepository) GetList(ctx context.Context) ([]*boards.Board, error) {
	var (
		boardsList           []*boards.Board
		boardId, userId      uuid.UUID
		name                 string
		description          *string
		visibility           bool
		pinCount             int
		portrait             *string
		createdAt, updatedAt time.Time
		deletedAt            *time.Time
	)

	rows, err := r.DB.QueryContext(ctx, QueryGetListBoards)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		err = rows.Scan(&boardId, &userId, &name, &description, &visibility, &pinCount, &portrait, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrScan, err)
		}

		board := boards.NewBoardFromDB(boardId, userId, name, description, visibility, pinCount, portrait, createdAt, updatedAt, deletedAt)
		boardsList = append(boardsList, board)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(got, ErrIterationRows, err)
	}

	return boardsList, nil
}

func (r boardRepository) GetListByUserId(ctx context.Context, id uuid.UUID) ([]*boards.Board, error) {
	var (
		boardsList           []*boards.Board
		boardId              uuid.UUID
		name                 string
		description          *string
		visibility           bool
		pinCount             int
		portrait             *string
		createdAt, updatedAt time.Time
		deletedAt            *time.Time
	)

	rows, err := r.DB.QueryContext(ctx, QueryGetListBoardsByUserId, id)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		err = rows.Scan(&boardId, &name, &description, &visibility, &pinCount, &portrait, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrScan, err)
		}

		board := boards.NewBoardFromDB(boardId, id, name, description, visibility, pinCount, portrait, createdAt, updatedAt, deletedAt)
		boardsList = append(boardsList, board)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(got, ErrIterationRows, err)
	}

	return boardsList, nil
}

func (r boardRepository) GetListByName(ctx context.Context, name string) ([]*boards.Board, error) {
	var (
		boardsList           []*boards.Board
		boardId, userId      uuid.UUID
		description          *string
		visibility           bool
		pinCount             int
		portrait             *string
		createdAt, updatedAt time.Time
		deletedAt            *time.Time
	)

	rows, err := r.DB.QueryContext(ctx, QueryGetListBoardsByName, name)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		err = rows.Scan(&boardId, &userId, &description, &visibility, &pinCount, &portrait, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return nil, fmt.Errorf(got, ErrScan, err)
		}

		board := boards.NewBoardFromDB(boardId, userId, name, description, visibility, pinCount, portrait, createdAt, updatedAt, deletedAt)
		boardsList = append(boardsList, board)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(got, ErrIterationRows, err)
	}

	return boardsList, nil
}

func (r boardRepository) GetById(ctx context.Context, id uuid.UUID) (*boards.Board, error) {
	var (
		boardId, userId      uuid.UUID
		name                 string
		description          *string
		visibility           bool
		pinCount             int
		portrait             *string
		createdAt, updatedAt time.Time
		deletedAt            *time.Time
	)

	err := r.DB.QueryRowContext(ctx, QueryGetBoardById, id).Scan(
		&boardId, &userId, &name, &description, &visibility, &pinCount, &portrait, &createdAt, &updatedAt, &deletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	board := boards.NewBoardFromDB(boardId, userId, name, description, visibility, pinCount, portrait, createdAt, updatedAt, deletedAt)

	return board, nil
}

func (r boardRepository) ExistById(ctx context.Context, id uuid.UUID) (bool, error) {
	var exist bool

	err := r.DB.QueryRowContext(ctx, QueryExistBoardById, id).Scan(&exist)
	if err != nil {
		return false, fmt.Errorf(got, ErrQuery, err)
	}

	return exist, nil
}

func (r boardRepository) Create(ctx context.Context, b *boards.Board) (*boards.Board, error) {
	var (
		boardId, userId      uuid.UUID
		name                 string
		description          *string
		visibility           bool
		pinCount             int
		portrait             *string
		createdAt, updatedAt time.Time
		deletedAt            *time.Time
	)

	err := r.DB.QueryRowContext(ctx, QueryCreateBoard,
		b.Id(), b.UserId(), b.Name(), b.Description(), b.Visibility(), b.PinCount(), b.Portrait(), b.CreatedAt(), b.UpdatedAt(),
	).Scan(
		&boardId, &userId, &name, &description, &visibility, &pinCount, &portrait, &createdAt, &updatedAt, &deletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf(got, ErrQuery, err)
	}

	board := boards.NewBoardFromDB(boardId, userId, name, description, visibility, pinCount, portrait, createdAt, updatedAt, deletedAt)

	return board, nil
}

func (r boardRepository) Update(ctx context.Context, b *boards.Board) error {
	err := r.DB.QueryRowContext(ctx, QueryCreateBoard,
		b.Id(), b.UserId(), b.Name(), b.Description(), b.Visibility(), b.PinCount(), b.Portrait(), b.CreatedAt(), b.UpdatedAt(),
	)

	if err != nil {
		return fmt.Errorf(got, ErrQuery, err)
	}

	return nil
}

func (r boardRepository) Delete(ctx context.Context, b *boards.Board) error {
	err := r.DB.QueryRowContext(ctx, QueryDeleteBoard, b.Id(), b.DeletedAt())
	if err != nil {
		return fmt.Errorf(got, ErrQuery, err)
	}

	return nil
}
