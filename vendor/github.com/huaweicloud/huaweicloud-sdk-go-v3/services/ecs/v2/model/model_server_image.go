/*
 * ECS
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 弹性云服务器镜像信息。
type ServerImage struct {
	// 镜像id
	Id string `json:"id"`
}

func (o ServerImage) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServerImage struct{}"
	}

	return strings.Join([]string{"ServerImage", string(data)}, " ")
}
