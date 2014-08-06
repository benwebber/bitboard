// Package util implements utility functions for the bitboard library.
package util

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
// Coordinate conversions
//-----------------------------------------------------------------------------

// Convert coordinates in algebraic notiton to Cartesian coordinates.
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
