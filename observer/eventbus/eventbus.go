/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 18:58:10
 * @LastEditTime: 2023-04-18 19:09:30
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package eventbus

import (
	"fmt"
	"reflect"
	"sync"
)

type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

// 异步事件总线
type AsyncEventBus struct {
	hanlders map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		hanlders: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

// Subscribe 订阅
func (bus *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	v := reflect.ValueOf(f)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	handler, ok := bus.hanlders[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	bus.hanlders[topic] = handler
	return nil
}

// Publish 发布
// 异步执行， 并且不会等待返回结果
func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.hanlders[topic]
	if !ok {
		fmt.Println("not found handlers in topic:", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		go handlers[i].Call(params)
	}
}
