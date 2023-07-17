// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package sqlserver

import (
	"github.com/GuanceCloud/cliutils/point"
	dkpt "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/io/point"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/plugins/inputs"
)

type Measurement struct {
	name     string
	tags     map[string]string
	fields   map[string]interface{}
	election bool
}

func (m *Measurement) Point() *point.Point {
	return nil
}

func (m *Measurement) Info() *inputs.MeasurementInfo {
	return nil
}

type MetricMeasurment struct {
	Measurement
}

func (m *MetricMeasurment) LineProto() (*dkpt.Point, error) {
	return dkpt.NewPoint(m.name, m.tags, m.fields, dkpt.MOptElection())
}

// Point implement MeasurementV2.
func (m *MetricMeasurment) Point() *point.Point {
	opts := point.DefaultMetricOptions()

	if m.election {
		opts = append(opts, point.WithExtraTags(dkpt.GlobalElectionTags()))
	}

	return point.NewPointV2([]byte(m.name),
		append(point.NewTags(m.tags), point.NewKVs(m.fields)...),
		opts...)
}

type LoggingMeasurment struct {
	Measurement
}

func (m *LoggingMeasurment) LineProto() (*dkpt.Point, error) {
	return dkpt.NewPoint(m.name, m.tags, m.fields, dkpt.LOptElection())
}

// Point implement MeasurementV2.
func (m *LoggingMeasurment) Point() *point.Point {
	opts := point.DefaultLoggingOptions()

	if m.election {
		opts = append(opts, point.WithExtraTags(dkpt.GlobalElectionTags()))
	}

	return point.NewPointV2([]byte(m.name),
		append(point.NewTags(m.tags), point.NewKVs(m.fields)...),
		opts...)
}

type SqlserverMeasurment struct {
	MetricMeasurment
}

//nolint:lll
func (m *SqlserverMeasurment) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver",
		Type: "metric",
		Fields: map[string]interface{}{
			"cpu_count":           newCountFieldInfo("Specifies the number of logical CPUs on the system. Not nullable"),
			"uptime":              newTimeFieldInfo("Total time elapsed since the last computer restart"),
			"committed_memory":    newByteFieldInfo("The amount of memory committed to the memory manager"),
			"physical_memory":     newByteFieldInfo("Total physical memory on the machine"),
			"virtual_memory":      newByteFieldInfo("Amount of virtual memory available to the process in user mode."),
			"target_memory":       newByteFieldInfo("Amount of memory that can be consumed by the memory manager. When this value is larger than the committed memory, then the memory manager will try to obtain more memory. When it is smaller, the memory manager will try to shrink the amount of memory committed."),
			"db_online":           newCountFieldInfo("num of database state in online"),
			"db_offline":          newCountFieldInfo("num of database state in offline"),
			"db_recovering":       newCountFieldInfo("num of database state in recovering"),
			"db_recovery_pending": newCountFieldInfo("num of database state in recovery_pending"),
			"db_restoring":        newCountFieldInfo("num of database state in restoring"),
			"db_suspect":          newCountFieldInfo("num of database state in suspect"),
			"server_memory":       newByteFieldInfo("memory used"),
		},
		Tags: map[string]interface{}{
			"sqlserver_host": inputs.NewTagInfo("host name which installed SQLServer"),
		},
	}
}

type Performance struct {
	MetricMeasurment
}

//nolint:lll
func (m *Performance) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_performance",
		Type: "metric",
		Desc: "performance counter maintained by the server,[detail](https://docs.microsoft.com/en-us/sql/relational-databases/system-dynamic-management-views/sys-dm-os-performance-counters-transact-sql?view=sql-server-ver15)",
		Fields: map[string]interface{}{
			"cntr_value": &inputs.FieldInfo{
				DataType: inputs.Float,
				Type:     inputs.Count,
				Unit:     inputs.NCount,
				Desc:     "Current value of the counter",
			},
		},
		Tags: map[string]interface{}{
			"object_name":    inputs.NewTagInfo("Category to which this counter belongs."),
			"counter_name":   inputs.NewTagInfo("Name of the counter. To get more information about a counter, this is the name of the topic to select from the list of counters in Use SQL Server Objects."),
			"counter_type":   inputs.NewTagInfo("Type of the counter"),
			"instance":       inputs.NewTagInfo("Name of the specific instance of the counter"),
			"sqlserver_host": inputs.NewTagInfo("host name which installed SQLServer"),
		},
	}
}

type WaitStatsCategorized struct {
	MetricMeasurment
}

