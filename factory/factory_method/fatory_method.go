/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-14 20:27:33
 * @LastEditTime: 2023-04-14 20:37:19
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

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type jsonRuleConfigParserFactory struct {
}

func (jParseF jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

type yamlRuleConfigParserFactory struct {
}

func (yParseF yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}
