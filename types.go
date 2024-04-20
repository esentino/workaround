package main

type Color int
type FigureSort int
type Location struct {
	x int
	y int
}

const (
	empty_field FigureSort = iota
	pawn
	rook
	bishop
)
const (
	white Color = iota
	black
	empty
)

type ChessBoard struct {
	color Color
	board [8][8]Figure
}

type Figure struct {
	sort       FigureSort
	first_move bool
	color      Color
}
