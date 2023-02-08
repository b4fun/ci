package cilog

import (
	"fmt"
	"io"
	"os"
)

// AzurePipelineOpts configures AzurePipeline logger.
type AzurePipelineOpts = applyOpts[azurePipelineT]

// AzurePipelineUseLogIssue controls AzurePipeline logger to log with LogIssue command.
//
// ref: https://learn.microsoft.com/en-us/azure/devops/pipelines/scripts/logging-commands?view=azure-devops&tabs=bash#logissue-log-an-error-or-warning
func AzurePipelineUseLogIssue(yes bool) AzurePipelineOpts {
	return applyOptsFunc[azurePipelineT](func(apt *azurePipelineT) {
		apt.useLogIssue = yes
	})
}

// azurePipelineT implements Logger for AzurePipelines CI.
type azurePipelineT struct {
	Mute

	useLogIssue bool

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

func (apt *azurePipelineT) SetOutput(out io.Writer) {
	apt.out = out
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
	if apt.useLogIssue {
		apt.logfln("##vso[task.logissue type=warning]%s", s)
	} else {
		apt.logfln("##[warning]%s", s)
	}
}

func (apt *azurePipelineT) ErrorLog(s string) {
	if apt.useLogIssue {
		apt.logfln("##vso[task.logissue type=error]%s", s)
	} else {
		apt.logfln("##[error]%s", s)
	}
}

func (apt *azurePipelineT) GroupLog(params GroupLogParams) (Logger, func()) {
	apt.logfln("##[group]%s", params.Name)
	closeSection := func() {
		apt.logfln("##[endgroup]")
	}
	return apt, closeSection
}
