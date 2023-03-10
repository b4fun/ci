package cilog

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneric(t *testing.T) {
	b := new(bytes.Buffer)
	gt := generic()
	gt.SetOutput(b)

	caps := gt.Capabilities()
	assert.Contains(t, caps, CapabilityLog)
	assert.True(t, Can(gt, CapabilityLog))
	assert.False(t, Can(gt, CapabilityColor))

	gt.Log("foobar")

	output := b.String()
	t.Log(output)
	assert.Equal(
		t,
		`foobar
`,
		output,
	)
}

func TestGeneric_UnsupportedCaps(t *testing.T) {
	b := new(bytes.Buffer)
	gt := generic()
	gt.SetOutput(b)

	Debug(gt, "debug")
	Notice(gt, "notice")
	Warning(gt, "warning")
	Error(gt, "error")
	g, finishedGroup := Group(gt, GroupLogParams{Name: "group"})
	g.Log("group log")
	finishedGroup()

	output := b.String()
	t.Log(output)
	assert.Equal(
		t,
		`debug
notice
warning
error
group log
`,
		output,
	)
}
