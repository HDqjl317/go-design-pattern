/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 18:23:28
 * @LastEditTime: 2023-04-18 18:27:31
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package proxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserProxy_Login(t *testing.T) {
	proxy := NewUserProxy(&User{})

	err := proxy.Login("test", "password")

	require.Nil(t, err)
}
