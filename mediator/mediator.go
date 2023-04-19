/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-19 09:53:41
 * @LastEditTime: 2023-04-19 10:05:41
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package mediator

import (
	"fmt"
	"reflect"
)

// Input 表示一个输入框
type Input string

func (i Input) String() string {
	return string(i)
}

// Selection 表示一个选择框
type Selection string

// Selected 当前选中的对象
func (s Selection) Selected() string {
	return string(s)
}

type Button struct {
	onClick func()
}

// SetOnClick 添加点击事件回调
func (b *Button) SetOnClick(f func()) {
	b.onClick = f
}

// 中介模式接口
type IMediator interface {
	HandleEvent(component interface{})
}

// Dialog 对话框组件
type Dialog struct {
	LoginButton         *Button
	RegButton           *Button
	Selection           *Selection
	UsernameInput       *Input
	PasswordInput       *Input
	RepeatPasswordInput *Input
}

func (d *Dialog) HandleEvent(component interface{}) {
	switch {
	case reflect.DeepEqual(component, d.Selection):
		if d.Selection.Selected() == "登录" {
			fmt.Println("select login")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
		} else if d.Selection.Selected() == "注册" {
			fmt.Println("select register")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
			fmt.Printf("show: %s\n", d.RepeatPasswordInput)
		}
		// others code...
	}
}
