// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

//go:build windows
// +build windows

// Package job limit cpu and memory in windows
package job

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

const MB = 1024 * 1024

var jobOpt *JobOptions

type JobOptions struct {
	CPUMax float64 `toml:"cpu_max"`
	MemMax int64   `toml:"mem_max_mb"`
}

func Run(opt *JobOptions) error {
	if opt == nil {
		return fmt.Errorf("opt is nil")
	}

	jobOpt = opt

	var (
		cpuInfo *JobObjectCPURateControlInformation
		memInfo *JobObjectExtendedLimitInformation
	)

	if opt.CPUMax == 0 && opt.MemMax == 0 {
		return fmt.Errorf("both CPUMax and MemMax are 0, ignore set cpu/mem limit")
	}

	pid := os.Getpid()
	name := "datakit"
	if u, err := uuid.NewRandom(); err == nil {
		name = u.String()
	}

	handle, err := CreateJobObject(nil, name)
	if err != nil {
		return fmt.Errorf("create job object failed: %w", err)
	}

	if opt.CPUMax > 0 {
		cpuInfo = &JobObjectCPURateControlInformation{
			CPURate: uint32(opt.CPUMax * 100),
		}
		cpuInfo.ControlFlags |= JOB_OBJECT_CPU_RATE_CONTROL_ENABLE
		cpuInfo.ControlFlags |= JOB_OBJECT_CPU_RATE_CONTROL_HARD_CAP

		if err := SetInformationJobObject_CPURateControlInformation(handle, cpuInfo); err != nil {
			return fmt.Errorf("set cpu limit error: %w", err)
		}
	}

	if opt.MemMax > 0 {
		memInfo = &JobObjectExtendedLimitInformation{}
		memInfo.ProcessMemoryLimit = uintptr(opt.MemMax * MB)
		memInfo.BasicLimitInformation.LimitFlags |= JOB_OBJECT_LIMIT_PROCESS_MEMORY

		if err := SetInformationJobObject_ExtendedLimitInformation(handle, memInfo); err != nil {
			return fmt.Errorf("set mem limit error: %w", err)
		}
	}

	processHandle, err := OpenProcess(
		STANDARD_RIGHTS_READ|PROCESS_QUERY_INFORMATION|SYNCHRONIZE|PROCESS_SET_INFORMATION,
		false,
		uint32(pid),
	)

	if err != nil {
		return fmt.Errorf("open process failed: %w", err)
	}

	defer CloseHandle(processHandle)

	if err := AssignProcessToJobObject(handle, processHandle); err != nil {
		return fmt.Errorf("assign process to job object failed: %s", err)
	}

	return nil
}

func Info() string {
	if jobOpt == nil {
		return "not ready"
	}

	return fmt.Sprintf("mem: %dMB, cpu: %.2f",
		jobOpt.MemMax, jobOpt.CPUMax)
}
