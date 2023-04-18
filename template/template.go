/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 15:24:41
 * @LastEditTime: 2023-04-18 15:32:29
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package template

import (
	"fmt"
)

type ISMS interface {
	send(content, phone string) error
}

// SMS 短信发送基类
type sms struct {
	ISMS
}

// 校验短信字数是否超过最大值
func (s *sms) Valid(content string) error {
	if len(content) > 63 {
		return fmt.Errorf("content is too long")
	}
	return nil
}

// Send 发送短信
func (s *sms) Send(content string, phone string) error {
	if err := s.Valid(content); err != nil {
		return err
	}
	return s.send(content, phone)
}

// 子类发送短信走电信通道
type TeleComSms struct {
	*sms
}

func NewTeleComSms() *TeleComSms {
	tel := &TeleComSms{}
	tel.sms = &sms{ISMS: tel}
	return tel
}

func (tel *TeleComSms) send(context string, phone string) error {
	fmt.Println("send by telecom success")
	return nil
}
