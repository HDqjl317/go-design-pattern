/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-14 20:05:36
 * @LastEditTime: 2023-04-14 20:13:41
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct {
}

func (jsonParse jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParser struct {
}

func (yamlParse yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

func NewRuleConfigParse(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}
