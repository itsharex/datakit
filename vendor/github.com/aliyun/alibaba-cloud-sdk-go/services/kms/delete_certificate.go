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

// DeleteCertificate invokes the kms.DeleteCertificate API synchronously
func (client *Client) DeleteCertificate(request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
	response = CreateDeleteCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCertificateWithChan invokes the kms.DeleteCertificate API asynchronously
func (client *Client) DeleteCertificateWithChan(request *DeleteCertificateRequest) (<-chan *DeleteCertificateResponse, <-chan error) {
	responseChan := make(chan *DeleteCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCertificate(request)
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

// DeleteCertificateWithCallback invokes the kms.DeleteCertificate API asynchronously
func (client *Client) DeleteCertificateWithCallback(request *DeleteCertificateRequest, callback func(response *DeleteCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCertificateResponse
		var err error
		defer close(result)
		response, err = client.DeleteCertificate(request)
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

// DeleteCertificateRequest is the request struct for api DeleteCertificate
type DeleteCertificateRequest struct {
	*requests.RpcRequest
	CertificateId string `position:"Query" name:"CertificateId"`
}

// DeleteCertificateResponse is the response struct for api DeleteCertificate
type DeleteCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCertificateRequest creates a request to invoke DeleteCertificate API
func CreateDeleteCertificateRequest() (request *DeleteCertificateRequest) {
	request = &DeleteCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "DeleteCertificate", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteCertificateResponse creates a response to parse from DeleteCertificate response
func CreateDeleteCertificateResponse() (response *DeleteCertificateResponse) {
	response = &DeleteCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
