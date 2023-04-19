/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-19 10:18:39
 * @LastEditTime: 2023-04-19 10:24:52
 * @Description:假设现在有一个游戏服务，我们正在实现一个游戏后端
				使用一个 goroutine 不断接收来自客户端请求的命令，并且将它放置到一个队列当中
				然后我们在另外一个 goroutine 中来执行它
 * Copyright © jiale_quan, All Rights Reserved
*/
package command

import "fmt"

// Command 命令
type Command func() error

// StartCommandFunc 返回一个 Command 命令
// 是因为正常情况下不会是这么简单的函数, 一般都会有一些参数
func StartCommandFunc() Command {
	return func() error {
		fmt.Println("game start")
		return nil
	}
}

func ArchiveCommandFunc() Command {
	return func() error {
		fmt.Println("game archive")
		return nil
	}
}
