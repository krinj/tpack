package tpack

import (
	"gotest.tools/assert"
	"testing"
)

func TestGetBestRapper(t *testing.T) {
	var result = GetBestRapper()
	assert.Equal(t, FamilyName, result)
}
