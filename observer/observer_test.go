/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 18:53:27
 * @LastEditTime: 2023-04-18 18:54:37
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	sub := &Subject{}
	sub.Register(&Observer1{})
	sub.Register(&Observer2{})
	sub.Notify("hi")
}
