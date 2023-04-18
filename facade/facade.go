/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 13:26:18
 * @LastEditTime: 2023-04-18 13:35:53
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package facade

// IUser 用户接口
type IUser interface {
	Login(phone, code int) (*User, error)
	Register(phone, code int) (*User, error)
}

type User struct {
	Name string
}

func newUser(name string) IUser {
	return User{Name: name}
}

func (u User) Login(phone int, code int) (*User, error) {
	// 校验操作 ...
	return &User{Name: "test login"}, nil
}

func (u User) Register(phone int, code int) (*User, error) {
	// 校验操作 ...
	// 创建用户
	return &User{Name: "test register"}, nil
}

type UserService struct {
	User IUser
}

func NewUserService(name string) UserService {
	return UserService{
		User: newUser(name),
	}
}

// Login or Register
func (u UserService) LoginorRegister(phone, code int) (*User, error) {
	user, err := u.User.Login(phone, code)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	return u.User.Register(phone, code)

}
