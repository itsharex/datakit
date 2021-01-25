/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListBackupsResponse struct {
	// 备份列表。
	Backups *[]BackupForList `json:"backups,omitempty"`
	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBackupsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupsResponse", string(data)}, " ")
}
