package test

import (
	"github.com/krinj/tpack/pkg"
	"testing"
)
import "gotest.tools/assert"

func TestExecute(t *testing.T) {
	var result = pkg.Execute(100)
	assert.Equal(t, result, 120)
}

