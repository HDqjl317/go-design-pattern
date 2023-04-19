/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 22:14:12
 * @LastEditTime: 2023-04-19 11:48:18
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package interpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAlertRule_Interpret(t *testing.T) {
	stats := map[string]float64{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	tests := []struct {
		name  string
		stats map[string]float64
		rule  string
		want  bool
	}{
		{
			name:  "case1",
			stats: stats,
			rule:  "a > 1 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case2",
			stats: stats,
			rule:  "a < 2 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case3",
			stats: stats,
			rule:  "a < 5 && b > 1 && c < 10",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewAlertRule(tt.rule)
			require.NoError(t, err)
			assert.Equal(t, tt.want, r.Interpret(tt.stats))
		})
	}
}
