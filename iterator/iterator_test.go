/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-19 08:56:46
 * @LastEditTime: 2023-04-19 08:58:37
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayInt_Iterator(t *testing.T) {
	data := ArrayInt{1, 3, 5, 7, 8}
	iterator := data.Iterator()
	i := 0
	for iterator.HasNext() {
		assert.Equal(t, data[i], iterator.CurrentItem())
		iterator.Next()
		i++
	}
}
