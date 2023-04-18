/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 14:20:06
 * @LastEditTime: 2023-04-18 14:41:37
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package flyweight

// 棋子享元
type ChesePieceUnit struct {
	ID    uint
	Name  string
	Color string
}

var uints = map[int]*ChesePieceUnit{
	1: {
		ID:    1,
		Name:  "车",
		Color: "rad",
	},
	2: {
		ID:    2,
		Name:  "炮",
		Color: "red",
	},
	// others ...
}

// 工厂
func NewChesePieceUint(id int) *ChesePieceUnit {
	return uints[id]
}

// 棋子
type ChesePiece struct {
	Uint *ChesePieceUnit
	X    int
	Y    int
}

type CheseBoard struct {
	chesePiece map[int]*ChesePiece
}

// 初始化棋盘
func NewChessBoard() *CheseBoard {
	board := &CheseBoard{chesePiece: map[int]*ChesePiece{}}
	for id := range uints {
		board.chesePiece[id] = &ChesePiece{
			Uint: NewChesePieceUint(id),
			X:    0,
			Y:    0,
		}
	}
	return board
}

func (c *CheseBoard) Move(id, x, y int) {
	c.chesePiece[id].X = x
	c.chesePiece[id].Y = y
}
