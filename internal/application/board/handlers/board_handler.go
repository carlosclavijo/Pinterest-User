package handlers

import boards "github.com/carlosclavijo/Pinterest-Services/internal/domain/board"

type BoardHandler struct {
	repository boards.BoardRepository
	factory    boards.BoardFactory
}

func NewBoardHandler(repository boards.BoardRepository, factory boards.BoardFactory) *BoardHandler {
	return &BoardHandler{
		repository: repository,
		factory:    factory,
	}
}
