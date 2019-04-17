package test

import (
	"github.com/krinj/tpack/pkg/tpack"
	"testing"
)
import "gotest.tools/assert"

func TestExecute(t *testing.T) {
	var result = tpack.Execute(100)
	assert.Equal(t, result, 120)
}

