package cilog

import (
	"testing"

	"github.com/b4fun/ci"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	assert.IsType(t, &githubActionsT{}, Get(ci.GithubActions))
	assert.IsType(t, &azurePipelineT{}, Get(ci.AzurePipelines))
	assert.IsType(t, &genericT{}, Get(ci.Custom))
}

func TestGroup_GroupUnsupported(t *testing.T) {
	l := &Mute{}
	subLogger, finish := Group(l, GroupLogParams{Name: "group"})
	assert.NotNil(t, finish)
	assert.Equal(t, l, subLogger)
}
