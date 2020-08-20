package domain

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

// CheckMaxYearOfServerLock invokes the domain.CheckMaxYearOfServerLock API synchronously
// api document: https://help.aliyun.com/api/domain/checkmaxyearofserverlock.html
func (client *Client) CheckMaxYearOfServerLock(request *CheckMaxYearOfServerLockRequest) (response *CheckMaxYearOfServerLockResponse, err error) {
	response = CreateCheckMaxYearOfServerLockResponse()
	err = client.DoAction(request, response)
	return
}

// CheckMaxYearOfServerLockWithChan invokes the domain.CheckMaxYearOfServerLock API asynchronously
// api document: https://help.aliyun.com/api/domain/checkmaxyearofserverlock.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckMaxYearOfServerLockWithChan(request *CheckMaxYearOfServerLockRequest) (<-chan *CheckMaxYearOfServerLockResponse, <-chan error) {
	responseChan := make(chan *CheckMaxYearOfServerLockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckMaxYearOfServerLock(request)
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

// CheckMaxYearOfServerLockWithCallback invokes the domain.CheckMaxYearOfServerLock API asynchronously
// api document: https://help.aliyun.com/api/domain/checkmaxyearofserverlock.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckMaxYearOfServerLockWithCallback(request *CheckMaxYearOfServerLockRequest, callback func(response *CheckMaxYearOfServerLockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckMaxYearOfServerLockResponse
		var err error
		defer close(result)
		response, err = client.CheckMaxYearOfServerLock(request)
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

// CheckMaxYearOfServerLockRequest is the request struct for api CheckMaxYearOfServerLock
type CheckMaxYearOfServerLockRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	CheckAction  string `position:"Query" name:"CheckAction"`
	Lang         string `position:"Query" name:"Lang"`
}

// CheckMaxYearOfServerLockResponse is the response struct for api CheckMaxYearOfServerLock
type CheckMaxYearOfServerLockResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MaxYear   int    `json:"MaxYear" xml:"MaxYear"`
}

// CreateCheckMaxYearOfServerLockRequest creates a request to invoke CheckMaxYearOfServerLock API
func CreateCheckMaxYearOfServerLockRequest() (request *CheckMaxYearOfServerLockRequest) {
	request = &CheckMaxYearOfServerLockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "CheckMaxYearOfServerLock", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckMaxYearOfServerLockResponse creates a response to parse from CheckMaxYearOfServerLock response
func CreateCheckMaxYearOfServerLockResponse() (response *CheckMaxYearOfServerLockResponse) {
	response = &CheckMaxYearOfServerLockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
