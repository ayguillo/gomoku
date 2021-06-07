package main

import (
	"fmt"
	s "gomoku/structures"
)

func main() {
	ctx := s.SContext{}
	ctx.NSize = 19
	ctx.Goban = make([][]s.Tnumber, ctx.NSize)
	index := 0
	for index < int(ctx.NSize) {
		ctx.Goban[index] = make([]s.Tnumber, ctx.NSize)
		index++
	}
	fmt.Println(ctx)
}
