/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 14:41:54
 * @LastEditTime: 2023-04-18 14:45:28
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package flyweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCheseBoard(t *testing.T) {
	board1 := NewChessBoard()
	board2 := NewChessBoard()

	board1.Move(1, 1, 2)
	board2.Move(2, 2, 3)

	assert.Equal(t, board1.chesePiece[1].Uint, board2.chesePiece[1].Uint)
	assert.Equal(t, board1.chesePiece[2].Uint, board2.chesePiece[2].Uint)
}
