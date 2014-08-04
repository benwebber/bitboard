// Package board64 implements 8x8 bitboards for games like chess, checkers,
// Reversi, and Othello.
package board64

import (
	"fmt"

	"github.com/benwebber/board64/util"
)

// A Bitboard represents game state.
//
// We use a little-endian mapping of bits to rank-and-file coordinates. For
// an 8x8 board, this mapping looks like:
//
//  8 | 56 57 58 59 60 61 62 63
//  7 | 48 49 50 51 52 53 54 55
//  6 | 40 41 42 43 44 45 46 47
//  5 | 32 33 34 35 36 37 38 39
//  4 | 24 25 26 27 28 29 30 31
//  3 | 16 17 18 19 20 21 22 23
//  2 | 8  9  10 11 12 13 14 15
//  1 | 0  1  2  3  4  5  6  7
//    -------------------------
//      a  b  c  d  e  f  g  h
//
// Coordinates
//
// We will define three sets of coordinates for interacting with the Bitboard:
//
//  * bit positions
//  * alegraic notation
//  * Cartesian (x, y) coordinates
//
// For example, the following positions are equivalent on an 8x8 board:
//
//  | Bit | Algebraic | Cartesian |
//  |-----|-----------|-----------|
//  | 0   | a1        | (0, 0)    |
//  | 22  | g3        | (6, 2)    |
//  | 28  | e4        | (4, 4)    |
//  | 35  | d5        | (3, 4)    |
//
// Construct a new Bitboard using NewBitboard. There are also convenience
// functions for constructing bitboards for specific games.
type Bitboard struct {
	Bitmaps []uint64 // Bitmaps for each colour/piece combination
	Symbols []string // Symbols representing each colour/piece combination
	Ranks   int      // Number of rows
	Files   int      // Number of columns
}

