/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 22:07:01
 * @LastEditTime: 2023-04-19 08:36:24
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package state

import "fmt"

// IState 状态
type IState interface {
	Approval(m *Machine)
	Reject(m *Machine)
	GetName() string // 获取当前状态名称
}

type Machine struct {
	state IState
}

func (m *Machine) SetState(state IState) {
	m.state = state
}

func (m *Machine) GetStateName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

// 领导审批
type leaderApproveState struct{}

func (leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader approval")
	m.SetState(GetFinanceApproveState())
}

func (leaderApproveState) GetName() string {
	return "LeaderApproveState"
}

func (leaderApproveState) Reject(m *Machine) {}

func GetLeaderApproveState() IState {
	return &leaderApproveState{}
}

// 财务审批
type financeApproveState struct{}

func (f financeApproveState) Approval(m *Machine) {
	fmt.Println("finance approval")
	fmt.Println("ok, done")
}

func (f financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

func (f financeApproveState) GetName() string {
	return "FinanceApproveState"
}

func GetFinanceApproveState() IState {
	return &financeApproveState{}
}
