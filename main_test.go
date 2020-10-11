package main

import (
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
	db := NewDB()
	db.Board[1][1] = 1
}


func TestDB_AddMove(t *testing.T) {
	db := NewDB()
	db.AddMove(0)
	db.PrintBoard()
}