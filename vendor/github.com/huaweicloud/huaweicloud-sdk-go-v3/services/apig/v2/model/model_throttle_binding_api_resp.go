/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"

	"strings"
)

type ThrottleBindingApiResp struct {
	// API的认证方式
	AuthType *string `json:"auth_type,omitempty"`
	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty"`
	// API的发布记录编号
	PublishId *string `json:"publish_id,omitempty"`
	// 与流控策略的绑定关系编号
	ThrottleApplyId *string `json:"throttle_apply_id,omitempty"`
	// 已绑定的流控策略的绑定时间
	ApplyTime *sdktime.SdkTime `json:"apply_time,omitempty"`
	// API描述
	Remark *string `json:"remark,omitempty"`
	// 发布的环境id
	RunEnvId *string `json:"run_env_id,omitempty"`
	// API类型
	Type *int32 `json:"type,omitempty"`
	// 绑定的流控策略名称
	ThrottleName *string `json:"throttle_name,omitempty"`
	// API的访问地址
	ReqUri *string `json:"req_uri,omitempty"`
	// 发布的环境名
	RunEnvName *string `json:"run_env_name,omitempty"`
	// API所属分组的编号
	GroupId *string `json:"group_id,omitempty"`
	// API名称
	Name *string `json:"name,omitempty"`
	// API编号
	Id *string `json:"id,omitempty"`
}

func (o ThrottleBindingApiResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ThrottleBindingApiResp struct{}"
	}

	return strings.Join([]string{"ThrottleBindingApiResp", string(data)}, " ")
}
