/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 08:39:50
 * @LastEditTime: 2023-04-18 08:44:27
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotfication_Notify(t *testing.T) {
	sender := NewEmailMsgSender([]string{"test@test.com"})
	n := NewErrorNotification(sender)
	err := n.Notify("test msg")

	assert.Nil(t, err)
}
