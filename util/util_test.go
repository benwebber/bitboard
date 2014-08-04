package util

import "testing"

var positionsAlgebraic = []string{
	"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1",
	"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2",
	"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3",
	"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4",
	"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5",
	"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6",
	"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7",
	"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8",
}

var positionsBit = []int{
	0, 1, 2, 3, 4, 5, 6, 7,
	8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 54, 55,
	56, 57, 58, 59, 60, 61, 62, 63,
}

func TestAlgebraicToBit(t *testing.T) {
	for i, p := range positionsAlgebraic {
		result := AlgebraicToBit(p, 8)
		if result != positionsBit[i] {
			t.Error("Expected", positionsBit[i], ", got", result)
		}
	}
}

func TestAlgebraicToCartesian(t *testing.T) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			bit := y*8 + x
			p := BitToAlgebraic(bit, 8)
			i, j := AlgebraicToCartesian(p, 8)
			if (i != x) || (j != y) {
				t.Error("Expected x:", x, "y:", y, ", got x:", i, "y:", j)
			}
		}
	}
}

func TestBitToAlgebraic(t *testing.T) {
	for i, p := range positionsBit {
		result := BitToAlgebraic(p, 8)
		if result != positionsAlgebraic[i] {
			t.Error("Expected", positionsAlgebraic[i], ", got", result)
		}
	}
}

func TestBitToCartesian(t *testing.T) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			bit := y*8 + x
			i, j := BitToCartesian(bit, 8)
			if (i != x) || (j != y) {
				t.Error("Expected x:", x, "y:", y, ", got x:", i, "y:", j)
			}
		}
	}
}

func TestCartesianToAlgebraic(t *testing.T) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			bit := y*8 + x
			p := CartesianToAlgebraic(x, y, 8)
			if p != positionsAlgebraic[bit] {
				t.Error("Expected", positionsBit[bit], "got", p)
			}
		}
	}
}

func TestCartesianToBit(t *testing.T) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			bit := CartesianToBit(x, y, 8)
			if bit != positionsBit[bit] {
				t.Error("Expected", positionsBit[bit], "got", bit)
			}
		}
	}
}
