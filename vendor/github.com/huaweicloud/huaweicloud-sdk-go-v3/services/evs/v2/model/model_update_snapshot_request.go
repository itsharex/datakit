/*
 * EVS
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
type UpdateSnapshotRequest struct {
	SnapshotId string                     `json:"snapshot_id"`
	Body       *UpdateSnapshotRequestBody `json:"body,omitempty"`
}

func (o UpdateSnapshotRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSnapshotRequest struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotRequest", string(data)}, " ")
}
