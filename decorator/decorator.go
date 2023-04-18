/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 08:48:55
 * @LastEditTime: 2023-04-18 08:51:41
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package decorator

type IDraw interface {
	Draw() string
}

type Square struct {
}

func (s Square) Draw() string {
	return "this is a square"
}

type ColorSquare struct {
	square Square
	color  string
}

func (c ColorSquare) Draw() string {
	return c.square.Draw() + ", color is " + c.color
}

func NewColorSquare(square Square, color string) ColorSquare {
	return ColorSquare{color: color, square: square}
}
