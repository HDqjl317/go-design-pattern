/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-11 20:42:56
 * @LastEditTime: 2023-04-11 20:49:30
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package singleton

import "sync"

// 饿汉模式
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}

// 懒汉模式
var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
