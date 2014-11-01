// Utility functions for the bitboard library.
package bitboard

import (
	"fmt"
	"strconv"
)

//-----------------------------------------------------------------------------
// Bit operations
//-----------------------------------------------------------------------------

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

// Calculate the union of integers.
func Union(i ...uint64) uint64 {
	var u uint64
	for _, v := range i {
		u = u | v
	}
	return u
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

//-----------------------------------------------------------------------------
// Flipping and rotating
//-----------------------------------------------------------------------------

// These are Go ports of the functions given on the Chess Programming wiki:
// <https://chessprogramming.wikispaces.com/Flipping+Mirroring+and+Rotating>.

// Flip a bitboard vertically about the centre ranks.
func FlipVertical(i uint64) uint64 {
	k1 := uint64(0x00FF00FF00FF00FF)
	k2 := uint64(0x0000FFFF0000FFFF)
	i = ((i >> 8) & k1) | ((i & k1) << 8)
	i = ((i >> 16) & k2) | ((i & k2) << 16)
	i = (i >> 32) | (i << 32)
	return i
}

// Flip a bitboard horizontally about the centre files.
func FlipHorizontal(i uint64) uint64 {
	k1 := uint64(0x5555555555555555)
	k2 := uint64(0x3333333333333333)
	k4 := uint64(0x0f0f0f0f0f0f0f0f)
	i = ((i >> 1) & k1) + 2*(i&k1)
	i = ((i >> 2) & k2) + 4*(i&k2)
	i = ((i >> 4) & k4) + 16*(i&k4)
	return i
}

// Flip a bitboard about the diagonal A1-H8.
func FlipDiagonalA1H8(i uint64) uint64 {
	var t uint64
	k1 := uint64(0x5500550055005500)
	k2 := uint64(0x3333000033330000)
	k4 := uint64(0x0f0f0f0f00000000)
	t = k4 & (i ^ (i << 28))
	i ^= t ^ (t >> 28)
	t = k2 & (i ^ (i << 14))
	i ^= t ^ (t >> 14)
	t = k1 & (i ^ (i << 7))
	i ^= t ^ (t >> 7)
	return i
}

func FlipDiagonalA8H1(i uint64) uint64 {
	var t uint64
	k1 := uint64(0xaa00aa00aa00aa00)
	k2 := uint64(0xcccc0000cccc0000)
	k4 := uint64(0xf0f0f0f00f0f0f0f)
	t = i ^ (i << 36)
	i ^= k4 & (t ^ (i >> 36))
	t = k2 & (i ^ (i << 18))
	i ^= t ^ (t >> 18)
	t = k1 & (i ^ (i << 9))
	i ^= t ^ (t >> 9)
	return i
}

// Rotate a bitboard by 180 degrees.
func Rotate180(i uint64) uint64 {
	return FlipHorizontal(FlipVertical(i))
}

// Rotate a bitboard by 90 degrees (clockwise).
func Rotate90(i uint64) uint64 {
	return FlipVertical(FlipDiagonalA1H8(i))
}

// Rotate a bitboard by 270 degrees (90 degrees counter-clockwise).
func Rotate270(i uint64) uint64 {
	return FlipDiagonalA1H8(FlipVertical(i))
}

//-----------------------------------------------------------------------------
// Coordinate conversions
//-----------------------------------------------------------------------------

// Convert coordinates in algebraic notation to Cartesian coordinates.
func AlgebraicToCartesian(p string, files int) (int, int) {
	symbols := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var x int
	for i, v := range symbols {
		if string(p[0]) == v {
			x = i
		}
	}
	y, _ := strconv.Atoi(string(p[1]))
	return x, (y - 1)
}

// Convert coordinates in algebraic notation to an integer bit position.
func AlgebraicToBit(p string, files int) int {
	x, y := AlgebraicToCartesian(p, files)
	return CartesianToBit(x, y, files)
}

// Convert an integer bit position to coordiantes in algebraic notation.
func BitToAlgebraic(p int, files int) string {
	x, y := BitToCartesian(p, files)
	return CartesianToAlgebraic(x, y, files)
}

// Convert an integer bit position to Cartesian coordinates.
func BitToCartesian(p int, files int) (int, int) {
	x := p % files
	y := p / files
	return x, y
}

// Convert Cartesian coordinates to coordinates in algebraic notation.
func CartesianToAlgebraic(x int, y int, files int) string {
	symbols := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return fmt.Sprintf("%v%v", symbols[x], y+1)
}

// Convert Cartesian coordinates to an integer bit position.
func CartesianToBit(x int, y int, files int) int {
	bit := y*files + x
	return bit
}
