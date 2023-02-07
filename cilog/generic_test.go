package cilog

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneric(t *testing.T) {
	b := new(bytes.Buffer)
	gt := generic(applyOptsFunc[genericT](func(gt *genericT) {
		gt.out = b
	}))

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
