/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 16:00:13
 * @LastEditTime: 2023-04-18 16:06:50
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package strategy

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_demo(t *testing.T) {
	data, sensitive := getData()
	strategyType := "file"
	if sensitive {
		strategyType = "encrypt_file"
	}
	storage, err := NewStorageStrategy(strategyType)
	assert.NoError(t, err)

	dir, err := os.MkdirTemp("", "strategy")
	assert.NoError(t, err)
	assert.NoError(t, storage.Save(dir+"/test.txt", data))
}

func getData() ([]byte, bool) {
	return []byte("test data"), false
}