//nolint:lll
func (m *WaitStatsCategorized) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_waitstats",
		Type: "metric",
		Desc: "information about all the waits encountered by threads that executed,[detail](https://docs.microsoft.com/en-us/sql/relational-databases/system-dynamic-management-views/sys-dm-os-wait-stats-transact-sql?view=sql-server-ver15)",
		Fields: map[string]interface{}{
			"max_wait_time_ms":    newTimeFieldInfo("Maximum wait time on this wait type."),
			"wait_time_ms":        newTimeFieldInfo("Total wait time for this wait type in milliseconds. This time is inclusive of signal_wait_time_ms"),
			"signal_wait_time_ms": newTimeFieldInfo("Difference between the time that the waiting thread was signaled and when it started running"),
			"resource_wait_ms":    newTimeFieldInfo("wait_time_ms-signal_wait_time_ms"),
			"waiting_tasks_count": newCountFieldInfo("Number of waits on this wait type. This counter is incremented at the start of each wait."),
		},
		Tags: map[string]interface{}{
			"sqlserver_host": inputs.NewTagInfo("host name which installed SQLServer"),
			"wait_type":      inputs.NewTagInfo("Name of the wait type. For more information, see Types of Waits, later in this topic"),
			"wait_category":  inputs.NewTagInfo("wait category info"),
		},
	}
}

type DatabaseIO struct {
	MetricMeasurment
}

//nolint:lll
func (m *DatabaseIO) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_database_io",
		Type: "metric",
		Desc: "I/O statistics for data and log files,[detail](https://docs.microsoft.com/en-us/sql/relational-databases/system-dynamic-management-views/sys-dm-io-virtual-file-stats-transact-sql?view=sql-server-ver15)",
		Fields: map[string]interface{}{
			"read_bytes":        newByteFieldInfo("Total number of bytes read on this file"),
			"write_bytes":       newByteFieldInfo("Number of writes made on this file"),
			"read_latency_ms":   newTimeFieldInfo("Total time, in milliseconds, that the users waited for reads issued on the file."),
			"write_latency_ms":  newTimeFieldInfo("Total time, in milliseconds, that users waited for writes to be completed on the file"),
			"reads":             newCountFieldInfo("Number of reads issued on the file."),
			"writes":            newCountFieldInfo("Number of writes issued on the file."),
			"rg_read_stall_ms":  newTimeFieldInfo("Does not apply to:: SQL Server 2008 through SQL Server 2012 (11.x).Total IO latency introduced by IO resource governance for reads"),
			"rg_write_stall_ms": newTimeFieldInfo("Does not apply to:: SQL Server 2008 through SQL Server 2012 (11.x).Total IO latency introduced by IO resource governance for writes. Is not nullable."),
		},
		Tags: map[string]interface{}{
			"database_name":     inputs.NewTagInfo("database name"),
			"file_type":         inputs.NewTagInfo("Description of the file type, `ROWS/LOG/FILESTREAM/FULLTEXT` (Full-text catalogs earlier than SQL Server 2008.)"),
			"logical_filename":  inputs.NewTagInfo("Logical name of the file in the database"),
			"physical_filename": inputs.NewTagInfo("Operating-system file name."),
			"sqlserver_host":    inputs.NewTagInfo("host name which installed SQLServer"),
		},
	}
}

type Schedulers struct {
	MetricMeasurment
}

