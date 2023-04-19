/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-19 10:05:50
 * @LastEditTime: 2023-04-19 10:12:37
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package mediator

import (
	"testing"
)

func TestDemo(t *testing.T) {
	usernameInput := Input("username input")
	passwordInput := Input("password input")
	repeatPwdInput := Input("repeat password input")

	selection := Selection("登录")
	d := &Dialog{
		Selection:           &selection,
		UsernameInput:       &usernameInput,
		PasswordInput:       &passwordInput,
		RepeatPasswordInput: &repeatPwdInput,
	}
	d.HandleEvent(&selection)

	regSelection := Selection("注册")
	d.Selection = &regSelection
	d.HandleEvent(&regSelection)
}
