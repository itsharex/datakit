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

// GenerateAndExportDataKey invokes the kms.GenerateAndExportDataKey API synchronously
func (client *Client) GenerateAndExportDataKey(request *GenerateAndExportDataKeyRequest) (response *GenerateAndExportDataKeyResponse, err error) {
	response = CreateGenerateAndExportDataKeyResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateAndExportDataKeyWithChan invokes the kms.GenerateAndExportDataKey API asynchronously
func (client *Client) GenerateAndExportDataKeyWithChan(request *GenerateAndExportDataKeyRequest) (<-chan *GenerateAndExportDataKeyResponse, <-chan error) {
	responseChan := make(chan *GenerateAndExportDataKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateAndExportDataKey(request)
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

// GenerateAndExportDataKeyWithCallback invokes the kms.GenerateAndExportDataKey API asynchronously
func (client *Client) GenerateAndExportDataKeyWithCallback(request *GenerateAndExportDataKeyRequest, callback func(response *GenerateAndExportDataKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateAndExportDataKeyResponse
		var err error
		defer close(result)
		response, err = client.GenerateAndExportDataKey(request)
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

// GenerateAndExportDataKeyRequest is the request struct for api GenerateAndExportDataKey
type GenerateAndExportDataKeyRequest struct {
	*requests.RpcRequest
	EncryptionContext string           `position:"Query" name:"EncryptionContext"`
	KeyId             string           `position:"Query" name:"KeyId"`
	KeySpec           string           `position:"Query" name:"KeySpec"`
	NumberOfBytes     requests.Integer `position:"Query" name:"NumberOfBytes"`
	WrappingAlgorithm string           `position:"Query" name:"WrappingAlgorithm"`
	PublicKeyBlob     string           `position:"Query" name:"PublicKeyBlob"`
	WrappingKeySpec   string           `position:"Query" name:"WrappingKeySpec"`
}

// GenerateAndExportDataKeyResponse is the response struct for api GenerateAndExportDataKey
type GenerateAndExportDataKeyResponse struct {
	*responses.BaseResponse
	KeyVersionId    string `json:"KeyVersionId" xml:"KeyVersionId"`
	KeyId           string `json:"KeyId" xml:"KeyId"`
	CiphertextBlob  string `json:"CiphertextBlob" xml:"CiphertextBlob"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	ExportedDataKey string `json:"ExportedDataKey" xml:"ExportedDataKey"`
}

// CreateGenerateAndExportDataKeyRequest creates a request to invoke GenerateAndExportDataKey API
func CreateGenerateAndExportDataKeyRequest() (request *GenerateAndExportDataKeyRequest) {
	request = &GenerateAndExportDataKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "GenerateAndExportDataKey", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGenerateAndExportDataKeyResponse creates a response to parse from GenerateAndExportDataKey response
func CreateGenerateAndExportDataKeyResponse() (response *GenerateAndExportDataKeyResponse) {
	response = &GenerateAndExportDataKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
