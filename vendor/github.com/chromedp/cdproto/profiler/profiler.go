// Package profiler provides the Chrome DevTools Protocol
// commands, types, and events for the Profiler domain.
//
// Generated by the cdproto-gen command.
package profiler

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// DisableParams [no description].
type DisableParams struct{}

// Disable [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Profiler.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// EnableParams [no description].
type EnableParams struct{}

// Enable [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Profiler.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// GetBestEffortCoverageParams collect coverage data for the current isolate.
// The coverage data may be incomplete due to garbage collection.
type GetBestEffortCoverageParams struct{}

// GetBestEffortCoverage collect coverage data for the current isolate. The
// coverage data may be incomplete due to garbage collection.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-getBestEffortCoverage
func GetBestEffortCoverage() *GetBestEffortCoverageParams {
	return &GetBestEffortCoverageParams{}
}

// GetBestEffortCoverageReturns return values.
type GetBestEffortCoverageReturns struct {
	Result []*ScriptCoverage `json:"result,omitempty"` // Coverage data for the current isolate.
}

// Do executes Profiler.getBestEffortCoverage against the provided context.
//
// returns:
//
//	result - Coverage data for the current isolate.
func (p *GetBestEffortCoverageParams) Do(ctx context.Context) (result []*ScriptCoverage, err error) {
	// execute
	var res GetBestEffortCoverageReturns
	err = cdp.Execute(ctx, CommandGetBestEffortCoverage, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Result, nil
}

// SetSamplingIntervalParams changes CPU profiler sampling interval. Must be
// called before CPU profiles recording started.
type SetSamplingIntervalParams struct {
	Interval int64 `json:"interval"` // New sampling interval in microseconds.
}

// SetSamplingInterval changes CPU profiler sampling interval. Must be called
// before CPU profiles recording started.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-setSamplingInterval
//
// parameters:
//
//	interval - New sampling interval in microseconds.
func SetSamplingInterval(interval int64) *SetSamplingIntervalParams {
	return &SetSamplingIntervalParams{
		Interval: interval,
	}
}

// Do executes Profiler.setSamplingInterval against the provided context.
func (p *SetSamplingIntervalParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetSamplingInterval, p, nil)
}

// StartParams [no description].
type StartParams struct{}

// Start [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-start
func Start() *StartParams {
	return &StartParams{}
}

// Do executes Profiler.start against the provided context.
func (p *StartParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStart, nil, nil)
}

// StartPreciseCoverageParams enable precise code coverage. Coverage data for
// JavaScript executed before enabling precise code coverage may be incomplete.
// Enabling prevents running optimized code and resets execution counters.
type StartPreciseCoverageParams struct {
	CallCount             bool `json:"callCount,omitempty"`             // Collect accurate call counts beyond simple 'covered' or 'not covered'.
	Detailed              bool `json:"detailed,omitempty"`              // Collect block-based coverage.
	AllowTriggeredUpdates bool `json:"allowTriggeredUpdates,omitempty"` // Allow the backend to send updates on its own initiative
}

// StartPreciseCoverage enable precise code coverage. Coverage data for
// JavaScript executed before enabling precise code coverage may be incomplete.
// Enabling prevents running optimized code and resets execution counters.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-startPreciseCoverage
//
// parameters:
func StartPreciseCoverage() *StartPreciseCoverageParams {
	return &StartPreciseCoverageParams{}
}

// WithCallCount collect accurate call counts beyond simple 'covered' or 'not
// covered'.
func (p StartPreciseCoverageParams) WithCallCount(callCount bool) *StartPreciseCoverageParams {
	p.CallCount = callCount
	return &p
}

// WithDetailed collect block-based coverage.
func (p StartPreciseCoverageParams) WithDetailed(detailed bool) *StartPreciseCoverageParams {
	p.Detailed = detailed
	return &p
}

// WithAllowTriggeredUpdates allow the backend to send updates on its own
// initiative.
func (p StartPreciseCoverageParams) WithAllowTriggeredUpdates(allowTriggeredUpdates bool) *StartPreciseCoverageParams {
	p.AllowTriggeredUpdates = allowTriggeredUpdates
	return &p
}

// StartPreciseCoverageReturns return values.
type StartPreciseCoverageReturns struct {
	Timestamp float64 `json:"timestamp,omitempty"` // Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
}

// Do executes Profiler.startPreciseCoverage against the provided context.
//
// returns:
//
//	timestamp - Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
func (p *StartPreciseCoverageParams) Do(ctx context.Context) (timestamp float64, err error) {
	// execute
	var res StartPreciseCoverageReturns
	err = cdp.Execute(ctx, CommandStartPreciseCoverage, p, &res)
	if err != nil {
		return 0, err
	}

	return res.Timestamp, nil
}

// StartTypeProfileParams enable type profile.
type StartTypeProfileParams struct{}

