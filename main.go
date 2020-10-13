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
		{1, 2},
		{1, -2},
		{-1, 2},
		{-1, -2},
		{2, 1},
		{2, -1},
		{-2, 1},
		{-2, -1},
	}

	return moves[index], nil
}

type BoardState struct {
	Board   [8][8]int
	Last    Move
	Count   int
	History []int
}

type BoardData struct {
	BD []BoardState
}

func NewBoardData() *BoardData {
	return &BoardData{}
}

func (bd *BoardData) Push(bs *BoardState) {
	tbs := NewBS()
	tbs.Board = bs.Board
	tbs.History = bs.History
	tbs.Count = bs.Count
	tbs.Last = bs.Last

	bd.BD = append(bd.BD, *tbs)
}

func NewBS() *BoardState {
	bs := &BoardState{}
	bs.Board = [8][8]int{}
	bs.Last = Move{0, 0}
	bs.Count = 1
	bs.Board[0][0] = bs.Count
	bs.History = append(bs.History, 0)
	return bs
}

func (db *BoardState) AddMove(index int) error {
	move, err := KnightOptions(index)
	if err != nil {
		return err
	}

	a := move.a + db.Last.a
	b := move.b + db.Last.b
	if a < 0 || a > 7 || b < 0 || b > 7 {
		return errors.New("off board")
	}

	if db.Board[a][b] != 0 {
		return errors.New("occupied square")
	}

	db.Count += 1
	db.Board[a][b] = db.Count
	db.Last = Move{a, b}
	db.History = append(db.History, index)
	return nil

}

func (db *BoardState) PrintBoard() {
	for i := 0; i < 7; i++ {
		fmt.Println(db.Board[i])
	}

}

func main() {

}
