/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-19 08:36:31
 * @LastEditTime: 2023-04-19 08:39:52
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachine_GetStateName(t *testing.T) {
	m := &Machine{state: GetLeaderApproveState()}
	assert.Equal(t, "LeaderApproveState", m.GetStateName())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetStateName())
	m.Reject()
	assert.Equal(t, "LeaderApproveState", m.GetStateName())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetStateName())
	m.Approval()
}
