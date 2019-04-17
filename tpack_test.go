package tpack

import "testing"
import "gotest.tools/assert"

func TestExecute(t *testing.T) {
	var result = Execute(100)
	assert.Equal(t, result, 120)
}

