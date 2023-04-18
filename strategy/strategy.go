/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 15:51:25
 * @LastEditTime: 2023-04-18 15:59:56
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package strategy

import (
	"fmt"
	"io/ioutil"
	"os"
)

// StorageStrategy 存储策略
type StorageStrategy interface {
	Save(name string, data []byte) error
}

var strategy = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := strategy[t]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy %s", t)
	}
	return s, nil
}

// 保存到文件
type fileStorage struct {
}

func (s *fileStorage) Save(name string, data []byte) error {
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

// 加密保存到文件
type encryptFileStorage struct {
}

func (s *encryptFileStorage) Save(name string, data []byte) error {
	data, err := encrypt(data)
	if err != nil {
		return nil
	}
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

func encrypt(data []byte) ([]byte, error) {
	// some code
	return data, nil
}
