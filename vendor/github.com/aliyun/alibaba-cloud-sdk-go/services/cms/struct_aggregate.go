package cms

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

// Aggregate is a nested struct in cms response
type Aggregate struct {
	Alias     string `json:"Alias" xml:"Alias"`
	FieldName string `json:"FieldName" xml:"FieldName"`
	Function  string `json:"Function" xml:"Function"`
	Min       string `json:"Min" xml:"Min"`
	Max       string `json:"Max" xml:"Max"`
}
