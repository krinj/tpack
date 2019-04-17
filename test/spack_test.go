package test

import (
	"github.com/krinj/tpack/pkg"
	"gotest.tools/assert"
	"testing"
)

func TestGetBestRapper(t *testing.T) {
	var result = pkg.GetBestRapper()
	assert.Equal(t, pkg.FamilyName, result)
}
