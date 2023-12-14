// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package profile

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/GuanceCloud/cliutils/point"
	"github.com/google/uuid"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/datakit"
	dkio "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/io"
)

const (
	eventJSONFile  = "event"
	profileTagsKey = "tags[]"
)

const (
	// Deprecated: use Collapsed instead.
	RawFlameGraph Format = "rawflamegraph" // flamegraph collapse
	Collapsed     Format = "collapse"      // flamegraph collapse format see https://github.com/brendangregg/FlameGraph/blob/master/stackcollapse.pl
	JFR           Format = "jfr"           // see https://github.com/openjdk/jmc#core-api-example
	PPROF         Format = "pprof"         // see https://github.com/google/pprof/blob/main/proto/profile.proto
)

const (
	Unknown       Profiler = "unknown"
	DDtrace       Profiler = "ddtrace"
	AsyncProfiler Profiler = "async-profiler"
	PySpy         Profiler = "py-spy"
	Pyroscope     Profiler = "pyroscope"
)

type Profiler string

type Format string

const (
	Java    Language = "java"
	Python  Language = "python"
	Golang  Language = "golang"
	Ruby    Language = "ruby"
	NodeJS  Language = "nodejs"
	PHP     Language = "php"
	DotNet  Language = "dotnet"
	UnKnown Language = "unknown"
)

type Language string

func (l Language) String() string {
	return string(l)
}

// line protocol tags.
const (
	TagEndPoint    = "endpoint"
	TagService     = "service"
	TagEnv         = "env"
	TagVersion     = "version"
	TagHost        = "host"
	TagLanguage    = "language"
	TagRuntime     = "runtime"
	TagOs          = "runtime_os"
	TagRuntimeArch = "runtime_arch"
)

// line protocol fields.
const (
	FieldProfileID  = "profile_id"
	FieldLibraryVer = "library_ver"
	FieldDatakitVer = "datakit_ver"
	FieldStart      = "start"
	FieldEnd        = "end"
	FieldDuration   = "duration"
	FieldFormat     = "format"
	FieldPid        = "pid"
	FieldRuntimeID  = "runtime_id"
	FieldFileSize   = "__file_size"
)

var langMaps = map[string]Language{
	"java":    Java,
	"jvm":     Java,
	"python":  Python,
	"ruby":    Ruby,
	"node.js": NodeJS,
	"nodejs":  NodeJS,
	"php":     PHP,
	"dotnet":  DotNet,
	"c#":      DotNet,
	"csharp":  DotNet,
	"golang":  Golang,
	"node":    NodeJS,
	"go":      Golang,
}

// Tags refer to parsed formValue["tags[]"].
type Tags map[string]string

type rfc3339Time time.Time

func (r rfc3339Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(r).Format(time.RFC3339Nano) + `"`), nil
}

type Metadata struct {
	Format       Format      `json:"format"`
	Profiler     Profiler    `json:"profiler"`
	Attachments  []string    `json:"attachments"`
	Language     Language    `json:"language"`
	TagsProfiler string      `json:"tags_profiler"`
	Start        rfc3339Time `json:"start"`
	End          rfc3339Time `json:"end"`
}

func newTags(originTags []string) Tags {
	pt := make(Tags)
	for _, tag := range originTags {
		// 有":"， 用:切割成键值对
		if strings.Index(tag, ":") > 0 {
			pairs := strings.SplitN(tag, ":", 2)
			pt[pairs[0]] = pairs[1]
		} else {
			// 没有":" 整个值做key, value为空
			pt[tag] = ""
		}
	}
	return pt
}

func (pt Tags) Get(name string, defVal ...string) string {
	if tag, ok := pt[name]; ok {
		return tag
	}
	if len(defVal) > 0 {
		return defVal[0]
	}
	return ""
}

func ResolveLanguage(runtimes []string) Language {
	for _, r := range runtimes {
		r = strings.ToLower(r)
		for name, lang := range langMaps {
			if strings.Contains(r, name) {
				return lang
			}
		}
	}
	return UnKnown
}

func randomProfileID() string {
	var (
		id  uuid.UUID
		err error
	)
	// generate google uuid
	for i := 0; i < 3; i++ {
		if id, err = uuid.NewRandom(); err == nil {
			return id.String()
		}
	}
	log.Error("call uuid.NewRandom() fail, use our uuid instead: ", err)
	return OurUUID()
}

