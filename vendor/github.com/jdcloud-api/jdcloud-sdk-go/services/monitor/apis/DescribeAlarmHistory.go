// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeAlarmHistoryRequest struct {

    core.JDCloudRequest

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 产品线标识，同一个产品线下可能存在多个product，如(redis下有redis2.8cluster、redis4.0) (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 产品标识,默认返回该product下所有dimension的数据。eg:product=redis2.8cluster（redis2.8cluster产品下包含redis2.8-shard与redis2.8-proxy、redis2.8-instance多个维度)。 (Optional) */
    Product *string `json:"product"`

    /* 维度标识、指定该参数时，查询只返回该维度的数据。如redis2.8cluster下存在实例、分片等多个维度 (Optional) */
    Dimension *string `json:"dimension"`

    /* 根据region筛选对应region的资源的报警历史 (Optional) */
    Region *string `json:"region"`

    /* 正在报警, 取值为1 (Optional) */
    IsAlarming *int `json:"isAlarming"`

    /* 报警的状态,1为报警恢复、2为报警、4为报警恢复无数据 (Optional) */
    Status *int `json:"status"`

    /* 开始时间 (Optional) */
    StartTime *string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime *string `json:"endTime"`

    /* 规则类型,默认查询1， 1表示资源监控，6表示站点监控,7表示可用性监控 (Optional) */
    RuleType *int `json:"ruleType"`

    /* 规则名称模糊搜索 (Optional) */
    RuleName *string `json:"ruleName"`

    /* serviceCodes - 产品线servicecode，精确匹配，支持多个
resourceIds - 资源Id，精确匹配，支持多个（必须指定serviceCode才会在该serviceCode下根据resourceIds过滤，否则该参数不生效）
alarmIds - 规则Id，精确匹配，支持多个 (Optional) */
    Filters []monitor.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAlarmHistoryRequest(
) *DescribeAlarmHistoryRequest {

	return &DescribeAlarmHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/groupAlarmsHistory",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
	}
}

/*
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param serviceCode: 产品线标识，同一个产品线下可能存在多个product，如(redis下有redis2.8cluster、redis4.0) (Optional)
 * param product: 产品标识,默认返回该product下所有dimension的数据。eg:product=redis2.8cluster（redis2.8cluster产品下包含redis2.8-shard与redis2.8-proxy、redis2.8-instance多个维度)。 (Optional)
 * param dimension: 维度标识、指定该参数时，查询只返回该维度的数据。如redis2.8cluster下存在实例、分片等多个维度 (Optional)
 * param region: 根据region筛选对应region的资源的报警历史 (Optional)
 * param isAlarming: 正在报警, 取值为1 (Optional)
 * param status: 报警的状态,1为报警恢复、2为报警、4为报警恢复无数据 (Optional)
 * param startTime: 开始时间 (Optional)
 * param endTime: 结束时间 (Optional)
 * param ruleType: 规则类型,默认查询1， 1表示资源监控，6表示站点监控,7表示可用性监控 (Optional)
 * param ruleName: 规则名称模糊搜索 (Optional)
 * param filters: serviceCodes - 产品线servicecode，精确匹配，支持多个
resourceIds - 资源Id，精确匹配，支持多个（必须指定serviceCode才会在该serviceCode下根据resourceIds过滤，否则该参数不生效）
alarmIds - 规则Id，精确匹配，支持多个 (Optional)
 */
func NewDescribeAlarmHistoryRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    serviceCode *string,
    product *string,
    dimension *string,
    region *string,
    isAlarming *int,
    status *int,
    startTime *string,
    endTime *string,
    ruleType *int,
    ruleName *string,
    filters []monitor.Filter,
) *DescribeAlarmHistoryRequest {

    return &DescribeAlarmHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/groupAlarmsHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        ServiceCode: serviceCode,
        Product: product,
        Dimension: dimension,
        Region: region,
        IsAlarming: isAlarming,
        Status: status,
        StartTime: startTime,
        EndTime: endTime,
        RuleType: ruleType,
        RuleName: ruleName,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAlarmHistoryRequestWithoutParam() *DescribeAlarmHistoryRequest {

    return &DescribeAlarmHistoryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/groupAlarmsHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeAlarmHistoryRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeAlarmHistoryRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param serviceCode: 产品线标识，同一个产品线下可能存在多个product，如(redis下有redis2.8cluster、redis4.0)(Optional) */
func (r *DescribeAlarmHistoryRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param product: 产品标识,默认返回该product下所有dimension的数据。eg:product=redis2.8cluster（redis2.8cluster产品下包含redis2.8-shard与redis2.8-proxy、redis2.8-instance多个维度)。(Optional) */
func (r *DescribeAlarmHistoryRequest) SetProduct(product string) {
    r.Product = &product
}

/* param dimension: 维度标识、指定该参数时，查询只返回该维度的数据。如redis2.8cluster下存在实例、分片等多个维度(Optional) */
func (r *DescribeAlarmHistoryRequest) SetDimension(dimension string) {
    r.Dimension = &dimension
}

/* param region: 根据region筛选对应region的资源的报警历史(Optional) */
func (r *DescribeAlarmHistoryRequest) SetRegion(region string) {
    r.Region = &region
}

/* param isAlarming: 正在报警, 取值为1(Optional) */
func (r *DescribeAlarmHistoryRequest) SetIsAlarming(isAlarming int) {
    r.IsAlarming = &isAlarming
}

/* param status: 报警的状态,1为报警恢复、2为报警、4为报警恢复无数据(Optional) */
func (r *DescribeAlarmHistoryRequest) SetStatus(status int) {
    r.Status = &status
}

/* param startTime: 开始时间(Optional) */
func (r *DescribeAlarmHistoryRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 结束时间(Optional) */
func (r *DescribeAlarmHistoryRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param ruleType: 规则类型,默认查询1， 1表示资源监控，6表示站点监控,7表示可用性监控(Optional) */
func (r *DescribeAlarmHistoryRequest) SetRuleType(ruleType int) {
    r.RuleType = &ruleType
}

/* param ruleName: 规则名称模糊搜索(Optional) */
func (r *DescribeAlarmHistoryRequest) SetRuleName(ruleName string) {
    r.RuleName = &ruleName
}

/* param filters: serviceCodes - 产品线servicecode，精确匹配，支持多个
resourceIds - 资源Id，精确匹配，支持多个（必须指定serviceCode才会在该serviceCode下根据resourceIds过滤，否则该参数不生效）
alarmIds - 规则Id，精确匹配，支持多个(Optional) */
func (r *DescribeAlarmHistoryRequest) SetFilters(filters []monitor.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAlarmHistoryRequest) GetRegionId() string {
    return ""
}

type DescribeAlarmHistoryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAlarmHistoryResult `json:"result"`
}

type DescribeAlarmHistoryResult struct {
    AlarmHistoryList []monitor.DescribedAlarmHistory `json:"alarmHistoryList"`
    NumberPages int64 `json:"numberPages"`
    NumberRecords int64 `json:"numberRecords"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
}