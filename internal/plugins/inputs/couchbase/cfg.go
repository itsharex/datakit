// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package couchbase

import (
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/io/point"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/plugins/inputs"
)

// See also https://github.com/couchbase/couchbase-exporter

type NodeMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
}

func (m *NodeMeasurement) LineProto() (*point.Point, error) {
	return point.NewPoint(m.name, m.tags, m.fields, point.MOptElection())
}

// Info ...
// nolint:lll
func (m *NodeMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "cbnode",
		Fields: map[string]interface{}{
			"interestingstats_ops":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total operations per second (including `XDCR`) to this bucket. (measured from cmd_get + cmd_set + incr_misses + incr_hits + decr_misses + decr_hits + delete_misses + delete_hits + ep_num_ops_del_meta + ep_num_ops_get_meta + ep_num_ops_set_meta)."},
			"interestingstats_curr_items":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Current number of unique items in Couchbase."},
			"uptime":                                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Uptime."},
			"interestingstats_couch_docs_data_size":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes of active data in this bucket. (measured from couch_docs_data_size)."},
			"interestingstats_vb_active_number_non_resident": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Interestingstats vb active number non resident."},
			"memory_total":                                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Memory total."},
			"rebalance_start":                                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Rebalance start."},
			"rebalance_success":                              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Rebalance success."},
			"systemstats_cpu_utilization_rate":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of CPU in use across all available cores on this server."},
			"systemstats_swap_total":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes of total swap space available on this server."},
			"systemstats_swap_used":                          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes of swap space in use on this server."},
			"graceful_failover_start":                        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Graceful failover start."},
			"interestingstats_couch_views_actual_disk_size":  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes of active items in all the views for this bucket on disk (measured from couch_views_actual_disk_size)."},
			"memory_free":                                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Memory free."},
			"failover":                                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Failover."},
			"failover_complete":                              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Failover complete."},
			"healthy":                                        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Bool, Desc: "Is this node healthy."},
			"systemstats_mem_free":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes of memory not in use on this server."},
			"interestingstats_couch_docs_actual_disk_size":   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The size of all data service files on disk for this bucket, including the data itself, metadata, and temporary files. (measured from couch_docs_actual_disk_size)."},
			"interestingstats_ep_bg_fetched":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of reads per second from disk for this bucket. (measured from ep_bg_fetched)."},
			"memcached_memory_reserved":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Memcached memory reserved."},
			"failover_incomplete":                            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Failover incomplete."},
			"interestingstats_curr_items_tot":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Current number of items in Couchbase including replicas."},
			"interestingstats_cmd_get":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of reads (get operations) per second from this bucket. (measured from cmd_get)."},
			"interestingstats_get_hits":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of get operations per second for data that this bucket contains. (measured from get_hits)."},
			"graceful_failover_success":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Graceful failover success."},
			"systemstats_mem_total":                          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes of total memory available on this server."},
			"interestingstats_couch_spatial_disk_size":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Interestingstats couch spatial disk size."},
			"interestingstats_couch_spatial_data_size":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Interestingstats couch spatial data size."},
			"interestingstats_couch_views_data_size":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes of active data for all the views in this bucket. (measured from couch_views_data_size)."},
			"interestingstats_mem_used":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Total memory used in bytes. (as measured from mem_used)."},
			"rebalance_failure":                              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Rebalance failure."},
			"rebalance_stop":                                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Rebalance stop."},
			"interestingstats_vb_replica_curr_items":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items in replica vBuckets in this bucket. (measured from vb_replica_curr_items)."},
			"memcached_memory_allocated":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Memcached memory allocated."},
			"failover_node":                                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Failover node."},
			"cluster_membership":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Whether or not node is part of the CB cluster."},
			"graceful_failover_fail":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Graceful failover fail."},
		},
		Tags: map[string]interface{}{
			"host":     inputs.NewTagInfo("Host name."),
			"instance": inputs.NewTagInfo("Instance endpoint."),
			"cluster":  inputs.NewTagInfo("Cluster name."),
			"node":     inputs.NewTagInfo("Node ip."),
		},
	}
}

type BucketInfoMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
}

func (m *BucketInfoMeasurement) LineProto() (*point.Point, error) {
	return point.NewPoint(m.name, m.tags, m.fields, point.MOptElection())
}

// Info ...
// nolint:lll
func (m *BucketInfoMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "cbbucketinfo",
		Fields: map[string]interface{}{
			"basic_memused_bytes":      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Basic memory used."},
			"basic_opspersec":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Basic ops per second."},
			"basic_quota_user_percent": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Basic quota percent used."},
			"basic_dataused_bytes":     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Basic data used."},
			"basic_diskfetches":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Bool, Desc: "Basic disk fetches."},
			"basic_diskused_bytes":     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Basic disk used."},
			"basic_itemcount":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "BucketItemCount first tries to retrieve an accurate bucket count via N1QL, but falls back to the REST API if that cannot be done (when there's no index to count all items in a bucket."},
		},
		Tags: map[string]interface{}{
			"host":     inputs.NewTagInfo("Host name."),
			"instance": inputs.NewTagInfo("Instance endpoint."),
			"bucket":   inputs.NewTagInfo("Bucket name."),
			"cluster":  inputs.NewTagInfo("Cluster name."),
		},
	}
}

type TaskMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
}

func (m *TaskMeasurement) LineProto() (*point.Point, error) {
	return point.NewPoint(m.name, m.tags, m.fields, point.MOptElection())
}

