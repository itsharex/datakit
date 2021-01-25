/*
 * DMS
 *
 * DMS Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type BatchCreateOrDeleteTagReqTags struct {
	// 键。最大长度36个unicode字符。  key不能为空，不能为空字符串。  不能包含下列字符：非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。
	Key *string `json:"key,omitempty"`
	// 值。每个值最大长度43个unicode字符。  value不能为空，可以空字符串。  不能包含下列字符：非打印字符ASCII(0-31), “=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。
	Value *string `json:"value,omitempty"`
}

func (o BatchCreateOrDeleteTagReqTags) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteTagReqTags struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteTagReqTags", string(data)}, " ")
}
