package sas

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionListArr is a nested struct in sas response
type ActionListArr struct {
	GmtModified  int64  `json:"GmtModified" xml:"GmtModified"`
	Id           int    `json:"Id" xml:"Id"`
	AliUid       int64  `json:"AliUid" xml:"AliUid"`
	GmtCreate    int64  `json:"GmtCreate" xml:"GmtCreate"`
	Url          string `json:"Url" xml:"Url"`
	IntervalTime int    `json:"IntervalTime" xml:"IntervalTime"`
	ActionName   string `json:"ActionName" xml:"ActionName"`
	Status       int    `json:"Status" xml:"Status"`
	ConfigList   string `json:"ConfigList" xml:"ConfigList"`
	GroupIdList  string `json:"GroupIdList" xml:"GroupIdList"`
	DingTalkLang string `json:"DingTalkLang" xml:"DingTalkLang"`
}
