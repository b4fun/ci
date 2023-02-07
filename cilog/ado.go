package cilog

import (
	"fmt"
	"io"
	"os"
)

// AzurePipelineOpts configures AzurePipeline logger.
type AzurePipelineOpts = applyOpts[azurePipelineT]

// azurePipelineT implements Logger for AzurePipelines CI.
type azurePipelineT struct {
	Mute

	out io.Writer // reserve for testing only for now
}

var _ Logger = (*azurePipelineT)(nil)

// AzurePipeline creates an AzurePipeline logger.
//
// refs:
// - https://learn.microsoft.com/en-us/azure/devops/pipelines/scripts/logging-commands
func AzurePipeline(opts ...AzurePipelineOpts) Logger {
	rv := &azurePipelineT{
		out: os.Stdout,
	}

	for _, o := range opts {
		o.apply(rv)
	}

	return rv
}

func (apt *azurePipelineT) Capabilities() map[Capability]struct{} {
	return map[Capability]struct{}{
		CapabilityLog:        {},
		CapabilityDebugLog:   {},
		CapabilityWarningLog: {},
		CapabilityErrorLog:   {},
		CapabilityGroupLog:   {},
	}
}

func (apt *azurePipelineT) logfln(s string, a ...interface{}) {
	fmt.Fprintf(apt.out, s, a...)
	fmt.Fprintln(apt.out)
}

func (apt *azurePipelineT) Log(s string) {
	apt.logfln(s)
}

func (apt *azurePipelineT) DebugLog(s string) {
	apt.logfln("##[debug]%s", s)
}

func (apt *azurePipelineT) WarningLog(s string) {
	apt.logfln("##[warning]%s", s)
}

func (apt *azurePipelineT) ErrorLog(s string) {
	apt.logfln("##[error]%s", s)
}

func (apt *azurePipelineT) GroupLog(params GroupLogParams) (Logger, func()) {
	apt.logfln("##[group]%s", params.Name)
	closeSection := func() {
		apt.logfln("##[endgroup]")
	}
	return apt, closeSection
}
