/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-14 20:38:31
 * @LastEditTime: 2023-04-14 20:38:31
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package factory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParserFactory
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParserFactory{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParserFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParserFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
