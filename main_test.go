package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {

	// Test valid moves
	for i := 0; i <= 7; i++ {
		_, err := KnightOptions(i)
		if err != nil {
			t.Fatalf("0 through 7 inclusive are valid moves" +
				" should not throw this error")
		}
	}

	// Test invalid moves. Should throw error
	_, err := KnightOptions(-1)
	if err == nil {
		t.Fatalf("Error was not thrown, it was" +
			" and out-of-bounds entry")
	}
	_, err = KnightOptions(8)
	if err == nil {
		t.Fatalf("Error was not thrown, it was" +
			" and out-of-bounds entry")
	}

}

func Test_InitializeDB(t *testing.T) {
	bs := NewBS()
	bs.Board[1][1] = 1
}

func Test_AddMove(t *testing.T) {
	bs := NewBS()
	err := bs.AddMove(0)
	if err != nil {
		t.Fatalf("%s", err)
	}
	bs.AddMove(0)
	bs.AddMove(1)
	bs.PrintBoard()
	if !reflect.DeepEqual(bs.History, []int{0, 0, 0, 1}) {
		t.Fatalf("Should have been equal")
	}

}

func TestBoardData_Push(t *testing.T) {
	bd := NewBoardData()
	bs := NewBS()

	bs.AddMove(0)
	bd.Push(bs)
	bs.AddMove(0)
	bd.Push(bs)

	bd.BD[1].PrintBoard()
}

func TestOffBoard(t *testing.T) {
	bs := NewBS()

	bs.AddMove(0)
	bs.AddMove(0)
	bs.AddMove(0)
	err := bs.AddMove(0)
	if err == nil {
		t.Fatalf("We're not getting an error")
	}
	if err.Error() != "off board" {
		t.Fatalf("Should be off board")
	}

	bs.PrintBoard()
}

func TestOccupiedSquare(t *testing.T) {
	bs := NewBS()
	bs.AddMove(0)
	bs.AddMove(1)
	err := bs.AddMove(2)
	fmt.Println(err)
	if err == nil {
		t.Fatalf("We're not getting an error")
	}
	if err.Error() != "occupied square" {
		t.Fatalf("Should be occupied square")
	}

	bs.PrintBoard()
}