// Info ...
// nolint:lll
func (m *TaskMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "cbtask",
		Fields: map[string]interface{}{
			"rebalance_progress":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Progress of a rebalance task."},
			"compacting_progress":              &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Progress of a bucket compaction task."},
			"xdcr_docs_checked":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of documents checked for changes."},
			"xdcr_docs_written":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of documents written to the destination cluster."},
			"xdcr_paused":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Is this replication paused."},
			"active_vbuckets_left":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of Active VBuckets remaining."},
			"node_rebalance_progress":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Progress of a rebalance task per node."},
			"cluster_logs_collection_progress": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Progress of a cluster logs collection task."},
			"xdcr_changes_left":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of updates still pending replication."},
			"xdcr_errors":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of errors."},
			"docs_total":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Docs total."},
			"docs_transferred":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Docs transferred."},
			"replica_vbuckets_left":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of Replica VBuckets remaining."},
		},
		Tags: map[string]interface{}{
			"host":     inputs.NewTagInfo("Host name."),
			"instance": inputs.NewTagInfo("Instance endpoint."),
			"bucket":   inputs.NewTagInfo("Bucket name."),
			"cluster":  inputs.NewTagInfo("Cluster name."),
			"node":     inputs.NewTagInfo("Node ip."),
			"source":   inputs.NewTagInfo("Source id."),
			"target":   inputs.NewTagInfo("Target id."),
		},
	}
}

type QueryMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
}

func (m *QueryMeasurement) LineProto() (*point.Point, error) {
	return point.NewPoint(m.name, m.tags, m.fields, point.MOptElection())
}

// Info ...
// nolint:lll
func (m *QueryMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "cbquery",
		Fields: map[string]interface{}{
			"errors":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of query errors."},
			"warnings":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of query warnings."},
			"avg_req_time":      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Average request time."},
			"invalid_requests":  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of invalid requests."},
			"request_time":      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Query request time."},
			"requests_5000ms":   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of requests that take longer than 5000 ms per second."},
			"result_count":      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Query result count."},
			"service_time":      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Query service time."},
			"avg_response_size": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Average response size."},
			"queued_requests":   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of queued requests."},
			"requests":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of query requests."},
			"requests_250ms":    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of requests that take longer than 250 ms per second."},
			"requests_500ms":    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of requests that take longer than 500 ms per second."},
			"avg_svc_time":      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Average service time."},
			"avg_result_count":  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Average result count."},
			"active_requests":   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Active number of requests."},
			"requests_1000ms":   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of requests that take longer than 1000 ms per second."},
			"result_size":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Query result size."},
			"selects":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of queries involving SELECT."},
		},
		Tags: map[string]interface{}{
			"host":     inputs.NewTagInfo("Host name."),
			"instance": inputs.NewTagInfo("Instance endpoint."),
			"cluster":  inputs.NewTagInfo("Cluster name."),
		},
	}
}

type IndexMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
}

func (m *IndexMeasurement) LineProto() (*point.Point, error) {
	return point.NewPoint(m.name, m.tags, m.fields, point.MOptElection())
}

// Info ...
// nolint:lll
func (m *IndexMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "cbindex",
		Fields: map[string]interface{}{
			"memory_quota":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Index Service memory quota."},
			"remaining_ram":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes of Index RAM quota still available on this server."},
			"frag_percent":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage fragmentation of the index."},
			"num_rows_returned":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total number of rows returned so far by the indexer."},
			"resident_percent":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of the data held in memory."},
			"num_docs_indexed":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of documents indexed by the indexer since last startup."},
			"items_count":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The number of items currently indexed."},
			"num_docs_pending_queued": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of documents pending to be indexed."},
			"avg_scan_latency":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationNS, Desc: "Average time to serve a scan request (nanoseconds)."},
			"cache_misses":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Accesses to this index data from disk."},
			"cache_hit_percent":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of memory accesses that were served from the managed cache."},
			"memory_used":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Index Service memory used."},
			"ram_percent":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of Index RAM quota in use across all indexes on this server."},
			"num_requests":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of requests served by the indexer since last startup."},
			"cache_hits":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Accesses to this index data from RAM."},
		},
		Tags: map[string]interface{}{
			"host":     inputs.NewTagInfo("Host name."),
			"instance": inputs.NewTagInfo("Instance endpoint."),
			"cluster":  inputs.NewTagInfo("Cluster name."),
			"keyspace": inputs.NewTagInfo("Key space name."),
		},
	}
}

type SearchMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
}

func (m *SearchMeasurement) LineProto() (*point.Point, error) {
	return point.NewPoint(m.name, m.tags, m.fields, point.MOptElection())
}

// Info ...
// nolint:lll
func (m *SearchMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "cbfts",
		Fields: map[string]interface{}{
			"num_bytes_used_ram":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Amount of RAM used by FTS on this server."},
			"total_queries_rejected_by_herder": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of fts queries rejected by the FTS throttler due to high memory consumption."},
			"curr_batches_blocked_by_herder":   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of DCP batches blocked by the FTS throttler due to high memory consumption."},
		},
		Tags: map[string]interface{}{
			"host":     inputs.NewTagInfo("Host name."),
			"instance": inputs.NewTagInfo("Instance endpoint."),
			"cluster":  inputs.NewTagInfo("Cluster name."),
		},
	}
}

type CbasMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
}

func (m *CbasMeasurement) LineProto() (*point.Point, error) {
	return point.NewPoint(m.name, m.tags, m.fields, point.MOptElection())
}

// Info ...
// nolint:lll
func (m *CbasMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "cbcbas",
		Fields: map[string]interface{}{
			"disk_used":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The total disk size used by Analytics."},
			"gc_count":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of JVM garbage collections for Analytics node."},
			"gc_time":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "The amount of time in milliseconds spent performing JVM garbage collections for Analytics node."},
			"heap_used":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of JVM heap used by Analytics on this server."},
			"io_reads":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number of disk bytes read on Analytics node per second."},
			"io_writes":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number of disk bytes written on Analytics node per second."},
			"system_load_avg": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "System load for Analytics node."},
			"thread_count":    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of threads for Analytics node."},
		},
		Tags: map[string]interface{}{
			"host":     inputs.NewTagInfo("Host name."),
			"instance": inputs.NewTagInfo("Instance endpoint."),
			"cluster":  inputs.NewTagInfo("Cluster name."),
		},
	}
}

type EventingMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
}

func (m *EventingMeasurement) LineProto() (*point.Point, error) {
	return point.NewPoint(m.name, m.tags, m.fields, point.MOptElection())
}

// Info ...
// nolint:lll
func (m *EventingMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "cbeventing",
		Fields: map[string]interface{}{
			"bucket_op_exception_count":      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Eventing bucket op exception count."},
			"test_checkpoint_failure_count":  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Test eventing bucket op exception count."},
			"test_on_delete_success":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Test on delete success."},
			"dcp_backlog":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Mutations yet to be processed by the function."},
			"failed_count":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Mutations for which the function execution failed."},
			"on_delete_failure":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number of disk bytes written on Analytics node per second."},
			"on_delete_success":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "System load for Analytics node."},
			"processed_count":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Mutations for which the function has finished processing."},
			"timeout_count":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Function execution timed-out while processing."},
			"test_failed_count":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Test failed count."},
			"test_on_update_failure":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Test on update failure."},
			"checkpoint_failure_count":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Checkpoint failure count."},
			"n1ql_op_exception_count":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Number of disk bytes read on Analytics node per second."},
			"test_dcp_backlog":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Test dcp backlog."},
			"test_on_delete_failure":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Test on delete failure."},
			"test_processed_count":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Test processed count."},
			"on_update_failure":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "On update failure."},
			"on_update_success":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "On update success."},
			"test_bucket_op_exception_count": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Test bucket op exception count."},
			"test_n1ql_op_exception_count":   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Test n1ql op exception count."},
			"test_on_update_success":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Test on update success."},
			"test_timeout_count":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Test timeout count."},
		},
		Tags: map[string]interface{}{
			"host":     inputs.NewTagInfo("Host name."),
			"instance": inputs.NewTagInfo("Instance endpoint."),
			"cluster":  inputs.NewTagInfo("Cluster name."),
		},
	}
}

type PerNodeBucketStatsMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
}

func (m *PerNodeBucketStatsMeasurement) LineProto() (*point.Point, error) {
	return point.NewPoint(m.name, m.tags, m.fields, point.MOptElection())
}

