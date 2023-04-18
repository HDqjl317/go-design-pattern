/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 20:48:47
 * @LastEditTime: 2023-04-18 20:48:54
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensitiveWordFilterChain_Filter(t *testing.T) {
	chain := &SensitiveWordFilterChain{}
	chain.AddFilter(&AdSensitiveWordFilter{})
	assert.Equal(t, false, chain.Filter("test"))

	chain.AddFilter(&PoliticalWordFilter{})
	assert.Equal(t, true, chain.Filter("test"))
}
