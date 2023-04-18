/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 09:18:54
 * @LastEditTime: 2023-04-18 09:30:34
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package adapter

import (
	"fmt"
)

// ICreateServer 创建云主机
type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

// AWSClient aws sdk
type AWSClient struct{}

// RunInstance 启动实例
func (c *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run success, cpu: %f, mem: %f", cpu, mem)
	return nil
}

// AwsClientAdpter 适配器
type AwsClientAdpter struct {
	Client AWSClient
}

func (a *AwsClientAdpter) CreateServer(cpu, mem float64) error {
	a.Client.RunInstance(cpu, mem)
	return nil
}

// AliyunClient aliyun sdk
type AliyunClient struct{}

// CreateServer 启动实例
func (c *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aws client run success, cpu: %d, mem: %d", cpu, mem)
	return nil
}

// AliyunClientAdapter 适配器
type AliyunClientAdapter struct {
	Client AliyunClient
}

// CreateServer 启动实例
func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.CreateServer(int(cpu), int(mem))
	return nil
}
