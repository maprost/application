package util_test

import (
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/assertion"
	"testing"
)

func TestJoinStrings(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(util.JoinStrings("", "sep", ""), "")
	assert.Equal(util.JoinStrings("A", "sep", ""), "A")
	assert.Equal(util.JoinStrings("", "sep", "B"), "B")
	assert.Equal(util.JoinStrings("A", "sep", "B"), "AsepB")
}
