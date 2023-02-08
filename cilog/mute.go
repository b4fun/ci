package cilog

import "io"

// Mute implements all Logger methods by muting output.
type Mute struct {
}

var _ Logger = (*Mute)(nil)

func (m *Mute) SetOutput(o io.Writer) {}

func (m *Mute) Capabilities() map[Capability]struct{} {
	return map[Capability]struct{}{}
}

func (m *Mute) Log(s string)        {}
func (m *Mute) DebugLog(s string)   {}
func (m *Mute) NoticeLog(s string)  {}
func (m *Mute) WarningLog(s string) {}
func (m *Mute) ErrorLog(s string)   {}

func (m *Mute) GroupLog(GroupLogParams) (Logger, func()) {
	return m, func() {}
}