func OurUUID() string {
	var buf [12]byte
	rand.Read(buf[:]) //nolint:gosec,errcheck
	random := hex.EncodeToString(buf[:8])

	nanos := strconv.FormatInt(time.Now().UnixNano(), 16)
	if len(nanos) > 14 {
		nanos = nanos[len(nanos)-14:]
	}
	host, err := os.Hostname()
	if err == nil && len(host) > 0 {
		host = hex.EncodeToString([]byte(host))
		if len(host) > 8 {
			host = host[len(host)-8:]
		}
	} else {
		host = hex.EncodeToString(buf[8:])
	}
	return fmt.Sprintf("%s-%s-%s-%s-%s", random[:8], random[8:12], random[12:], host, nanos)
}

func resolveStartTime(formValue map[string][]string) (time.Time, error) {
	return resolveTime(formValue, []string{"recording-start", "start"})
}

func resolveEndTime(formValue map[string][]string) (time.Time, error) {
	return resolveTime(formValue, []string{"recording-end", "end"})
}

func resolveTime(formValue map[string][]string, formFields []string) (time.Time, error) {
	var tm time.Time

	if len(formFields) == 0 {
		return tm, fmt.Errorf("form time fields is empty")
	}

	var err error
	for _, field := range formFields {
		if timeVal := formValue[field]; len(timeVal) > 0 {
			tVal := timeVal[0]
			if strings.Contains(tVal, ".") {
				tm, err = time.Parse(time.RFC3339Nano, tVal)
			} else {
				tm, err = time.Parse(time.RFC3339, tVal)
			}
			if err == nil {
				return tm, nil
			}
		}
	}
	if err != nil {
		return tm, err
	}

	return tm, errors.New("there is not proper form time field")
}

func getForm(field string, formValues map[string][]string) string {
	if val := formValues[field]; len(val) > 0 {
		return val[0]
	}
	return ""
}

func resolveLang(formValue map[string][]string, pt Tags) Language {
	var runtimes []string

	if v := pt.Get("language"); v != "" {
		runtimes = append(runtimes, v)
	}

	if v := pt.Get("runtime"); v != "" {
		runtimes = append(runtimes, v)
	}

	formKeys := []string{
		"runtime",
		"language",
		"family",
	}

	for _, field := range formKeys {
		if v := getForm(field, formValue); v != "" {
			runtimes = append(runtimes, v)
		}
	}

	return ResolveLanguage(runtimes)
}

func interface2String(i interface{}) (string, error) {
	switch baseV := i.(type) {
	case string:
		return baseV, nil
	case float64:
		return strconv.FormatFloat(baseV, 'g', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(baseV), 'g', -1, 64), nil
	case int64:
		return strconv.FormatInt(baseV, 10), nil
	case uint64:
		return strconv.FormatUint(baseV, 10), nil
	case int:
		return strconv.Itoa(baseV), nil
	case uint:
		return strconv.FormatUint(uint64(baseV), 10), nil
	case bool:
		if baseV {
			return "true", nil
		} else {
			return "false", nil
		}
	}
	return "", fmt.Errorf("not suppoerted interface type: %T", i)
}

func json2StringMap(m map[string]interface{}) map[string][]string {
	formatted := make(map[string][]string, len(m))

	for k, v := range m {
		switch baseV := v.(type) {
		case []interface{}:
			for _, elem := range baseV {
				if elemStr, err := interface2String(elem); err == nil {
					formatted[k] = append(formatted[k], elemStr)
				}
			}
		default:
			if vStr, err := interface2String(v); err == nil {
				formatted[k] = append(formatted[k], vStr)
			}
		}
	}
	return formatted
}

type resolvedMetadata struct {
	formValue map[string][]string
	tags      Tags
}

