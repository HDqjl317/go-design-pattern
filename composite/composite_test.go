/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 14:05:49
 * @LastEditTime: 2023-04-18 14:05:57
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package composite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrganizetion(t *testing.T) {
	got := NewOrgnization().Count()
	assert.Equal(t, 20, got)
}
