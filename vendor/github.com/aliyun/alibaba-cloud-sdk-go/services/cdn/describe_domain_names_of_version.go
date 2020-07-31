package cdn

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

// DescribeDomainNamesOfVersion invokes the cdn.DescribeDomainNamesOfVersion API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainnamesofversion.html
func (client *Client) DescribeDomainNamesOfVersion(request *DescribeDomainNamesOfVersionRequest) (response *DescribeDomainNamesOfVersionResponse, err error) {
	response = CreateDescribeDomainNamesOfVersionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainNamesOfVersionWithChan invokes the cdn.DescribeDomainNamesOfVersion API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainnamesofversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainNamesOfVersionWithChan(request *DescribeDomainNamesOfVersionRequest) (<-chan *DescribeDomainNamesOfVersionResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainNamesOfVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainNamesOfVersion(request)
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

// DescribeDomainNamesOfVersionWithCallback invokes the cdn.DescribeDomainNamesOfVersion API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainnamesofversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainNamesOfVersionWithCallback(request *DescribeDomainNamesOfVersionRequest, callback func(response *DescribeDomainNamesOfVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainNamesOfVersionResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainNamesOfVersion(request)
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

// DescribeDomainNamesOfVersionRequest is the request struct for api DescribeDomainNamesOfVersion
type DescribeDomainNamesOfVersionRequest struct {
	*requests.RpcRequest
	VersionId string           `position:"Query" name:"VersionId"`
	PageSize  string           `position:"Query" name:"PageSize"`
	PageIndex requests.Integer `position:"Query" name:"PageIndex"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDomainNamesOfVersionResponse is the response struct for api DescribeDomainNamesOfVersion
type DescribeDomainNamesOfVersionResponse struct {
	*responses.BaseResponse
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	Contents   []Content `json:"Contents" xml:"Contents"`
}

// CreateDescribeDomainNamesOfVersionRequest creates a request to invoke DescribeDomainNamesOfVersion API
func CreateDescribeDomainNamesOfVersionRequest() (request *DescribeDomainNamesOfVersionRequest) {
	request = &DescribeDomainNamesOfVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainNamesOfVersion", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainNamesOfVersionResponse creates a response to parse from DescribeDomainNamesOfVersion response
func CreateDescribeDomainNamesOfVersionResponse() (response *DescribeDomainNamesOfVersionResponse) {
	response = &DescribeDomainNamesOfVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
