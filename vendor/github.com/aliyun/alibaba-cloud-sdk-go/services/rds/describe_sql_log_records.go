package rds

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeSQLLogRecords invokes the rds.DescribeSQLLogRecords API synchronously
// api document: https://help.aliyun.com/api/rds/describesqllogrecords.html
func (client *Client) DescribeSQLLogRecords(request *DescribeSQLLogRecordsRequest) (response *DescribeSQLLogRecordsResponse, err error) {
	response = CreateDescribeSQLLogRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSQLLogRecordsWithChan invokes the rds.DescribeSQLLogRecords API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqllogrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLLogRecordsWithChan(request *DescribeSQLLogRecordsRequest) (<-chan *DescribeSQLLogRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLLogRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLLogRecords(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeSQLLogRecordsWithCallback invokes the rds.DescribeSQLLogRecords API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqllogrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLLogRecordsWithCallback(request *DescribeSQLLogRecordsRequest, callback func(response *DescribeSQLLogRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLLogRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLLogRecords(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeSQLLogRecordsRequest is the request struct for api DescribeSQLLogRecords
type DescribeSQLLogRecordsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	StartTime            string           `position:"Query" name:"StartTime"`
	QueryKeywords        string           `position:"Query" name:"QueryKeywords"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	Database             string           `position:"Query" name:"Database"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	SQLId                requests.Integer `position:"Query" name:"SQLId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Form                 string           `position:"Query" name:"Form"`
	User                 string           `position:"Query" name:"User"`
}

// DescribeSQLLogRecordsResponse is the response struct for api DescribeSQLLogRecords
type DescribeSQLLogRecordsResponse struct {
	*responses.BaseResponse
	RequestId        string                       `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int64                        `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                          `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                          `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeSQLLogRecords `json:"Items" xml:"Items"`
}

// CreateDescribeSQLLogRecordsRequest creates a request to invoke DescribeSQLLogRecords API
func CreateDescribeSQLLogRecordsRequest() (request *DescribeSQLLogRecordsRequest) {
	request = &DescribeSQLLogRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLLogRecords", "rds", "openAPI")
	return
}

// CreateDescribeSQLLogRecordsResponse creates a response to parse from DescribeSQLLogRecords response
func CreateDescribeSQLLogRecordsResponse() (response *DescribeSQLLogRecordsResponse) {
	response = &DescribeSQLLogRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
