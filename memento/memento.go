/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-19 09:24:56
 * @LastEditTime: 2023-04-19 09:25:03
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package memento

// 用于保存数据
type InputText struct {
	content string
}

func (in *InputText) Append(content string) {
	in.content += content
}

func (in *InputText) GetText() string {
	return in.content
}

type Snapshot struct {
	content string
}

// Snapshot 创建快照
func (in *InputText) Snapshot() *Snapshot {
	return &Snapshot{content: in.content}
}

// 从快照中恢复
func (in *InputText) Restore(s *Snapshot) {
	in.content = s.GetText()
}

func (s *Snapshot) GetText() string {
	return s.content
}