// PrettyPrint pretty-prints a Bitboard using the symbols for each colour/piece
// combination. Empty squares are represented by periods.
func (b *Bitboard) PrettyPrint() {
	for r := b.Ranks; r > 0; r-- {
		for f := 0; f < b.Files; f++ {
			p := (r-1)*b.Files + f
			i := b.GetBitmapIndex(p)
			if i != -1 {
				fmt.Print(b.Symbols[i])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// GetBitmapIndex returns the array index of the bitmap including a particular
// square.
func (b *Bitboard) GetBitmapIndex(p int) int {
	for i := 0; i < len(b.Bitmaps); i++ {
		if util.GetBit(&b.Bitmaps[i], p) != 0 {
			return i
		}
	}
	return -1 // not found
}

// Convert coordinates in algebraic notation to an integer bit position.
// Wrap util.AlgebraicToBit to automatically pass in number of files.
func (b *Bitboard) AlgebraicToBit(p string) int {
	return util.AlgebraicToBit(p, b.Files)
}

// Convert coordinates in algebraic notiton to Cartesian coordinates.
// Wrap util.AlgebraicToCartesian to automatically pass in number of files.
func (b *Bitboard) AlgebraicToCartesian(p string) (int, int) {
	return util.AlgebraicToCartesian(p, b.Files)
}

// Convert an integer bit position to coordiantes in algebraic notation.
// Wrap util.BitToAlgebraic to automatically pass in number of files.
func (b *Bitboard) BitToAlgebraic(p int) string {
	return util.BitToAlgebraic(p, b.Files)
}

// Convert an integer bit position to Cartesian coordinates.
// Wrap util.BitToCartesian to automatically pass in number of files.
func (b *Bitboard) BitToCartesian(p int) (int, int) {
	return util.BitToCartesian(p, b.Files)
}

// Convert Cartesian coordinates to coordinates in algebraic notation.
// Wrap util.CartesianToAlgebraic to automatically pass in number of files.
func (b *Bitboard) CartesianToAlgebraic(x int, y int) string {
	return util.CartesianToAlgebraic(x, y, b.Files)
}

// Convert Cartesian coordinates to an integer bit position.
// Wrap util.CartesianToBit to automatically pass in number of files.
func (b *Bitboard) CartesianToBit(x int, y int) int {
	return util.CartesianToBit(x, y, b.Files)
}

// Move a piece from algebraic position p1 to p2.
func (b *Bitboard) MovePieceAlgebraic(m int, p1 string, p2 string) {
}

// Move a piece from bit position p1 to p2.
func (b *Bitboard) MovePieceBit(m int, p1 int, p2 int) {
	util.ClearBit(&b.Bitmaps[m], p1)
	util.SetBit(&b.Bitmaps[m], p2)
}

// Move a piece using Cartesian coordinates.
func (b *Bitboard) MovePieceCartesian(m int, x1 int, y1 int, x2 int, y2 int) {
}

// NewBitboard constructs a new Bitboard.
func NewBitboard() *Bitboard {
	return &Bitboard{}
}

// NewChessBoard is a convenience function for constructing a new chess board.
func NewChessBoard() *Bitboard {
	bitmaps := []uint64{
		uint64(0x0000000000000081), // White Rooks
		uint64(0x0000000000000042), // White Knights
		uint64(0x0000000000000024), // White Bishops
		uint64(0x0000000000000008), // White Queen
		uint64(0x0000000000000010), // White King
		uint64(0x000000000000ff00), // White Pawns
		uint64(0x8100000000000000), // Black Rooks
		uint64(0x4200000000000000), // Black Knights
		uint64(0x2400000000000000), // Black Bishops
		uint64(0x0800000000000000), // Black Queen
		uint64(0x1000000000000000), // Black King
		uint64(0x00ff000000000000), // Black Pawns
	}
	symbols := []string{
		"R", "N", "B", "Q", "K", "P",
		"r", "n", "b", "q", "k", "p",
	}
	return &Bitboard{bitmaps, symbols, 8, 8}
}

// NewCheckersBoard is a convenience function for constructing a new checkers
// (English draughts) board.
func NewCheckersBoard() *Bitboard {
	bitmaps := []uint64{
		uint64(0xaa55aa0000000000), // Red
		uint64(0x000000000055aa55), // White
	}
	symbols := []string{"R", "W"}
	return &Bitboard{bitmaps, symbols, 8, 8}
}

// NewOthelloBoard is a convenience function for constructing a new Othello
// board.
//
// Othello differs from Reversi only in starting position.
func NewOthelloBoard() *Bitboard {
	bitmaps := []uint64{
		uint64(0x0000001008000000), // Black
		uint64(0x0000000810000000), // White
	}
	symbols := []string{"B", "W"}
	return &Bitboard{bitmaps, symbols, 8, 8}
}

// NewReversiBoard is a convenience function for constructing a new Reversi
// board.
func NewReversiBoard() *Bitboard {
	bitmaps := []uint64{
		uint64(0x0000000000000000), // Black
		uint64(0x0000000000000000), // White
	}
	symbols := []string{"B", "W"}
	return &Bitboard{bitmaps, symbols, 8, 8}
}

// NewTicTacToeBoard is a convenience function for constructing a new
// Tic-Tac-Toe board.
func NewTicTacToeBoard() *Bitboard {
	bitmaps := []uint64{
		uint64(0x0000000000000000), // X
		uint64(0x0000000000000000), // O
	}
	symbols := []string{"X", "O"}
	return &Bitboard{bitmaps, symbols, 3, 3}
}

// NewConnectFourBoard is a convenience function for constructing a new Connect
// Four board.
func NewConnectFourBoard() *Bitboard {
	bitmaps := []uint64{
		uint64(0x0000000000000000), // Red
		uint64(0x0000000000000000), // Yellow
	}
	symbols := []string{"R", "Y"}
	return &Bitboard{bitmaps, symbols, 6, 7}
}
