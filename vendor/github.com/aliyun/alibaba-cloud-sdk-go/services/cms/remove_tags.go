package cms

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

// RemoveTags invokes the cms.RemoveTags API synchronously
// api document: https://help.aliyun.com/api/cms/removetags.html
func (client *Client) RemoveTags(request *RemoveTagsRequest) (response *RemoveTagsResponse, err error) {
	response = CreateRemoveTagsResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveTagsWithChan invokes the cms.RemoveTags API asynchronously
// api document: https://help.aliyun.com/api/cms/removetags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveTagsWithChan(request *RemoveTagsRequest) (<-chan *RemoveTagsResponse, <-chan error) {
	responseChan := make(chan *RemoveTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveTags(request)
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

// RemoveTagsWithCallback invokes the cms.RemoveTags API asynchronously
// api document: https://help.aliyun.com/api/cms/removetags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveTagsWithCallback(request *RemoveTagsRequest, callback func(response *RemoveTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveTagsResponse
		var err error
		defer close(result)
		response, err = client.RemoveTags(request)
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

// RemoveTagsRequest is the request struct for api RemoveTags
type RemoveTagsRequest struct {
	*requests.RpcRequest
	GroupIds *[]string        `position:"Query" name:"GroupIds"  type:"Repeated"`
	Tag      *[]RemoveTagsTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// RemoveTagsTag is a repeated param struct in RemoveTagsRequest
type RemoveTagsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// RemoveTagsResponse is the response struct for api RemoveTags
type RemoveTagsResponse struct {
	*responses.BaseResponse
	Code      string          `json:"Code" xml:"Code"`
	Message   string          `json:"Message" xml:"Message"`
	Success   bool            `json:"Success" xml:"Success"`
	RequestId string          `json:"RequestId" xml:"RequestId"`
	Tag       TagInRemoveTags `json:"Tag" xml:"Tag"`
}

// CreateRemoveTagsRequest creates a request to invoke RemoveTags API
func CreateRemoveTagsRequest() (request *RemoveTagsRequest) {
	request = &RemoveTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "RemoveTags", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveTagsResponse creates a response to parse from RemoveTags response
func CreateRemoveTagsResponse() (response *RemoveTagsResponse) {
	response = &RemoveTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
