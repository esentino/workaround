package main

import (
	"fmt"

	"github.com/rivo/tview"
)

var app *tview.Application
var boardView *tview.Table
var chessboard ChessBoard

func main() {
	app = tview.NewApplication()
	boardView = tview.NewTable()

	chessboard = ChessBoard{
		color: white,
	}

	chessboard = fill(chessboard)
	var ok = allowed_move(1, 1, chessboard)
	fmt.Println(ok)
	grid := tview.NewGrid()
	grid.AddItem(boardView, 0, 0, 1, 1, 10, 0, false)
	command_input := tview.NewInputField()
	// command_input.SetDoneFunc(make_move).Autocomplete().SetAutocompleteFunc()
	grid.AddItem(command_input, 0, 1, 1, 1, 20, 0, true)
	boardView.SetFixed(8, 8)
	redraw_board(*boardView, chessboard)
	boardView.SetBorders(true)
	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func redraw_board(boardView tview.Table, chessboard ChessBoard) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			var figure_string string
			switch chessboard.board[j][i].sort {
			case pawn:
				if chessboard.board[j][i].color == white {
					figure_string = " ♙ "
				} else {
					figure_string = " ♟︎ "
				}
			case rook:
				if chessboard.board[j][i].color == white {
					figure_string = " ♖ "
				} else {
					figure_string = " ♜ "
				}
			case bishop:
				if chessboard.board[j][i].color == white {
					figure_string = " ♗ "
				} else {
					figure_string = " ♝ "
				}
			default:
				figure_string = "   "
			}
			boardView.SetCell(i+1, j+1, tview.NewTableCell(figure_string))
		}
	}
	for i := 1; i < 9; i++ {
		var c = rune('`' + i)
		boardView.SetCell(0, i, tview.NewTableCell(" "+string(c)+" "))
		boardView.SetCell(9, i, tview.NewTableCell(" "+string(c)+" "))

		var n = rune('0' + i)
		boardView.SetCell(i, 0, tview.NewTableCell(" "+string(n)+" "))
		boardView.SetCell(i, 9, tview.NewTableCell(" "+string(n)+" "))
	}
}

func allowed_move(x int, y int, chessboard ChessBoard) []Location {
	fmt.Println(chessboard.board[x][y].sort)
	figure := chessboard.board[x][y]
	switch figure.sort {
	case pawn:
		if figure.color == white {
			return []Location{
				{x, y + 1}, {x, y + 2},
			}
		} else {
			return []Location{
				{x, y - 1}, {x, y - 2},
			}
		}
	default:
		return []Location{}
	}

}

func fill(chessboard ChessBoard) ChessBoard {
	chessboard = fillEmpty(chessboard)
	chessboard = fillPawns(chessboard, white)
	chessboard = fillPawns(chessboard, black)
	chessboard = fillRooks(chessboard, white)
	chessboard = fillRooks(chessboard, black)
	chessboard = fillBishops(chessboard, white)
	chessboard = fillBishops(chessboard, black)
	return chessboard
}

func fillBishops(chessboard ChessBoard, color Color) ChessBoard {
	var row int
	switch color {
	case white:
		row = 0
	default:
		row = 7
	}
	chessboard.board[2][row] = Figure{
		sort:       bishop,
		first_move: true,
		color:      color,
	}
	chessboard.board[5][row] = Figure{
		sort:       bishop,
		first_move: true,
		color:      color,
	}
	return chessboard
}

func fillRooks(chessboard ChessBoard, color Color) ChessBoard {
	var row int
	switch color {
	case white:
		row = 0
	default:
		row = 7
	}
	chessboard.board[0][row] = Figure{
		sort:       rook,
		first_move: true,
		color:      color,
	}
	chessboard.board[7][row] = Figure{
		sort:       rook,
		first_move: true,
		color:      color,
	}
	return chessboard
}

func fillPawns(chessboard ChessBoard, color Color) ChessBoard {
	var row int
	switch color {
	case white:
		row = 1
	default:
		row = 6
	}
	for column := 0; column < 8; column++ {
		chessboard.board[column][row] = Figure{
			sort:       pawn,
			first_move: true,
			color:      color,
		}
	}
	return chessboard
}

func fillEmpty(chessboard ChessBoard) ChessBoard {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			chessboard.board[i][j] = Figure{empty_field, true, empty}
		}
	}
	return chessboard
}
