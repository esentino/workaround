package main

import "testing"

func TestFillChessboard(t *testing.T) {
	chessboard := ChessBoard{color: white}
	chessboard = fill(chessboard)
	if chessboard.board[1][1].color != white {
		panic("should be white figure in field 1, 1")
	}
	if chessboard.board[1][1].first_move != true {
		panic("clean figure on board should have first move")
	}
}