//nolint:lll
func (m *Schedulers) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_schedulers",
		Type: "metric",
		Desc: "one row per scheduler in SQL Server where each scheduler is mapped to an individual processor,[detail](https://docs.microsoft.com/en-us/sql/relational-databases/system-dynamic-management-views/sys-dm-os-schedulers-transact-sql?view=sql-server-ver15)",
		Fields: map[string]interface{}{
			"active_workers_count":      newCountFieldInfo("Number of workers that are active. An active worker is never preemptive, must have an associated task, and is either running, runnable, or suspended. Is not nullable."),
			"context_switches_count":    newCountFieldInfo("Number of context switches that have occurred on this scheduler"),
			"current_tasks_count":       newCountFieldInfo("Number of current tasks that are associated with this scheduler."),
			"current_workers_count":     newCountFieldInfo("Number of workers that are associated with this scheduler. This count includes workers that are not assigned any task. Is not nullable."),
			"is_idle":                   newBoolFieldInfo("Scheduler is idle. No workers are currently running"),
			"is_online":                 newBoolFieldInfo("If SQL Server is configured to use only some of the available processors on the server, this configuration can mean that some schedulers are mapped to processors that are not in the affinity mask. If that is the case, this column returns 0. This value means that the scheduler is not being used to process queries or batches."),
			"load_factor":               newCountFieldInfo("Internal value that indicates the perceived load on this scheduler"),
			"pending_disk_io_count":     newCountFieldInfo("Number of pending I/Os that are waiting to be completed."),
			"preemptive_switches_count": newCountFieldInfo("Number of times that workers on this scheduler have switched to the preemptive mode"),
			"runnable_tasks_count":      newCountFieldInfo("Number of workers, with tasks assigned to them, that are waiting to be scheduled on the runnable queue."),
			"total_cpu_usage_ms":        newTimeFieldInfo("Applies to: SQL Server 2016 (13.x) and laterTotal CPU consumed by this scheduler as reported by non-preemptive workers."),
			"total_scheduler_delay_ms":  newTimeFieldInfo("Applies to: SQL Server 2016 (13.x) and laterThe time between one worker switching out and another one switching in"),
			"work_queue_count":          newCountFieldInfo("Number of tasks in the pending queue. These tasks are waiting for a worker to pick them up"),
			"yield_count":               newCountFieldInfo("Internal value that is used to indicate progress on this scheduler. This value is used by the Scheduler Monitor to determine whether a worker on the scheduler is not yielding to other workers on time."),
		},
		Tags: map[string]interface{}{
			"cpu_id":         inputs.NewTagInfo("CPU ID assigned to the scheduler."),
			"sqlserver_host": inputs.NewTagInfo("host name which installed SQLServer"),
			"scheduler_id":   inputs.NewTagInfo("ID of the scheduler. All schedulers that are used to run regular queries have ID numbers less than 1048576. Those schedulers that have IDs greater than or equal to 1048576 are used internally by SQL Server, such as the dedicated administrator connection scheduler. Is not nullable."),
		},
	}
}

type VolumeSpace struct {
	MetricMeasurment
}

//nolint:lll
func (m *VolumeSpace) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_volumespace",
		Type: "metric",
		Fields: map[string]interface{}{
			"volume_available_space_bytes": newByteFieldInfo("Available free space on the volume"),
			"volume_total_space_bytes":     newByteFieldInfo("Total size in bytes of the volume"),
			"volume_used_space_bytes":      newByteFieldInfo("Used size in bytes of the volume"),
		},
		Tags: map[string]interface{}{
			"sqlserver_host":     inputs.NewTagInfo("host name which installed SQLServer"),
			"volume_mount_point": inputs.NewTagInfo("Mount point at which the volume is rooted. Can return an empty string. Returns null on Linux operating system."),
		},
	}
}

type LockRow struct {
	LoggingMeasurment
}

//nolint:lll
func (m *LockRow) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_lock_row",
		Type: "logging",
		Fields: map[string]interface{}{
			"blocking_session_id":     newCountFieldInfo("ID of the session that is blocking the request"),
			"session_id":              newCountFieldInfo("ID of the session to which this request is related"),
			"cpu_time":                newTimeFieldInfo("CPU time in milliseconds that is used by the request"),
			"logical_reads":           newCountFieldInfo("Number of logical reads that have been performed by the request"),
			"row_count":               newCountFieldInfo("Number of rows returned on the session up to this point"),
			"memory_usage":            newCountFieldInfo("Number of 8-KB pages of memory used by this session"),
			"last_request_start_time": newTimeFieldInfo("Time at which the last request on the session began, in second"),
			"last_request_end_time":   newTimeFieldInfo("Time of the last completion of a request on the session, in second"),
			"host_name":               newStringFieldInfo("Name of the client workstation that is specific to a session"),
			"login_name":              newStringFieldInfo("SQL Server login name under which the session is currently executing"),
			"session_status":          newStringFieldInfo("Status of the session"),
			"message":                 newStringFieldInfo("Text of the SQL query"),
		},
		Tags: map[string]interface{}{},
	}
}

type LockTable struct {
	LoggingMeasurment
}

//nolint:lll
func (m *LockTable) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_lock_table",
		Type: "logging",
		Fields: map[string]interface{}{
			"request_session_id": newCountFieldInfo("Session ID that currently owns this request"),
			"object_name":        newStringFieldInfo("Name of the entity in a database with which a resource is associated"),
			"db_name":            newStringFieldInfo("Name of the database under which this resource is scoped"),
			"resource_type":      newStringFieldInfo("Represents the resource type"),
			"request_mode":       newStringFieldInfo("Mode of the request"),
			"request_status":     newStringFieldInfo("Current status of this request"),
		},
		Tags: map[string]interface{}{},
	}
}

type LockDead struct {
	LoggingMeasurment
}

