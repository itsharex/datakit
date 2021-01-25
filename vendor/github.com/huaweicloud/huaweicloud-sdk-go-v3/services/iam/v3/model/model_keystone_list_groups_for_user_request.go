/*
 * IAM
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
type KeystoneListGroupsForUserRequest struct {
	UserId string `json:"user_id"`
}

func (o KeystoneListGroupsForUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListGroupsForUserRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListGroupsForUserRequest", string(data)}, " ")
}
