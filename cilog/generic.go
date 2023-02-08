package cilog

import (
	"fmt"
	"io"
	"os"
)

type genericT struct {
	Mute

	out io.Writer
}

func generic(opts ...applyOpts[genericT]) Logger {
	rv := &genericT{
		out: os.Stdout,
	}

	for _, o := range opts {
		o.apply(rv)
	}

	return rv
}

var _ Logger = (*genericT)(nil)

func (gt *genericT) SetOutput(out io.Writer) {
	gt.out = out
}

func (gt *genericT) Capabilities() map[Capability]struct{} {
	return map[Capability]struct{}{
		CapabilityLog: {},
	}
}

func (gt *genericT) logfln(s string, a ...interface{}) {
	fmt.Fprintf(gt.out, s, a...)
	fmt.Fprintln(gt.out)
}

func (gt *genericT) Log(s string) {
	gt.logfln(s)
}
