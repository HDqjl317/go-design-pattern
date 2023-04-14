/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-14 20:45:52
 * @LastEditTime: 2023-04-14 21:21:09
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

type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

type jsonSystemConfigParser struct {
}

func (jsonSys jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct {
}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}
