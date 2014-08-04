// Package board64 implements 8x8 bitboards for games like chess, checkers,
// Reversi, and Othello.
package board64

import "fmt"

// A Bitboard represents game state.
//
// We use a little-endian mapping of bits to rank-and-file coordinates. For
// an 8x8 board, this mapping looks like:
//
//   8 | 56 57 58 59 60 61 62 63
//   7 | 48 49 50 51 52 53 54 55
//   6 | 40 41 42 43 44 45 46 47
//   5 | 32 33 34 35 36 37 38 39
//   4 | 24 25 26 27 28 29 30 31
//   3 | 16 17 18 19 20 21 22 23
//   2 | 8  9  10 11 12 13 14 15
//   1 | 0  1  2  3  4  5  6  7
//     -------------------------
//       a  b  c  d  e  f  g  h
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
		if GetBit(&b.Bitmaps[i], p) != 0 {
			return i
		}
	}
	return -1 // not found
}

// CoordsToBit converts rank and file coordinates to an integer bit position.
func (b *Bitboard) CoordsToBit(file string, rank int) int {
	files := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var f int
	for i, v := range files {
		if file == v {
			f = i
		}
	}
	bit := (rank-1)*b.Files + f
	return bit
}

// BitToSquareIndex converts an integer bit position to a square index (e.g.,
// e4).
func (b *Bitboard) BitToSquareIndex(p int) string {
	files := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	r := p/b.Files + 1
	f := p % b.Files
	return fmt.Sprintf("%v%v", files[f], r)
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

// SetBit sets (sets to 1) the bit at position p.
func SetBit(i *uint64, p int) {
	var mask uint64
	mask = (1 << uint(p))
	*i |= mask
}

// ClearBit clears (sets to 0) the bit at position p.
func ClearBit(i *uint64, p int) {
	var mask uint64
	mask = ^(1 << uint(p))
	*i &= mask
}

// ToggleBit toggles the value of the bit at position p.
func ToggleBit(i *uint64, p int) {
	var mask uint64
	mask = (1 << uint(p))
	*i ^= mask
}

// GetBit returns the value of the bit at position p.
func GetBit(i *uint64, p int) int {
	return int((*i >> uint(p)) & 1)
}

func IsBitSet(i uint64, p int) bool {
	var mask uint64
	mask = 1 << uint64(p)
	return (i & mask) != 0
}

// PopCount calculates the population count (Hamming weight) of an integer
// using a divide-and-conquer approach.
//
// See <http://en.wikipedia.org/wiki/Hamming_weight> for a complete description
// of this implementation.
func PopCount(i uint64) int {
	var mask1, mask2, mask4 uint64
	mask1 = 0x5555555555555555 // 0101...
	mask2 = 0x3333333333333333 // 00110011..
	mask4 = 0x0f0f0f0f0f0f0f0f // 00001111...
	i -= (i >> 1) & mask1
	i = (i & mask2) + ((i >> 2) & mask2)
	i = (i + (i >> 4)) & mask4
	i += i >> 8
	i += i >> 16
	i += i >> 32
	return int(i & 0x7f)
}
