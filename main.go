package main

import (
	"errors"
	"fmt"
)

type Move struct {
	a int
	b int
}

func KnightOptions(index int) (Move, error) {
	if index < 0 || index > 7 {
		return Move{}, errors.New("index out of bounds")
	}

	moves := []Move{
		Move{1, 2},
		Move{1, -2},
		Move{-1, 2},
		Move{-1, -2},
		Move{2, 1},
		Move{2, -1},
		Move{-2, 1},
		Move{-2, -1},
	}

	return moves[index], nil
}

type DB struct {
	Board [8][8]int
	Last  Move
	Count int
}

func NewDB() *DB {
	db := &DB{}
	db.Board = [8][8]int{}
	return &DB{}
}

func (db *DB) AddMove(index int) {
	move, err := KnightOptions(index)
	if err != nil {
		return
	}
	db.Count += 1
	db.Board[move.a][move.b] = db.Count
	db.Last = move

}

func (db *DB) PrintBoard() {
	for i := 0; i < 7; i++ {
		fmt.Println(db.Board[i])
	}

}

func main() {

}
