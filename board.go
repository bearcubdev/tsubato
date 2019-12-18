package tsubato

import (
	"encoding/json"
)

type boardList struct {
	Boards []Board
}

var localBoards map[string]Board

func getBoards() (boardList, error) {
	JSON, err := get("/boards.json")
	if err != nil {
		return boardList{}, err
	}

	tempBoards := boardList{}
	err = json.Unmarshal(JSON, &tempBoards)
	if err != nil {
		return boardList{}, err
	}

	return tempBoards, nil
}

func loadBoards() error {
	localBoards = make(map[string]Board)
	tempBoards, err := getBoards()
	if err != nil {
		return err
	}
	for _, board := range tempBoards.Boards {
		localBoards[board.Board] = board
	}
	return nil
}

// GetBoard returns the information for a board
func GetBoard(name string) (Board, error) {
	if len(localBoards) == 0 {
		err := loadBoards()
		if err != nil {
			return Board{}, err
		}
	}
	return localBoards[name], nil
}

// GetBoards returns the information for all boards
func GetBoards() ([]Board, error) {
	if len(localBoards) == 0 {
		err := loadBoards()
		if err != nil {
			return []Board{}, err
		}
	}
	boardList := []Board{}
	for _, board := range localBoards {
		boardList = append(boardList, board)
	}
	return boardList, nil
}
