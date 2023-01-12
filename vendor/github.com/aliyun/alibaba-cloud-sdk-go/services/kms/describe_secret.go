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

// DescribeSecret invokes the kms.DescribeSecret API synchronously
func (client *Client) DescribeSecret(request *DescribeSecretRequest) (response *DescribeSecretResponse, err error) {
	response = CreateDescribeSecretResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSecretWithChan invokes the kms.DescribeSecret API asynchronously
func (client *Client) DescribeSecretWithChan(request *DescribeSecretRequest) (<-chan *DescribeSecretResponse, <-chan error) {
	responseChan := make(chan *DescribeSecretResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecret(request)
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

// DescribeSecretWithCallback invokes the kms.DescribeSecret API asynchronously
func (client *Client) DescribeSecretWithCallback(request *DescribeSecretRequest, callback func(response *DescribeSecretResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecretResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecret(request)
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

// DescribeSecretRequest is the request struct for api DescribeSecret
type DescribeSecretRequest struct {
	*requests.RpcRequest
	SecretName string `position:"Query" name:"SecretName"`
	FetchTags  string `position:"Query" name:"FetchTags"`
}

// DescribeSecretResponse is the response struct for api DescribeSecret
type DescribeSecretResponse struct {
	*responses.BaseResponse
	UpdateTime        string               `json:"UpdateTime" xml:"UpdateTime"`
	CreateTime        string               `json:"CreateTime" xml:"CreateTime"`
	NextRotationDate  string               `json:"NextRotationDate" xml:"NextRotationDate"`
	EncryptionKeyId   string               `json:"EncryptionKeyId" xml:"EncryptionKeyId"`
	RotationInterval  string               `json:"RotationInterval" xml:"RotationInterval"`
	Arn               string               `json:"Arn" xml:"Arn"`
	ExtendedConfig    string               `json:"ExtendedConfig" xml:"ExtendedConfig"`
	LastRotationDate  string               `json:"LastRotationDate" xml:"LastRotationDate"`
	RequestId         string               `json:"RequestId" xml:"RequestId"`
	Description       string               `json:"Description" xml:"Description"`
	SecretName        string               `json:"SecretName" xml:"SecretName"`
	AutomaticRotation string               `json:"AutomaticRotation" xml:"AutomaticRotation"`
	SecretType        string               `json:"SecretType" xml:"SecretType"`
	PlannedDeleteTime string               `json:"PlannedDeleteTime" xml:"PlannedDeleteTime"`
	DKMSInstanceId    string               `json:"DKMSInstanceId" xml:"DKMSInstanceId"`
	Tags              TagsInDescribeSecret `json:"Tags" xml:"Tags"`
}

// CreateDescribeSecretRequest creates a request to invoke DescribeSecret API
func CreateDescribeSecretRequest() (request *DescribeSecretRequest) {
	request = &DescribeSecretRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "DescribeSecret", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSecretResponse creates a response to parse from DescribeSecret response
func CreateDescribeSecretResponse() (response *DescribeSecretResponse) {
	response = &DescribeSecretResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
