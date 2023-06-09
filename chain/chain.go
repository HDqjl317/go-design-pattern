/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 20:37:10
 * @LastEditTime: 2023-04-18 20:48:16
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package chain

// SensitiveWordFilter 敏感词过滤, 判定是否是敏感词
type SensitiveWordFilter interface {
	Filter(content string) bool
}

// SensitiveWordFilterChain 职责链
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

// 添加一个过滤器
func (c *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	c.filters = append(c.filters, filter)
}

// 执行过滤
func (c *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range c.filters {
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

type AdSensitiveWordFilter struct{}

func (f *AdSensitiveWordFilter) Filter(content string) bool {
	// some code
	return false
}

type PoliticalWordFilter struct{}

func (f *PoliticalWordFilter) Filter(content string) bool {
	// some code
	return true
}
