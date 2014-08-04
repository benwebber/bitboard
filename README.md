# board64

This library aims to implement a flexible system for storing state in games played on 8&times;8 grids. It does not implement game mechanics.

## Bitboards

Bitboards are bitmaps representing the positions of a certain type of piece (player/piece combination) on a game board. They are a compact way to store state in games like chess, checkers, and Reversi.

For example, we could use the following bitboard to represent the positions of White's pawns after White opens with **1. e4**.

```
8 | 0 0 0 0 0 0 0 0
7 | 0 0 0 0 0 0 0 0
6 | 0 0 0 0 0 0 0 0
5 | 0 0 0 0 0 0 0 0
4 | 0 0 0 0 1 0 0 0
3 | 0 0 0 0 0 0 0 0
2 | 1 1 1 1 0 1 1 1
1 | 0 0 0 0 0 0 0 0
    ---------------
    a b c d e f g h
```

Because an 8&times;8 grid has 64 squares, we can represent all the positions of a particular type using a single 64-bit integer. Of course, we can also represent smaller boards using the same technique.

For more information on bitboards, check out the [Chess Programming Wiki](https://chessprogramming.wikispaces.com/Bitboards).

## Example 

```go
package main

import "github.com/benwebber/board64"

func main() {
	b := board64.NewChessBoard()
	b.PrettyPrint()
}
```