func parseMetadata(req *http.Request) (*resolvedMetadata, int64, error) {
	filesize := int64(0)
	for _, files := range req.MultipartForm.File {
		for _, f := range files {
			filesize += f.Size
		}
	}

	if req.MultipartForm.Value != nil {
		if _, ok := req.MultipartForm.Value[profileTagsKey]; ok {
			return &resolvedMetadata{
				formValue: req.MultipartForm.Value,
				tags:      newTags(req.MultipartForm.Value[profileTagsKey]),
			}, filesize, nil
		}
	}
	if eventFiles, ok := req.MultipartForm.File[eventJSONFile]; ok {
		if len(eventFiles) == 1 {
			fp, err := eventFiles[0].Open()
			if err != nil {
				return nil, filesize, fmt.Errorf("read event json file fail: %w", err)
			}
			defer func() {
				_ = fp.Close()
			}()

			var events map[string]interface{}
			decoder := json.NewDecoder(fp)
			if err := decoder.Decode(&events); err != nil {
				return nil, filesize, fmt.Errorf("resolve the event file fail: %w", err)
			}
			metadata := json2StringMap(events)
			if len(metadata["tags_profiler"]) == 1 {
				metadata[profileTagsKey] = strings.Split(metadata["tags_profiler"][0], ",")
				delete(metadata, "tags_profiler")
			}
			if _, ok := metadata[profileTagsKey]; !ok {
				return nil, filesize, fmt.Errorf("the profiling data format not supported, tags[] field missing")
			}
			return &resolvedMetadata{
				formValue: metadata,
				tags:      newTags(metadata[profileTagsKey]),
			}, filesize, nil
		}
	}
	return nil, filesize, fmt.Errorf("the profiling data format not supported, check your datadog trace library version")
}

func mkPoint(req *http.Request, metadata *resolvedMetadata, fileSize int64) (string, int64, *point.Point, error) {
	profileStart, err := resolveStartTime(metadata.formValue)
	if err != nil {
		return "", 0, nil, fmt.Errorf("can not resolve profile start time: %w", err)
	}

	profileEnd, err := resolveEndTime(metadata.formValue)
	if err != nil {
		return "", 0, nil, fmt.Errorf("can not resolve profile end time: %w", err)
	}

	startMicroSeconds, endMicroSeconds := profileStart.UnixMicro(), profileEnd.UnixMicro()

	tags := metadata.tags

	runtime := getForm("runtime", metadata.formValue)
	if runtime == "" {
		runtime = tags.Get("runtime")
	}

	serviceName := tags.Get("service")
	if serviceName == "" {
		serviceName = "unnamed-service"
	}

	var kvs point.KVs
	kvs = kvs.AddTag(TagHost, tags.Get("host"))
	kvs = kvs.AddTag(TagEndPoint, req.URL.Path)
	kvs = kvs.AddTag(TagService, serviceName)
	kvs = kvs.AddTag(TagOs, tags.Get("runtime_os"))
	kvs = kvs.AddTag(TagRuntimeArch, tags.Get("runtime_arch"))
	kvs = kvs.AddTag(TagEnv, tags.Get("env"))
	kvs = kvs.AddTag(TagVersion, tags.Get("version"))
	kvs = kvs.AddTag(TagLanguage, resolveLang(metadata.formValue, tags).String())
	kvs = kvs.AddTag(TagRuntime, runtime)

	profileID := randomProfileID()

	kvs = kvs.Add(FieldProfileID, profileID, false, false)
	kvs = kvs.Add(FieldRuntimeID, tags.Get("runtime-id"), false, false)
	kvs = kvs.Add(FieldPid, tags.Get("pid"), false, false)
	kvs = kvs.Add(FieldFormat, getForm("format", metadata.formValue), false, false)
	kvs = kvs.Add(FieldLibraryVer, tags.Get("profiler_version"), false, false)
	kvs = kvs.Add(FieldDatakitVer, datakit.Version, false, false)
	kvs = kvs.Add(FieldStart, startMicroSeconds, false, false)
	kvs = kvs.Add(FieldEnd, endMicroSeconds, false, false)
	kvs = kvs.Add(FieldDuration, endMicroSeconds-startMicroSeconds, false, false) // unit: microsecond
	kvs = kvs.Add(FieldFileSize, fileSize, false, false)

	// user custom tags
	for tName, tVal := range tags {
		// line protocol field name must not contain "."
		tName = strings.ReplaceAll(tName, ".", "_")
		kvs = kvs.Add(tName, tVal, false, false)
	}

	opts := point.CommonLoggingOptions()
	opts = append(opts, point.WithTime(profileEnd))
	pt := point.NewPointV2(inputName, kvs, opts...)
	if err != nil {
		return "", 0, nil, fmt.Errorf("build profile point fail: %w", err)
	}

	return profileID, profileStart.UnixNano(), pt, nil
}

func sendToIO(pt *point.Point) error {
	lang, _ := pt.Get(TagLanguage).(string)

	if err := iptGlobal.feeder.Feed(inputName+"/"+lang,
		point.Profiling,
		[]*point.Point{pt},
		&dkio.Option{CollectCost: time.Since(pt.Time())}); err != nil {
		return fmt.Errorf("unable to feed profiling point: %w", err)
	}
	return nil
}
