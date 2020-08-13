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

// SubmitEmailVerification invokes the domain.SubmitEmailVerification API synchronously
// api document: https://help.aliyun.com/api/domain/submitemailverification.html
func (client *Client) SubmitEmailVerification(request *SubmitEmailVerificationRequest) (response *SubmitEmailVerificationResponse, err error) {
	response = CreateSubmitEmailVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitEmailVerificationWithChan invokes the domain.SubmitEmailVerification API asynchronously
// api document: https://help.aliyun.com/api/domain/submitemailverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitEmailVerificationWithChan(request *SubmitEmailVerificationRequest) (<-chan *SubmitEmailVerificationResponse, <-chan error) {
	responseChan := make(chan *SubmitEmailVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitEmailVerification(request)
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

// SubmitEmailVerificationWithCallback invokes the domain.SubmitEmailVerification API asynchronously
// api document: https://help.aliyun.com/api/domain/submitemailverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitEmailVerificationWithCallback(request *SubmitEmailVerificationRequest, callback func(response *SubmitEmailVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitEmailVerificationResponse
		var err error
		defer close(result)
		response, err = client.SubmitEmailVerification(request)
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

// SubmitEmailVerificationRequest is the request struct for api SubmitEmailVerification
type SubmitEmailVerificationRequest struct {
	*requests.RpcRequest
	SendIfExist  requests.Boolean `position:"Query" name:"SendIfExist"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
	Email        string           `position:"Query" name:"Email"`
}

// SubmitEmailVerificationResponse is the response struct for api SubmitEmailVerification
type SubmitEmailVerificationResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	SuccessList []SendResult `json:"SuccessList" xml:"SuccessList"`
	FailList    []SendResult `json:"FailList" xml:"FailList"`
	ExistList   []SendResult `json:"ExistList" xml:"ExistList"`
}

// CreateSubmitEmailVerificationRequest creates a request to invoke SubmitEmailVerification API
func CreateSubmitEmailVerificationRequest() (request *SubmitEmailVerificationRequest) {
	request = &SubmitEmailVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SubmitEmailVerification", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitEmailVerificationResponse creates a response to parse from SubmitEmailVerification response
func CreateSubmitEmailVerificationResponse() (response *SubmitEmailVerificationResponse) {
	response = &SubmitEmailVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
