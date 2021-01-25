/*
 * CloudIDE
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StartInstanceRequest struct {
	InstanceId string              `json:"instance_id"`
	Body       *StartInstanceParam `json:"body,omitempty"`
}

func (o StartInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceRequest", string(data)}, " ")
}
