/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 08:53:32
 * @LastEditTime: 2023-04-18 08:55:11
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
	assert.Equal(t, "this is a square, color is red", got)
}