// StartTypeProfile enable type profile.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-startTypeProfile
func StartTypeProfile() *StartTypeProfileParams {
	return &StartTypeProfileParams{}
}

// Do executes Profiler.startTypeProfile against the provided context.
func (p *StartTypeProfileParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStartTypeProfile, nil, nil)
}

// StopParams [no description].
type StopParams struct{}

// Stop [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-stop
func Stop() *StopParams {
	return &StopParams{}
}

// StopReturns return values.
type StopReturns struct {
	Profile *Profile `json:"profile,omitempty"` // Recorded profile.
}

// Do executes Profiler.stop against the provided context.
//
// returns:
//
//	profile - Recorded profile.
func (p *StopParams) Do(ctx context.Context) (profile *Profile, err error) {
	// execute
	var res StopReturns
	err = cdp.Execute(ctx, CommandStop, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Profile, nil
}

// StopPreciseCoverageParams disable precise code coverage. Disabling
// releases unnecessary execution count records and allows executing optimized
// code.
type StopPreciseCoverageParams struct{}

// StopPreciseCoverage disable precise code coverage. Disabling releases
// unnecessary execution count records and allows executing optimized code.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-stopPreciseCoverage
func StopPreciseCoverage() *StopPreciseCoverageParams {
	return &StopPreciseCoverageParams{}
}

// Do executes Profiler.stopPreciseCoverage against the provided context.
func (p *StopPreciseCoverageParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStopPreciseCoverage, nil, nil)
}

// StopTypeProfileParams disable type profile. Disabling releases type
// profile data collected so far.
type StopTypeProfileParams struct{}

// StopTypeProfile disable type profile. Disabling releases type profile data
// collected so far.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-stopTypeProfile
func StopTypeProfile() *StopTypeProfileParams {
	return &StopTypeProfileParams{}
}

// Do executes Profiler.stopTypeProfile against the provided context.
func (p *StopTypeProfileParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStopTypeProfile, nil, nil)
}

// TakePreciseCoverageParams collect coverage data for the current isolate,
// and resets execution counters. Precise code coverage needs to have started.
type TakePreciseCoverageParams struct{}

// TakePreciseCoverage collect coverage data for the current isolate, and
// resets execution counters. Precise code coverage needs to have started.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-takePreciseCoverage
func TakePreciseCoverage() *TakePreciseCoverageParams {
	return &TakePreciseCoverageParams{}
}

// TakePreciseCoverageReturns return values.
type TakePreciseCoverageReturns struct {
	Result    []*ScriptCoverage `json:"result,omitempty"`    // Coverage data for the current isolate.
	Timestamp float64           `json:"timestamp,omitempty"` // Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
}

// Do executes Profiler.takePreciseCoverage against the provided context.
//
// returns:
//
//	result - Coverage data for the current isolate.
//	timestamp - Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
func (p *TakePreciseCoverageParams) Do(ctx context.Context) (result []*ScriptCoverage, timestamp float64, err error) {
	// execute
	var res TakePreciseCoverageReturns
	err = cdp.Execute(ctx, CommandTakePreciseCoverage, nil, &res)
	if err != nil {
		return nil, 0, err
	}

	return res.Result, res.Timestamp, nil
}

// TakeTypeProfileParams collect type profile.
type TakeTypeProfileParams struct{}

// TakeTypeProfile collect type profile.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Profiler#method-takeTypeProfile
func TakeTypeProfile() *TakeTypeProfileParams {
	return &TakeTypeProfileParams{}
}

// TakeTypeProfileReturns return values.
type TakeTypeProfileReturns struct {
	Result []*ScriptTypeProfile `json:"result,omitempty"` // Type profile for all scripts since startTypeProfile() was turned on.
}

// Do executes Profiler.takeTypeProfile against the provided context.
//
// returns:
//
//	result - Type profile for all scripts since startTypeProfile() was turned on.
func (p *TakeTypeProfileParams) Do(ctx context.Context) (result []*ScriptTypeProfile, err error) {
	// execute
	var res TakeTypeProfileReturns
	err = cdp.Execute(ctx, CommandTakeTypeProfile, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Result, nil
}

// Command names.
const (
	CommandDisable               = "Profiler.disable"
	CommandEnable                = "Profiler.enable"
	CommandGetBestEffortCoverage = "Profiler.getBestEffortCoverage"
	CommandSetSamplingInterval   = "Profiler.setSamplingInterval"
	CommandStart                 = "Profiler.start"
	CommandStartPreciseCoverage  = "Profiler.startPreciseCoverage"
	CommandStartTypeProfile      = "Profiler.startTypeProfile"
	CommandStop                  = "Profiler.stop"
	CommandStopPreciseCoverage   = "Profiler.stopPreciseCoverage"
	CommandStopTypeProfile       = "Profiler.stopTypeProfile"
	CommandTakePreciseCoverage   = "Profiler.takePreciseCoverage"
	CommandTakeTypeProfile       = "Profiler.takeTypeProfile"
)
