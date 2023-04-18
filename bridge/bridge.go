/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 08:33:25
 * @LastEditTime: 2023-04-18 08:39:36
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package bridge

type IMsgSender interface {
	Send(msg string) error
}

// EmailMsgSender 发送电子邮件
type EmailMsgSender struct {
	emails []string
}

func (es *EmailMsgSender) Send(msg string) error {
	// some code
	return nil
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

type INotification interface {
	Notify(msg string) error
}

// ErrorNotification 错误通知接口
type ErrorNotification struct {
	sender IMsgSender
}

func (errnotiy ErrorNotification) Notify(msg string) error {
	return errnotiy.sender.Send(msg)
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}