//nolint:lll
func (m *LockDead) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_lock_dead",
		Type: "logging",
		Fields: map[string]interface{}{
			"request_session_id":   newCountFieldInfo("Session ID that currently owns this request"),
			"blocking_session_id":  newCountFieldInfo("ID of the session that is blocking the request"),
			"blocking_object_name": newStringFieldInfo("Indicates the name of the object to which this partition belongs"),
			"db_name":              newStringFieldInfo("Name of the database under which this resource is scoped"),
			"resource_type":        newStringFieldInfo("Represents the resource type"),
			"request_mode":         newStringFieldInfo("Mode of the request"),
			"requesting_text":      newStringFieldInfo("Text of the SQL query which is requesting"),
			"blocking_text":        newStringFieldInfo("Text of the SQL query which is blocking"),
			"message":              newStringFieldInfo("Text of the SQL query which is blocking"),
		},
		Tags: map[string]interface{}{},
	}
}

type LogicalIO struct {
	LoggingMeasurment
}

//nolint:lll
func (m *LogicalIO) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_logical_io",
		Type: "logging",
		Fields: map[string]interface{}{
			"avg_logical_io":       newCountFieldInfo("Average number of logical writes and logical reads"),
			"total_logical_io":     newCountFieldInfo("Total number of logical writes and logical reads"),
			"total_logical_reads":  newCountFieldInfo("Total amount of logical reads"),
			"total_logical_writes": newCountFieldInfo("Total amount of logical writes"),
			"creation_time":        newCountFieldInfo("The Unix time at which the plan was compiled, in millisecond"),
			"execution_count":      newCountFieldInfo("Number of times that the plan has been executed since it was last compiled"),
			"last_execution_time":  newCountFieldInfo("Last time at which the plan started executing, unix time in millisecond"),
		},
		Tags: map[string]interface{}{
			"message": inputs.NewTagInfo("Text of the SQL query"),
		},
	}
}

type WorkerTime struct {
	LoggingMeasurment
}

//nolint:lll
func (m *WorkerTime) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_worker_time",
		Type: "logging",
		Fields: map[string]interface{}{
			"creation_time":       newCountFieldInfo("The Unix time at which the plan was compiled, in millisecond"),
			"execution_count":     newCountFieldInfo("Number of times that the plan has been executed since it was last compiled"),
			"last_execution_time": newCountFieldInfo("Last time at which the plan started executing, unix time in millisecond"),
			"total_worker_time":   newCountFieldInfo("Total amount of CPU time, reported in milliseconds"),
			"avg_worker_time":     newCountFieldInfo("Average amount of CPU time, reported in milliseconds"),
		},
		Tags: map[string]interface{}{
			"message": inputs.NewTagInfo("Text of the SQL query"),
		},
	}
}

type DatabaseSize struct {
	MetricMeasurment
}

//nolint:lll
func (m *DatabaseSize) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_database_size",
		Type: "metric",
		Fields: map[string]interface{}{
			"data_size": newKByteFieldInfo("The size of file of Rows"),
			"log_size":  newKByteFieldInfo("The size of file of Log"),
		},
		Tags: map[string]interface{}{
			"database_name": inputs.NewTagInfo("Name of the database"),
		},
	}
}

type DatabaseFilesMeasurement struct {
	MetricMeasurment
}

//nolint:lll
func (m *DatabaseFilesMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_database_files",
		Type: "metric",
		Fields: map[string]interface{}{
			"size": &inputs.FieldInfo{
				DataType: inputs.Int,
				Type:     inputs.Gauge,
				Unit:     inputs.SizeKB,
				Desc:     "Current size of the database file",
			},
		},
		Tags: map[string]interface{}{
			"database":      inputs.NewTagInfo("Database name"),
			"state":         inputs.NewTagInfo("Database file state: 0 = Online, 1 = Restoring, 2 = Recovering, 3 = Recovery_Pending, 4 = Suspect, 5 = Unknown, 6 = Offline, 7 = Defunct"),
			"physical_name": inputs.NewTagInfo("Operating-system file name"),
			"state_desc":    inputs.NewTagInfo("Description of the file state"),
			"file_id":       inputs.NewTagInfo("ID of the file within database"),
			"file_type":     inputs.NewTagInfo("File type: 0 = Rows, 1 = Log, 2 = File-Stream, 3 = Identified for informational purposes only, 4 = Full-text"),
		},
	}
}

type DatabaseBackupMeasurement struct {
	MetricMeasurment
}

//nolint:lll
func (m *DatabaseBackupMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "sqlserver_database_backup",
		Type: "metric",
		Fields: map[string]interface{}{
			"backup_count": &inputs.FieldInfo{
				DataType: inputs.Int,
				Type:     inputs.Gauge,
				Unit:     inputs.Count,
				Desc:     "The total count of successful backups made for a database",
			},
		},
		Tags: map[string]interface{}{
			"database": inputs.NewTagInfo("Database name"),
		},
	}
}
