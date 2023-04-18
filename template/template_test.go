/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 15:32:49
 * @LastEditTime: 2023-04-18 15:32:49
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSms_Send(t *testing.T) {
	tel := NewTeleComSms()
	err := tel.Send("hello sms!", "123456789")
	assert.NoError(t, err)
}
