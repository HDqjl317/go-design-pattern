/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 19:10:26
 * @LastEditTime: 2023-04-18 19:12:34
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package eventbus

import (
	"fmt"
	"testing"
	"time"
)

func sub1(msg1, msg2 string) {
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("sub1, %s %s\n", msg1, msg2)
}

func sub2(msg1, msg2 string) {
	fmt.Printf("sub2, %s %s\n", msg1, msg2)
}

func TestAsyncEventBus_Publish(t *testing.T) {
	bus := NewAsyncEventBus()
	bus.Subscribe("topic:1", sub1)
	bus.Subscribe("topic:1", sub2)
	bus.Publish("topic:1", "test1", "test2")
	bus.Publish("topic:1", "testA", "testB")
	time.Sleep(1 * time.Second)
}
