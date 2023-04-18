/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 13:32:53
 * @LastEditTime: 2023-04-18 13:38:08
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Login(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.User.Login(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}

func TestUserService_LoginOrRegister(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.LoginorRegister(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}

func TestUserService_Register(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.User.Register(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test register"}, user)
}
