/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 09:34:43
 * @LastEditTime: 2023-04-18 09:38:19
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package adapter

import "testing"

func TestAliyunClientAdapter_CreateServer(t *testing.T) {
	// 确保 adapter 实现了目标接口
	var a ICreateServer = &AliyunClientAdapter{
		Client: AliyunClient{},
	}
	a.CreateServer(1.0, 2.0)
}

func TestAwsClientAdapter_CreateServer(t *testing.T) {
	// 确保 adapter 实现了目标接口
	var a ICreateServer = &AwsClientAdpter{
		Client: AWSClient{},
	}
	a.CreateServer(1.0, 2.0)
}
