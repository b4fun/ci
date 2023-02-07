package cilog

import (
	"bytes"
	"testing"

	"github.com/b4fun/ci"
	"github.com/stretchr/testify/assert"
)

func TestAzurePipeline(t *testing.T) {
	b := new(bytes.Buffer)
	apt := AzurePipeline(applyOptsFunc[azurePipelineT](func(apt *azurePipelineT) {
		apt.out = b
	}))

	caps := apt.Capabilities()
	assert.Contains(t, caps, CapabilityLog)
	assert.Contains(t, caps, CapabilityDebugLog)
	assert.Contains(t, caps, CapabilityWarningLog)
	assert.Contains(t, caps, CapabilityErrorLog)
	assert.Contains(t, caps, CapabilityGroupLog)

	apt.Log("foobar")
	apt.DebugLog("debug foobar")
	apt.WarningLog("warning foobar")
	apt.ErrorLog("error foobar")
	func() {
		gapt, closeGroup := apt.GroupLog(GroupLogParams{Name: "group"})
		defer closeGroup()

		gapt.Log("foobar")
		gapt.DebugLog("debug foobar")
		gapt.WarningLog("warning foobar")
		gapt.ErrorLog("error foobar")
	}()

	output := b.String()
	if ci.Detect() != ci.AzurePipelines {
		t.Log(output)
	}

	assert.Equal(
		t,
		`foobar
##[debug]debug foobar
##[warning]warning foobar
##[error]error foobar
##[group]group
foobar
##[debug]debug foobar
##[warning]warning foobar
##[error]error foobar
##[endgroup]
`,
		output,
	)
}

func TestAzurePipeline_useLogIssue(t *testing.T) {
	b := new(bytes.Buffer)
	apt := AzurePipeline(
		applyOptsFunc[azurePipelineT](func(apt *azurePipelineT) {
			apt.out = b
		}),
		AzurePipelineUseLogIssue(true),
	)

	apt.Log("foobar")
	apt.DebugLog("debug foobar")
	apt.WarningLog("warning foobar")
	apt.ErrorLog("error foobar")

	output := b.String()
	if ci.Detect() != ci.AzurePipelines {
		t.Log(output)
	}

	assert.Equal(
		t,
		`foobar
##[debug]debug foobar
##vso[task.logissue type=warning]warning foobar
##vso[task.logissue type=error]error foobar
`,
		output,
	)
}
