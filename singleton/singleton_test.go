/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-11 20:51:10
 * @LastEditTime: 2023-04-11 21:19:30
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package singleton_test

import (
	"fmt"
	"go-design-pattern/singleton"
	"testing"
)

func _assert(condition bool, msg string) {
	if !condition {
		panic(fmt.Sprintf("assertion failed:" + msg))
	}
}

func TestGetInstance(t *testing.T) {
	_assert(singleton.GetInstance() == singleton.GetInstance(), "Two singleton instances.")
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetInstance() != singleton.GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestGetLazyInstance(t *testing.T) {
	_assert(singleton.GetLazyInstance() == singleton.GetLazyInstance(), "Two singleton instances.")
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetLazyInstance() != singleton.GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
