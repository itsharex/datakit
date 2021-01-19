package aegis

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

// ModifySaveVulBatch invokes the aegis.ModifySaveVulBatch API synchronously
// api document: https://help.aliyun.com/api/aegis/modifysavevulbatch.html
func (client *Client) ModifySaveVulBatch(request *ModifySaveVulBatchRequest) (response *ModifySaveVulBatchResponse, err error) {
	response = CreateModifySaveVulBatchResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySaveVulBatchWithChan invokes the aegis.ModifySaveVulBatch API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifysavevulbatch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySaveVulBatchWithChan(request *ModifySaveVulBatchRequest) (<-chan *ModifySaveVulBatchResponse, <-chan error) {
	responseChan := make(chan *ModifySaveVulBatchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySaveVulBatch(request)
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

// ModifySaveVulBatchWithCallback invokes the aegis.ModifySaveVulBatch API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifysavevulbatch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySaveVulBatchWithCallback(request *ModifySaveVulBatchRequest, callback func(response *ModifySaveVulBatchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySaveVulBatchResponse
		var err error
		defer close(result)
		response, err = client.ModifySaveVulBatch(request)
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

// ModifySaveVulBatchRequest is the request struct for api ModifySaveVulBatch
type ModifySaveVulBatchRequest struct {
	*requests.RpcRequest
	BatchName  string `position:"Query" name:"BatchName"`
	AliasName  string `position:"Query" name:"AliasName"`
	StatusList string `position:"Query" name:"StatusList"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Level      string `position:"Query" name:"Level"`
	Resource   string `position:"Query" name:"Resource"`
	Name       string `position:"Query" name:"Name"`
	Dealed     string `position:"Query" name:"Dealed"`
	Remark     string `position:"Query" name:"Remark"`
	Type       string `position:"Query" name:"Type"`
	Necessity  string `position:"Query" name:"Necessity"`
	Uuids      string `position:"Query" name:"Uuids"`
}

// ModifySaveVulBatchResponse is the response struct for api ModifySaveVulBatch
type ModifySaveVulBatchResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySaveVulBatchRequest creates a request to invoke ModifySaveVulBatch API
func CreateModifySaveVulBatchRequest() (request *ModifySaveVulBatchRequest) {
	request = &ModifySaveVulBatchRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifySaveVulBatch", "vipaegis", "openAPI")
	return
}

// CreateModifySaveVulBatchResponse creates a response to parse from ModifySaveVulBatch response
func CreateModifySaveVulBatchResponse() (response *ModifySaveVulBatchResponse) {
	response = &ModifySaveVulBatchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
