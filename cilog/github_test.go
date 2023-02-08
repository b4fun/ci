package cilog

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitHubActions(t *testing.T) {
	b := new(bytes.Buffer)
	gat := GitHubActions()
	gat.SetOutput(b)

	caps := gat.Capabilities()
	assert.Contains(t, caps, CapabilityLog)
	assert.Contains(t, caps, CapabilityColor)
	assert.Contains(t, caps, CapabilityDebugLog)
	assert.Contains(t, caps, CapabilityNoticeLog)
	assert.Contains(t, caps, CapabilityWarningLog)
	assert.Contains(t, caps, CapabilityErrorLog)
	assert.Contains(t, caps, CapabilityGroupLog)

	gat.Log("foobar")
	gat.DebugLog("debug foobar")
	gat.NoticeLog("notice foobar")
	gat.WarningLog("warning foobar")
	gat.ErrorLog("error foobar")
	func() {
		ggat, closeGroup := gat.GroupLog(GroupLogParams{Name: "group"})
		defer closeGroup()

		ggat.Log("foobar")
		ggat.DebugLog("debug foobar")
		ggat.NoticeLog("notice foobar")
		ggat.WarningLog("warning foobar")
		ggat.ErrorLog("error foobar")
	}()

	output := b.String()
	t.Log(output)

	assert.Equal(
		t,
		`foobar
::debug::debug foobar
::notice::notice foobar
::warning::warning foobar
::error::error foobar
::group::group
foobar
::debug::debug foobar
::notice::notice foobar
::warning::warning foobar
::error::error foobar
::endgroup::
`,
		output,
	)
}
