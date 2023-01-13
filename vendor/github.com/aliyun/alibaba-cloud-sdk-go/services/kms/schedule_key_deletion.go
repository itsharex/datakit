package kms

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

// ScheduleKeyDeletion invokes the kms.ScheduleKeyDeletion API synchronously
func (client *Client) ScheduleKeyDeletion(request *ScheduleKeyDeletionRequest) (response *ScheduleKeyDeletionResponse, err error) {
	response = CreateScheduleKeyDeletionResponse()
	err = client.DoAction(request, response)
	return
}

// ScheduleKeyDeletionWithChan invokes the kms.ScheduleKeyDeletion API asynchronously
func (client *Client) ScheduleKeyDeletionWithChan(request *ScheduleKeyDeletionRequest) (<-chan *ScheduleKeyDeletionResponse, <-chan error) {
	responseChan := make(chan *ScheduleKeyDeletionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ScheduleKeyDeletion(request)
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

// ScheduleKeyDeletionWithCallback invokes the kms.ScheduleKeyDeletion API asynchronously
func (client *Client) ScheduleKeyDeletionWithCallback(request *ScheduleKeyDeletionRequest, callback func(response *ScheduleKeyDeletionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ScheduleKeyDeletionResponse
		var err error
		defer close(result)
		response, err = client.ScheduleKeyDeletion(request)
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

// ScheduleKeyDeletionRequest is the request struct for api ScheduleKeyDeletion
type ScheduleKeyDeletionRequest struct {
	*requests.RpcRequest
	PendingWindowInDays requests.Integer `position:"Query" name:"PendingWindowInDays"`
	KeyId               string           `position:"Query" name:"KeyId"`
}

// ScheduleKeyDeletionResponse is the response struct for api ScheduleKeyDeletion
type ScheduleKeyDeletionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateScheduleKeyDeletionRequest creates a request to invoke ScheduleKeyDeletion API
func CreateScheduleKeyDeletionRequest() (request *ScheduleKeyDeletionRequest) {
	request = &ScheduleKeyDeletionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "ScheduleKeyDeletion", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateScheduleKeyDeletionResponse creates a response to parse from ScheduleKeyDeletion response
func CreateScheduleKeyDeletionResponse() (response *ScheduleKeyDeletionResponse) {
	response = &ScheduleKeyDeletionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
