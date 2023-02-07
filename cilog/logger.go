package cilog

import "github.com/b4fun/ci"

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

// GroupLogParams specifies the parameter for GroupLog.
type GroupLogParams struct {
	// Name - name of the group
	Name string
}

// Logger provides improved CI logging output.
type Logger interface {
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
