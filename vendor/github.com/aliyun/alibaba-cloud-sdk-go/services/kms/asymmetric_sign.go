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

// AsymmetricSign invokes the kms.AsymmetricSign API synchronously
func (client *Client) AsymmetricSign(request *AsymmetricSignRequest) (response *AsymmetricSignResponse, err error) {
	response = CreateAsymmetricSignResponse()
	err = client.DoAction(request, response)
	return
}

// AsymmetricSignWithChan invokes the kms.AsymmetricSign API asynchronously
func (client *Client) AsymmetricSignWithChan(request *AsymmetricSignRequest) (<-chan *AsymmetricSignResponse, <-chan error) {
	responseChan := make(chan *AsymmetricSignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AsymmetricSign(request)
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

// AsymmetricSignWithCallback invokes the kms.AsymmetricSign API asynchronously
func (client *Client) AsymmetricSignWithCallback(request *AsymmetricSignRequest, callback func(response *AsymmetricSignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AsymmetricSignResponse
		var err error
		defer close(result)
		response, err = client.AsymmetricSign(request)
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

// AsymmetricSignRequest is the request struct for api AsymmetricSign
type AsymmetricSignRequest struct {
	*requests.RpcRequest
	KeyVersionId string `position:"Query" name:"KeyVersionId"`
	Digest       string `position:"Query" name:"Digest"`
	KeyId        string `position:"Query" name:"KeyId"`
	Algorithm    string `position:"Query" name:"Algorithm"`
}

// AsymmetricSignResponse is the response struct for api AsymmetricSign
type AsymmetricSignResponse struct {
	*responses.BaseResponse
	KeyVersionId string `json:"KeyVersionId" xml:"KeyVersionId"`
	KeyId        string `json:"KeyId" xml:"KeyId"`
	Value        string `json:"Value" xml:"Value"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateAsymmetricSignRequest creates a request to invoke AsymmetricSign API
func CreateAsymmetricSignRequest() (request *AsymmetricSignRequest) {
	request = &AsymmetricSignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "AsymmetricSign", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAsymmetricSignResponse creates a response to parse from AsymmetricSign response
func CreateAsymmetricSignResponse() (response *AsymmetricSignResponse) {
	response = &AsymmetricSignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
