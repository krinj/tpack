package test

import (
	"github.com/krinj/tpack/pkg/spack"
	"gotest.tools/assert"
	"testing"
)

func TestGetBestRapper(t *testing.T) {
	var result = spack.GetBestRapper()
	assert.Equal(t, spack.FamilyName, result)
}
