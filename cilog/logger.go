package cilog

import (
	"io"

	"github.com/b4fun/ci"
)

type Capability int

const (
	// CapabilityLog - prints a basic log
	CapabilityLog Capability = iota
	// CapabilityColor - prints log with color
	CapabilityColor
	// CapabilityDebugLog - emitting debug log
	CapabilityDebugLog
	// CapabilityNoticeLog - emitting notice log
	CapabilityNoticeLog
	// CapabilityWarningLog - emitting warning log
	CapabilityWarningLog
	// CapabilityErrorLog - emitting error log
	CapabilityErrorLog
	// CapabilityGroupLog - emitting logs as a group
	CapabilityGroupLog
)

// Can tells if the logger has specified logging capability.
func Can(l Logger, capability Capability) bool {
	_, exists := l.Capabilities()[capability]
	return exists
}

// Group starts a group logger.
// If the given logger doesn't support group logging,
// the original logger will be returned and a noop finish function will be returned.
func Group(l Logger, params GroupLogParams) (Logger, func()) {
	if !Can(l, CapabilityGroupLog) {
		return l, func() {}
	}

	return l.GroupLog(params)
}

// Debug - emitting debug log.
// If the given logger doesn't support CapabilityDebugLog, fallbacks to Log.
func Debug(l Logger, s string) {
	p := l.DebugLog

	if !Can(l, CapabilityDebugLog) {
		p = l.Log
	}

	p(s)
}

// Notice - emitting notice log.
// If the given logger doesn't support CapabilityNoticeLog, fallbacks to Log.
func Notice(l Logger, s string) {
	p := l.DebugLog

	if !Can(l, CapabilityNoticeLog) {
		p = l.Log
	}

	p(s)
}

// Warning - emitting warning log.
// If the given logger doesn't support CapabilityWarningLog, fallbacks to Log.
func Warning(l Logger, s string) {
	p := l.WarningLog

	if !Can(l, CapabilityWarningLog) {
		p = l.Log
	}

	p(s)
}

// Error - emitting error log.
// If the given logger doesn't support CapabilityErrorLog, fallbacks to Log.
func Error(l Logger, s string) {
	p := l.ErrorLog

	if !Can(l, CapabilityErrorLog) {
		p = l.Log
	}

	p(s)
}

// GroupLogParams specifies the parameter for GroupLog.
type GroupLogParams struct {
	// Name - name of the group
	Name string
}

// Logger provides improved CI logging output.
type Logger interface {
	// SetOutput sets the output of the logger.
	// This is useful unit testing usage.
	SetOutput(io.Writer)

	// Capabilities returns the capabilities of the logger.
	Capabilities() map[Capability]struct{}

	// Log - prints a basic log (CapabilityLog)
	Log(s string)

	// DebugLog - emitting debug log (CapabilityDebugLog)
	DebugLog(s string)

	// NoticeLog - emitting notice log (CapabilityNoticeLog)
	NoticeLog(s string)

	// WarningLog - emitting warning log (CapabilityWarningLog)
	WarningLog(s string)

	// ErrorLog - emitting error log (CapabilityErrorLog)
	ErrorLog(s string)

	// GroupLog - emitting logs as a group (CapabilityGroupLog)
	GroupLog(params GroupLogParams) (groupLogger Logger, endGroup func())
}

// Get returns a logger based on CI environment.
// If no supported logger defined, a generic logger will be returned.
func Get(name ci.Name) Logger {
	switch name {
	case ci.GithubActions:
		return GitHubActions()
	case ci.AzurePipelines:
		return AzurePipeline()
	default:
		return generic()
	}
}