// Info ...
// nolint:lll
func (m *PerNodeBucketStatsMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "cbpernodebucket",
		Fields: map[string]interface{}{
			"ep_max_size":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The maximum amount of memory this bucket can use."},
			"ep_num_ops_del_ret_meta":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of delRetMeta operations per second for this bucket as the target for `XDCR`."},
			"ep_num_ops_get_meta":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of metadata read operations per second for this bucket as the target for `XDCR`."},
			"vb_pending_queue_drain":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of pending items per second being written to disk in this bucket and should be transient during rebalancing."},
			"ep_dcp_2i_items_sent":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of indexes items sent."},
			"ep_ops_update":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items updated on disk per second for this bucket."},
			"vb_active_itm_memory":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of active user data cached in RAM in this bucket."},
			"swap_total":                              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total amount of swap available."},
			"avg_bg_wait_seconds":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "avg_bg_wait_seconds."},
			"hit_ratio":                               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Rate, Desc: "Hit ratio."},
			"disk_write_queue":                        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items waiting to be written to disk in this bucket. (measured from ep_queue_size+ep_flusher_todo)."},
			"ep_dcp_xdcr_total_backlog_size":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp `XDCR` total backlog size."},
			"vb_total_queue_age":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Vb total queue age."},
			"mem_free":                                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Amount of Memory free."},
			"avg_disk_commit_time":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average disk commit time in seconds as from disk_update histogram of timings."},
			"disk_update_total":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Disk update total."},
			"ep_clock_cas_drift_threshold_exceeded":   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Ep clock cas drift threshold exceeded."},
			"ep_dcp_other_producer_count":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of other senders for this bucket."},
			"ep_num_ops_set_ret_meta":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of setRetMeta operations per second for this bucket as the target for `XDCR`."},
			"vb_avg_replica_queue_age":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average age in seconds of replica items in the replica item queue for this bucket."},
			"curr_items_tot":                          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total number of items in this bucket."},
			"ep_dcp_cbas_items_remaining":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items remaining to be sent to consumer in this bucket."},
			"ep_dcp_replica_total_bytes":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Number of bytes per second being sent for replication DCP connections for this bucket."},
			"ep_diskqueue_items":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total number of items waiting to be written to disk in this bucket."},
			"vb_active_queue_drain":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of active items per second being written to disk in this bucket."},
			"vb_active_queue_fill":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of active items per second being put on the active item disk queue in this bucket."},
			"vb_pending_queue_size":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of pending items waiting to be written to disk in this bucket and should be transient during rebalancing."},
			"incr_hits":                               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of increment hits."},
			"ep_resident_items_rate":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of all items cached in RAM in this bucket."},
			"ep_dcp_views_indexes_backoff":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp views indexes backoff."},
			"disk_commit_count":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Disk commit count."},
			"ep_dcp_cbas_backoff":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of backoffs per second for analytics DCP connections (measured from ep_dcp_cbas_backoff)."},
			"ep_dcp_other_items_remaining":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items remaining to be sent to consumer in this bucket (measured from ep_dcp_other_items_remaining)."},
			"ep_meta_data_memory":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Total amount of item metadata consuming RAM in this bucket."},
			"get_hits":                                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of get hits."},
			"vb_active_num":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of vBuckets in the active state for this bucket."},
			"vb_pending_num":                          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of vBuckets in the pending state for this bucket and should be transient during rebalancing."},
			"couch_docs_disk_size":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The size of all data files for this bucket, including the data itself, meta data and temporary files."},
			"cas_misses":                              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Compare and Swap misses."},
			"decr_misses":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Decrement misses."},
			"ep_dcp_2i_total_backlog_size":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp 2i total backlog size."},
			"ep_dcp_fts_count":                        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp fts count."},
			"ep_dcp_views_producer_count":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of views producers."},
			"cpu_local_ms":                            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "CPU local ms."},
			"vb_replica_ops_update":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items updated on replica vBucket per second for this bucket."},
			"disk_commit_total":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Disk commit total."},
			"ep_dcp_fts_backlog_size":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp fts backlog size."},
			"ep_dcp_fts_total_bytes":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp fts total bytes."},
			"ep_dcp_other_backoff":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of backoffs for other DCP connections."},
			"ep_dcp_xdcr_producer_count":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of `XDCR` senders for this bucket."},
			"ep_ops_create":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total number of new items being inserted into this bucket."},
			"vb_pending_ops_update":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items updated on pending vBucket per second for this bucket."},
			"vb_avg_total_queue_age":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average age in seconds of all items in the disk write queue for this bucket."},
			"vb_active_resident_items_ratio":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of active items cached in RAM in this bucket."},
			"couch_views_data_size":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The size of active data on for all the indexes in this bucket."},
			"ep_dcp_fts_backoff":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp fts backoff."},
			"ep_dcp_views_count":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of views DCP connections."},
			"ep_dcp_views_total_backlog_size":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp views total backlog size."},
			"vb_replica_num":                          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of vBuckets in the replica state for this bucket."},
			"couch_views_actual_disk_size":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The size of all active items in all the indexes for this bucket on disk."},
			"bytes_written":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes written."},
			"delete_hits":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of delete operations per second for this bucket."},
			"delete_misses":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of delete operations per second for data that this bucket does not contain. (measured from delete_misses)."},
			"ep_active_hlc_drift":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Ep active hlc drift."},
			"vb_pending_resident_items_ratio":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of items in pending state buckets cached in RAM in this bucket."},
			"cmd_set":                                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of writes (set operations) per second to this bucket."},
			"ep_dcp_cbas_producer_count":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of analytics senders for this bucket (measured from ep_dcp_cbas_producer_count)."},
			"ep_dcp_other_total_backlog_size":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp other total backlog size."},
			"ep_dcp_xdcr_count":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of internal `XDCR` DCP connections in this bucket."},
			"vb_pending_queue_fill":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of pending items per second being put on the pending item disk queue in this bucket and should be transient during rebalancing."},
			"vb_replica_num_non_resident":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Vb replica num non resident."},
			"vb_replica_queue_drain":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of replica items per second being written to disk in this bucket."},
			"ep_dcp_views_indexes_total_bytes":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp views indexes total bytes."},
			"ep_dcp_2i_backoff":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of backoffs for indexes DCP connections."},
			"ep_dcp_2i_total_bytes":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number of bytes per second being sent for indexes DCP connections."},
			"ep_dcp_replica_items_sent":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being sent for a producer for this bucket."},
			"ep_dcp_replica_total_backlog_size":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp replica total backlog size."},
			"ep_dcp_views_items_remaining":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of views items remaining to be sent."},
			"ep_replica_hlc_drift_count":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Ep replica hlc drift count."},
			"vb_replica_resident_items_ratio":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Percentage of active items cached in RAM in this bucket."},
			"cas_bad_val":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Compare and Swap bad values."},
			"curr_items":                              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items in active vBuckets in this bucket."},
			"ep_dcp_total_bytes":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp total bytes."},
			"ep_num_non_resident":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of non-resident items."},
			"swap_used":                               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of swap space in use on this server."},
			"vb_pending_ops_create":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "New items per second being instead into pending vBuckets in this bucket and should be transient during rebalancing."},
			"couch_views_disk_size":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Couch views disk size."},
			"ep_data_read_failed":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of disk read failures. (measured from ep_data_read_failed)."},
			"ep_dcp_fts_items_sent":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp fts items sent."},
			"ep_dcp_views_items_sent":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of views items sent."},
			"ep_num_ops_del_meta":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of delete operations per second for this bucket as the target for `XDCR`."},
			"ep_num_value_ejects":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total number of items per second being ejected to disk in this bucket."},
			"incr_misses":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of increment misses."},
			"vb_replica_meta_data_memory":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of replica item metadata consuming in RAM in this bucket."},
			"couch_docs_actual_disk_size":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The size of all data files for this bucket, including the data itself, meta data and temporary files."},
			"ep_data_write_failed":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of disk write failures. (measured from ep_data_write_failed)."},
			"ep_dcp_2i_count":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of indexes DCP connections."},
			"ep_dcp_cbas_count":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of internal analytics DCP connections in this bucket (measured from ep_dcp_cbas_count)."},
			"ep_dcp_xdcr_items_sent":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being sent for a producer for this bucket."},
			"evictions":                               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of evictions."},
			"vb_replica_curr_items":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items in replica vBuckets in this bucket."},
			"avg_replica_timestamp_drift":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average drift (in seconds) between mutation timestamps and the local time for replica vBuckets. (measured from ep_replica_hlc_drift and ep_replica_hlc_drift_count)."},
			"couch_spatial_disk_size":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Couch spatial disk size."},
			"couch_spatial_ops":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Couch spatial ops."},
			"ep_dcp_fts_items_remaining":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp fts items remaining."},
			"ep_dcp_other_count":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of other DCP connections in this bucket."},
			"ep_replica_ahead_exceptions":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of all items cached in RAM in this bucket."},
			"vb_replica_queue_age":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Sum of disk replica queue item age in milliseconds."},
			"xdc_ops":                                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total `XDCR` operations per second for this bucket."},
			"ep_dcp_views_indexes_count":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp views indexes count."},
			"ep_active_hlc_drift_count":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Ep active hlc drift count."},
			"ep_mem_high_wat":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "High water mark for auto-evictions."},
			"ops":                                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total amount of operations per second to this bucket."},
			"vb_active_ops_create":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "New items per second being inserted into active vBuckets in this bucket."},
			"vb_replica_ops_create":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "New items per second being inserted into replica vBuckets in this bucket."},
			"vb_replica_queue_size":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of replica items waiting to be written to disk in this bucket."},
			"mem_actual_free":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Amount of RAM available on this server."},
			"couch_total_disk_size":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The total size on disk of all data and view files for this bucket."},
			"ep_mem_low_wat":                          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Low water mark for auto-evictions."},
			"ep_overhead":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Extra memory used by transient data like persistence queues or checkpoints."},
			"vb_pending_meta_data_memory":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of pending item metadata consuming RAM in this bucket and should be transient during rebalancing."},
			"vb_replica_queue_fill":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of replica items per second being put on the replica item disk queue in this bucket."},
			"mem_actual_used":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Memory actual used."},
			"rest_requests":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Rate of http requests on port 8091."},
			"ep_cache_miss_rate":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Percentage of reads per second to this bucket from disk as opposed to RAM."},
			"bg_wait_count":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Bg wait count."},
			"decr_hits":                               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Decrement hits."},
			"ep_dcp_replica_items_remaining":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items remaining to be sent to consumer in this bucket."},
			"get_misses":                              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of get misses."},
			"vb_active_meta_data_memory":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of active item metadata consuming RAM in this bucket."},
			"vb_pending_itm_memory":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of pending user data cached in RAM in this bucket and should be transient during rebalancing."},
			"couch_docs_data_size":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The size of active data in this bucket."},
			"ep_dcp_cbas_items_sent":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being sent for a producer for this bucket."},
			"ep_dcp_fts_producer_count":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp fts producer count."},
			"ep_dcp_replica_backoff":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of backoffs for replication DCP connections."},
			"ep_dcp_replica_producer_count":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of replication senders for this bucket."},
			"ep_dcp_views_backoff":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of backoffs for views DCP connections."},
			"ep_flusher_todo":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items currently being written."},
			"vb_replica_itm_memory":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of replica user data cached in RAM in this bucket."},
			"cmd_get":                                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of reads (get operations) per second from this bucket."},
			"ep_active_ahead_exceptions":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total number of ahead exceptions (when timestamp drift between mutations and local time has exceeded 5000000 μs) per second for all active vBuckets."},
			"vb_active_ops_update":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items updated on active vBucket per second for this bucket."},
			"vb_active_queue_items":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Vb active queue items."},
			"ep_bg_fetched":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of reads per second from disk for this bucket."},
			"ep_dcp_2i_items_remaining":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of indexes items remaining to be sent."},
			"ep_dcp_replica_count":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of internal replication DCP connections in this bucket."},
			"ep_dcp_views_total_bytes":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number bytes per second being sent for views DCP connections."},
			"vb_active_eject":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being ejected to disk from active vBuckets in this bucket."},
			"avg_active_timestamp_drift":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average drift (in seconds) between mutation timestamps and the local time for active vBuckets. (measured from ep_active_hlc_drift and ep_active_hlc_drift_count)."},
			"ep_dcp_views_indexes_producer_count":     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp views indexes producer count."},
			"ep_dcp_views_indexes_total_backlog_size": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp views indexes total backlog size."},
			"ep_diskqueue_drain":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total number of items per second being written to disk in this bucket."},
			"ep_kv_size":                              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total amount of user data cached in RAM in this bucket."},
			"ep_dcp_views_indexes_items_remaining":    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp views indexes items remaining."},
			"vb_replica_eject":                        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being ejected to disk from replica vBuckets in this bucket."},
			"vb_avg_active_queue_age":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Sum of disk queue item age in milliseconds."},
			"mem_total":                               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Memory total."},
			"disk_update_count":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Disk update count."},
			"ep_queue_size":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items queued for storage."},
			"vb_pending_curr_items":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items in pending vBuckets in this bucket and should be transient during rebalancing."},
			"vb_pending_queue_age":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Sum of disk pending queue item age in milliseconds."},
			"ep_tmp_oom_errors":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of back-offs sent per second to client SDKs due to OOM situations from this bucket."},
			"avg_disk_update_time":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Average disk update time in microseconds as from disk_update histogram of timings."},
			"couch_docs_fragmentation":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "How much fragmented data there is to be compacted compared to real data for the data files in this bucket."},
			"couch_views_ops":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "All the view reads for all design documents including scatter gather."},
			"ep_dcp_2i_producers":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of indexes producers."},
			"ep_dcp_xdcr_backoff":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of backoffs for `XDCR` DCP connections."},
			"ep_item_commit_failed":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of times a transaction failed to commit due to storage errors."},
			"ep_oom_errors":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of times unrecoverable OOMs happened while processing operations."},
			"mem_used":                                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Amount of memory used."},
			"vb_active_queue_age":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Sum of disk queue item age in milliseconds."},
			"cpu_idle_ms":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "CPU idle milliseconds."},
			"cas_hits":                                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of operations with a CAS id per second for this bucket."},
			"ep_dcp_xdcr_items_remaining":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items remaining to be sent to consumer in this bucket."},
			"hibernated_requests":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of streaming requests on port 8091 now idle."},
			"couch_views_fragmentation":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "How much fragmented data there is to be compacted compared to real data for the view index files in this bucket."},
			"ep_dcp_views_indexes_items_sent":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp views indexes items sent."},
			"ep_dcp_cbas_total_backlog_size":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ep dcp `cbas` total backlog size."},
			"ep_dcp_other_total_bytes":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number of bytes per second being sent for other DCP connections for this bucket."},
			"ep_replica_hlc_drift":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "The sum of the total Absolute Drift, which is the accumulated drift observed by the vBucket. Drift is always accumulated as an absolute value."},
			"vb_active_num_non_resident":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of non resident vBuckets in the active state for this bucket."},
			"bg_wait_total":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Bg wait total."},
			"curr_connections":                        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of connections to this server including connections from external client SDKs, proxies, DCP requests and internal statistic gathering."},
			"ep_dcp_other_items_sent":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being sent for a producer for this bucket (measured from ep_dcp_other_items_sent)."},
			"ep_num_ops_set_meta":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of set operations per second for this bucket as the target for `XDCR`."},
			"ep_vb_total":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total number of vBuckets for this bucket."},
			"hibernated_waked":                        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Rate of streaming request wakeups on port 8091."},
			"mem_used_sys":                            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Memory used sys."},
			"couch_spatial_data_size":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Couch spatial data size."},
			"bytes_read":                              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes Read."},
			"ep_diskqueue_fill":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total number of items per second being put on the disk queue in this bucket."},
			"vb_pending_eject":                        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being ejected to disk from pending vBuckets in this bucket and should be transient during rebalancing."},
			"cpu_utilization_rate":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of CPU in use across all available cores on this server."},
			"ep_dcp_xdcr_total_bytes":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number of bytes per second being sent for `XDCR` DCP connections for this bucket."},
			"misses":                                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of misses."},
			"vb_active_queue_size":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of active items waiting to be written to disk in this bucket."},
			"vb_pending_num_non_resident":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of non resident vBuckets in the pending state for this bucket."},
			"vb_avg_pending_queue_age":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average age in seconds of pending items in the pending item queue for this bucket and should be transient during rebalancing."},
		},
		Tags: map[string]interface{}{
			"host":     inputs.NewTagInfo("Host name."),
			"instance": inputs.NewTagInfo("Instance endpoint."),
			"bucket":   inputs.NewTagInfo("Bucket name."),
			"cluster":  inputs.NewTagInfo("Cluster name."),
			"node":     inputs.NewTagInfo("Node ip."),
		},
	}
}

type BucketStatsMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
}

func (m *BucketStatsMeasurement) LineProto() (*point.Point, error) {
	return point.NewPoint(m.name, m.tags, m.fields, point.MOptElection())
}

// Info ...
// nolint:lll
func (m *BucketStatsMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: "cbbucketstat",
		Fields: map[string]interface{}{
			"ep_resident_items_rate":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Percentage of all items cached in RAM in this bucket."},
			"get_misses":                            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of get misses."},
			"vbuckets_avg_pending_queue_age":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average age in seconds of pending items in the pending item queue for this bucket and should be transient during rebalancing."},
			"vbuckets_pending_meta_data_memory":     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of pending item metadata consuming RAM in this bucket and should be transient during rebalancing."},
			"vbuckets_replica_queue_size":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of replica items waiting to be written to disk in this bucket."},
			"ep_dcp_2i_producers":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of indexes producers."},
			"ep_bg_fetched":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of reads per second from disk for this bucket."},
			"ep_queue_size":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items queued for storage."},
			"vbuckets_active_queue_drain":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of active items per second being written to disk in this bucket."},
			"cmd_set":                               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of writes (set operations) per second to this bucket."},
			"ep_dcp_2i_items_sent":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of indexes items sent."},
			"ep_dcp_other_items_sent":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being sent for a producer for this bucket."},
			"vbuckets_pending_curr_items":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items in pending vBuckets in this bucket and should be transient during rebalancing."},
			"couch_docs_actual_disk_size":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The size of all data files for this bucket, including the data itself, meta data and temporary files."},
			"ep_dcp_replica_total_bytes":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number of bytes per second being sent for replication DCP connections for this bucket."},
			"vbuckets_replica_eject":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being ejected to disk from replica vBuckets in this bucket."},
			"decr_hits":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Decrement hits."},
			"ep_active_hlc_drift":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Ep active hlc drift."},
			"ep_dcp_views_total_bytes":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number bytes per second being sent for views DCP connections."},
			"swap_bytes":                            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Total amount of swap available."},
			"vbuckets_total_queue_age":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Sum of disk queue item age in milliseconds."},
			"delete_misses":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Delete misses."},
			"vbuckets_active_num_non_resident":      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of non resident vBuckets in the active state for this bucket."},
			"vbuckets_pending_num":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of vBuckets in the pending state for this bucket and should be transient during rebalancing."},
			"couch_views_actual_disk_size":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The size of all active items in all the indexes for this bucket on disk."},
			"disk_write_queue":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items waiting to be written to disk in this bucket."},
			"ep_num_ops_del_meta":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of delete operations per second for this bucket as the target for `XDCR`."},
			"ep_replica_hlc_drift":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "The sum of the total Absolute Drift, which is the accumulated drift observed by the vBucket. Drift is always accumulated as an absolute value."},
			"mem_free_bytes":                        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Amount of Memory free."},
			"vbuckets_active_ops_create":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "New items per second being inserted into active vBuckets in this bucket."},
			"cas_misses":                            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Compare and Swap misses."},
			"couch_docs_data_size":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The size of active data in this bucket."},
			"ep_active_ahead_exceptions":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total number of ahead exceptions for  all active vBuckets."},
			"ep_dcp_replicas":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of internal replication DCP connections in this bucket."},
			"ep_num_non_resident":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of non-resident items."},
			"couch_total_disk_size":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The total size on disk of all data and view files for this bucket."},
			"ep_dcp_2i_total_backlog_size":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp 2i total backlog size."},
			"ep_dcp_other_backoff":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of backoffs for other DCP connections."},
			"ep_dcp_views_producers":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of views producers."},
			"ep_meta_data_memory":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Total amount of item metadata consuming RAM in this bucket."},
			"ep_num_value_ejects":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total number of items per second being ejected to disk in this bucket."},
			"evictions":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of evictions."},
			"delete_hits":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of delete operations per second for this bucket."},
			"ep_flusher_todo":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items currently being written."},
			"incr_misses":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of increment misses."},
			"vbuckets_replica_ops_create":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "New items per second being inserted into replica vBuckets in this bucket."},
			"ep_cache_miss_rate":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of reads per second to this bucket from disk as opposed to RAM."},
			"couch_views_data_size":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The size of active data on for all the indexes in this bucket."},
			"ep_dcp_other_items_remaining":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items remaining to be sent to consumer in this bucket."},
			"ep_dcp_views_total_backlog_size":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp views total backlog size."},
			"ep_vbuckets":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total number of vBuckets for this bucket."},
			"vbuckets_avg_total_queue_age":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average age in seconds of all items in the disk write queue for this bucket."},
			"xdc_ops":                               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total `XDCR` operations per second for this bucket."},
			"couch_views_fragmentation":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "How much fragmented data there is to be compacted compared to real data for the view index files in this bucket."},
			"ep_dcp_other_producers":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of other senders for this bucket."},
			"hibernated_waked":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Rate of streaming request wakeups on port 8091."},
			"incr_hits":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of increment hits."},
			"vbuckets_active_resident_items_ratio":  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of active items cached in RAM in this bucket."},
			"vbuckets_replica_curr_items":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items in replica vBuckets in this bucket."},
			"vbuckets_replica_num":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of vBuckets in the replica state for this bucket."},
			"avg_active_timestamp_drift":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average drift (in seconds) per mutation on active vBuckets."},
			"ep_num_ops_get_meta":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of metadata read operations per second for this bucket as the target for `XDCR`."},
			"couch_views_ops":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "All the view reads for all design documents including scatter gather."},
			"ops":                                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total amount of operations per second to this bucket."},
			"vbuckets_pending_queue_age":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Sum of disk pending queue item age in milliseconds."},
			"ep_dcp_others":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of other DCP connections in this bucket."},
			"vbuckets_active_meta_data_memory":      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Amount of active item metadata consuming RAM in this bucket."},
			"disk_updates":                          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Disk updates."},
			"rest_requests":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Rate of http requests on port 8091."},
			"vbuckets_active_itm_memory":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of active user data cached in RAM in this bucket."},
			"vbuckets_replica_meta_data_memory":     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Amount of replica item metadata consuming in RAM in this bucket."},
			"ep_dcp_xdcr_connections":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of internal `XDCR` DCP connections in this bucket."},
			"ep_dcp_2i_items_remaining":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of indexes items remaining to be sent."},
			"ep_dcp_other_total_bytes":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number of bytes per second being sent for other DCP connections for this bucket."},
			"ep_item_commit_failed":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of times a transaction failed to commit due to storage errors."},
			"ep_num_ops_set_ret_meta":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of setRetMeta operations per second for this bucket as the target for `XDCR`."},
			"get_hits":                              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of get hits."},
			"mem_actual_free":                       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Amount of RAM available on this server."},
			"cas_badval":                            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Compare and Swap bad values."},
			"vbuckets_pending_eject":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being ejected to disk from pending vBuckets in this bucket and should be transient during rebalancing."},
			"written_bytes":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes written."},
			"curr_items":                            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items in active vBuckets in this bucket."},
			"ep_dcp_2i_backoff":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of backoffs for indexes DCP connections."},
			"ep_dcp_replica_backoff":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of backoffs for replication DCP connections."},
			"ep_dcp_xdcr_items_sent":                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being sent for a producer for this bucket."},
			"ep_ops_update":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items updated on disk per second for this bucket."},
			"ep_tmp_oom_errors":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of back-offs sent per second to client SDKs due to OOM situations from this bucket."},
			"mem_used_bytes":                        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Amount of memory used."},
			"avg_disk_update_time":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationUS, Desc: "Average disk update time in microseconds as from disk_update histogram of timings."},
			"vbuckets_replica_ops_update":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items updated on replica vBucket per second for this bucket."},
			"vbuckets_pending_itm_memory":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of pending user data cached in RAM in this bucket and should be transient during rebalancing."},
			"cpu_idle_ms":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "CPU idle milliseconds."},
			"ep_num_ops_set_meta":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of set operations per second for this bucket as the target for `XDCR`."},
			"vbuckets_pending_ops_update":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items updated on pending vBucket per second for this bucket."},
			"vbuckets_replica_queue_fill":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of replica items per second being put on the replica item disk queue in this bucket."},
			"couch_docs_fragmentation":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "How much fragmented data there is to be compacted compared to real data for the data files in this bucket."},
			"ep_dcp_xdcr_total_bytes":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number of bytes per second being sent for `XDCR` DCP connections for this bucket."},
			"ep_dcp_2i_total_bytes":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.BytesPerSec, Desc: "Number bytes per second being sent for indexes DCP connections."},
			"ep_dcp_replica_producers":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of replication senders for this bucket."},
			"ep_dcp_xdcr_producers":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of `XDCR` senders for this bucket."},
			"ep_ops_create":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total number of new items being inserted into this bucket."},
			"vbuckets_pending_num_non_resident":     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of non resident vBuckets in the pending state for this bucket."},
			"ep_dcp_other_total_backlog_size":       &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp other total backlog size."},
			"mem_used_sys_bytes":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "System memory in use."},
			"vbuckets_pending_ops_create":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "New items per second being instead into pending vBuckets in this bucket and should be transient during rebalancing."},
			"vbuckets_replica_num_non_resident":     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "vb replica num non resident."},
			"ep_mem_high_wat_bytes":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "High water mark for auto-evictions."},
			"ep_dcp_replica_items_sent":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being sent for a producer for this bucket."},
			"ep_dcp_replica_total_backlog_size":     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp replica total backlog size."},
			"vbuckets_active_ops_update":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items updated on active vBucket per second for this bucket."},
			"vbuckets_active_queue_size":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of active items waiting to be written to disk in this bucket."},
			"vbuckets_pending_queue_drain":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of pending items per second being written to disk in this bucket and should be transient during rebalancing."},
			"vbuckets_replica_queue_drain":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of replica items per second being written to disk in this bucket."},
			"couch_docs_disk_size":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The size of all data files for this bucket, including the data itself, meta data and temporary files."},
			"ep_replica_ahead_exceptions":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total number of ahead exceptions (when timestamp drift between mutations and local time has exceeded 5000000 μs) per second for all replica vBuckets."},
			"mem_bytes":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Total amount of memory available."},
			"ep_dcp_views_items_remaining":          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of views items remaining to be sent."},
			"cpu_local_ms":                          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "CPU local ms."},
			"curr_connections":                      &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of connections to this server including connections from external client SDKs, proxies, DCP requests and internal statistic gathering."},
			"decr_misses":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Decrement misses."},
			"ep_diskqueue_items":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total number of items waiting to be written to disk in this bucket."},
			"avg_bg_wait_seconds":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average background fetch time in seconds."},
			"ep_dcp_views_connections":              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of views DCP connections."},
			"ep_diskqueue_drain":                    &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total number of items per second being written to disk in this bucket."},
			"ep_diskqueue_fill":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Total number of items per second being put on the disk queue in this bucket."},
			"ep_num_ops_del_ret_meta":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of delRetMeta operations per second for this bucket as the target for `XDCR`."},
			"ep_overhead":                           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Extra memory used by transient data like persistence queues or checkpoints."},
			"vbuckets_pending_resident_items_ratio": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of items in pending state vb cached in RAM in this bucket."},
			"ep_clock_cas_drift_threshold_exceeded": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Ep clock cas drift threshold exceeded."},
			"avg_disk_commit_time":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average disk commit time in seconds as from disk_update histogram of timings."},
			"cmd_get":                               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of reads (get operations) per second from this bucket."},
			"ep_dcp_xdcr_items_remaining":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items remaining to be sent to consumer in this bucket."},
			"ep_dcp_xdcr_total_backlog_size":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Ep dcp `XDCR` total backlog size."},
			"hibernated_requests":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of streaming requests on port 8091 now idle."},
			"hit_ratio":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Hit ratio."},
			"misses":                                &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of misses."},
			"avg_replica_timestamp_drift":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average drift (in seconds) per mutation on replica vBuckets."},
			"vbuckets_replica_queue_age":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Sum of disk replica queue item age in milliseconds."},
			"vbuckets_avg_active_queue_age":         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average age in seconds of active items in the active item queue for this bucket."},
			"ep_dcp_views_items_sent":               &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of views items sent."},
			"ep_dcp_xdcr_backoff":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of backoffs for `XDCR` DCP connections."},
			"ep_oom_errors":                         &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of times unrecoverable OOMs happened while processing operations."},
			"vbuckets_active_queue_fill":            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of active items per second being put on the active item disk queue in this bucket."},
			"disk_commits":                          &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Disk commits."},
			"curr_items_tot":                        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total number of items in this bucket."},
			"ep_dcp_2i_connections":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of indexes DCP connections."},
			"ep_dcp_views_backoffs":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of backoffs for views DCP connections."},
			"ep_max_size_bytes":                     &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The maximum amount of memory this bucket can use."},
			"mem_actual_used_bytes":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Memory actually used in bytes."},
			"vbuckets_active_num":                   &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of vBuckets in the active state for this bucket."},
			"vbuckets_pending_queue_fill":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of pending items per second being put on the pending item disk queue in this bucket and should be transient during rebalancing."},
			"read_bytes":                            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Bytes read."},
			"vbuckets_pending_queue_size":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of pending items waiting to be written to disk in this bucket and should be transient during rebalancing."},
			"cpu_utilization_rate":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of CPU in use across all available cores on this server."},
			"ep_kv_size":                            &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Total amount of user data cached in RAM in this bucket."},
			"ep_mem_low_wat_bytes":                  &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Low water mark for auto-evictions."},
			"vbuckets_replica_resident_items_ratio": &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.Percent, Desc: "Percentage of replica items cached in RAM in this bucket."},
			"cas_hits":                              &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of operations with a CAS id per second for this bucket."},
			"swap_used":                             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "Amount of swap space in use on this server."},
			"vbuckets_active_eject":                 &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.RequestsPerSec, Desc: "Number of items per second being ejected to disk from active vBuckets in this bucket."},
			"vbuckets_active_queue_age":             &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "Sum of disk queue item age in milliseconds."},
			"vbuckets_avg_replica_queue_age":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.DurationSecond, Desc: "Average age in seconds of replica items in the replica item queue for this bucket."},
			"vbuckets_replica_itm_memory":           &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Amount of replica user data cached in RAM in this bucket."},
			"ep_dcp_replica_items_remaining":        &inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Number of items remaining to be sent to consumer in this bucket."},
		},
		Tags: map[string]interface{}{
			"host":     inputs.NewTagInfo("Host name."),
			"instance": inputs.NewTagInfo("Instance endpoint."),
			"bucket":   inputs.NewTagInfo("Bucket name."),
			"cluster":  inputs.NewTagInfo("Cluster name."),
		},
	}
}
