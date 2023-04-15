/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-15 10:56:08
 * @LastEditTime: 2023-04-15 11:02:05
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package prototype

import (
	"encoding/json"
	"time"
)

type Keyword struct {
	word     string
	visit    int
	UpdateAt *time.Time
}

// Clone 这里用序列化与反序列化的方式进行深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

type Keywords map[string]*Keyword

func (words Keywords) Clone(updateWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range words {
		newKeywords[k] = v
	}

	for _, word := range updateWords {
		newKeywords[word.word] = word.Clone()
	}
	return newKeywords
}
