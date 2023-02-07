package cilog

import (
	"fmt"
	"io"
	"os"
)

// GitHubOpts configures GitHubActions logger.
type GitHubOpts = applyOpts[githubActionsT]

type githubActionsT struct {
	Mute

	out io.Writer // reserve for testing only for now
}

var _ Logger = (*githubActionsT)(nil)

// GitHubActions creates a GitHubActions logger.
//
// refs:
// - https://github.com/actions/toolkit/blob/master/docs/commands.md
func GitHubActions(opts ...GitHubOpts) Logger {
	rv := &githubActionsT{
		out: os.Stdout,
	}

	for _, o := range opts {
		o.apply(rv)
	}

	return rv
}

func (gat *githubActionsT) Capabilities() map[Capability]struct{} {
	return map[Capability]struct{}{
		CapabilityLog:        {},
		CapabilityColor:      {},
		CapabilityDebugLog:   {},
		CapabilityNoticeLog:  {},
		CapabilityWarningLog: {},
		CapabilityErrorLog:   {},
		CapabilityGroupLog:   {},
	}
}

func (gat *githubActionsT) logfln(s string, a ...interface{}) {
	fmt.Fprintf(gat.out, s, a...)
	fmt.Fprintln(gat.out)
}

func (gat *githubActionsT) Log(s string) {
	gat.logfln(s)
}

func (gat *githubActionsT) DebugLog(s string) {
	gat.logfln("::debug::%s", s)
}

func (gat *githubActionsT) NoticeLog(s string) {
	gat.logfln("::notice::%s", s)
}

func (gat *githubActionsT) WarningLog(s string) {
	gat.logfln("::warning::%s", s)
}

func (gat *githubActionsT) ErrorLog(s string) {
	gat.logfln("::error::%s", s)
}

func (gat *githubActionsT) GroupLog(params GroupLogParams) (Logger, func()) {
	gat.logfln("::group::%s", params.Name)
	closeSection := func() {
		gat.logfln("::endgroup::")
	}
	return gat, closeSection
}
