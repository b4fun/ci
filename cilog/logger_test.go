package cilog

import (
	"testing"

	"github.com/b4fun/ci"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	assert.IsType(t, &githubActionsT{}, Get(ci.GithubActions))
	assert.IsType(t, &azurePipelineT{}, Get(ci.AzurePipelines))
	assert.IsType(t, &Mute{}, Get(ci.Custom))
}
