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

// CertificatePrivateKeySign invokes the kms.CertificatePrivateKeySign API synchronously
func (client *Client) CertificatePrivateKeySign(request *CertificatePrivateKeySignRequest) (response *CertificatePrivateKeySignResponse, err error) {
	response = CreateCertificatePrivateKeySignResponse()
	err = client.DoAction(request, response)
	return
}

// CertificatePrivateKeySignWithChan invokes the kms.CertificatePrivateKeySign API asynchronously
func (client *Client) CertificatePrivateKeySignWithChan(request *CertificatePrivateKeySignRequest) (<-chan *CertificatePrivateKeySignResponse, <-chan error) {
	responseChan := make(chan *CertificatePrivateKeySignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CertificatePrivateKeySign(request)
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

// CertificatePrivateKeySignWithCallback invokes the kms.CertificatePrivateKeySign API asynchronously
func (client *Client) CertificatePrivateKeySignWithCallback(request *CertificatePrivateKeySignRequest, callback func(response *CertificatePrivateKeySignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CertificatePrivateKeySignResponse
		var err error
		defer close(result)
		response, err = client.CertificatePrivateKeySign(request)
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

// CertificatePrivateKeySignRequest is the request struct for api CertificatePrivateKeySign
type CertificatePrivateKeySignRequest struct {
	*requests.RpcRequest
	MessageType   string `position:"Query" name:"MessageType"`
	CertificateId string `position:"Query" name:"CertificateId"`
	Message       string `position:"Query" name:"Message"`
	Algorithm     string `position:"Query" name:"Algorithm"`
}

// CertificatePrivateKeySignResponse is the response struct for api CertificatePrivateKeySign
type CertificatePrivateKeySignResponse struct {
	*responses.BaseResponse
	SignatureValue string `json:"SignatureValue" xml:"SignatureValue"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	CertificateId  string `json:"CertificateId" xml:"CertificateId"`
}

// CreateCertificatePrivateKeySignRequest creates a request to invoke CertificatePrivateKeySign API
func CreateCertificatePrivateKeySignRequest() (request *CertificatePrivateKeySignRequest) {
	request = &CertificatePrivateKeySignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "CertificatePrivateKeySign", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCertificatePrivateKeySignResponse creates a response to parse from CertificatePrivateKeySign response
func CreateCertificatePrivateKeySignResponse() (response *CertificatePrivateKeySignResponse) {
	response = &CertificatePrivateKeySignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
