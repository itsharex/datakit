/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateSecurityGroupRequestBody struct {
	// 功能说明：是否只预检此次请求 取值范围： -true：发送检查请求，不会更新安全组。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。 -false（默认值）：发送正常请求，并直接更新安全组。
	DryRun        *bool                      `json:"dry_run,omitempty"`
	SecurityGroup *UpdateSecurityGroupOption `json:"security_group"`
}

func (o UpdateSecurityGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupRequestBody", string(data)}, " ")
}
