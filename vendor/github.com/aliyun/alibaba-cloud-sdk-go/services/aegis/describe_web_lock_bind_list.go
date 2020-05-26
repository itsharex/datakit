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

// DescribeWebLockBindList invokes the aegis.DescribeWebLockBindList API synchronously
// api document: https://help.aliyun.com/api/aegis/describeweblockbindlist.html
func (client *Client) DescribeWebLockBindList(request *DescribeWebLockBindListRequest) (response *DescribeWebLockBindListResponse, err error) {
	response = CreateDescribeWebLockBindListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebLockBindListWithChan invokes the aegis.DescribeWebLockBindList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeweblockbindlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebLockBindListWithChan(request *DescribeWebLockBindListRequest) (<-chan *DescribeWebLockBindListResponse, <-chan error) {
	responseChan := make(chan *DescribeWebLockBindListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebLockBindList(request)
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

// DescribeWebLockBindListWithCallback invokes the aegis.DescribeWebLockBindList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeweblockbindlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebLockBindListWithCallback(request *DescribeWebLockBindListRequest, callback func(response *DescribeWebLockBindListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebLockBindListResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebLockBindList(request)
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

// DescribeWebLockBindListRequest is the request struct for api DescribeWebLockBindList
type DescribeWebLockBindListRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Remark      string           `position:"Query" name:"Remark"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Lang        string           `position:"Query" name:"Lang"`
	Status      string           `position:"Query" name:"Status"`
}

// DescribeWebLockBindListResponse is the response struct for api DescribeWebLockBindList
type DescribeWebLockBindListResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	PageSize    int    `json:"PageSize" xml:"PageSize"`
	CurrentPage int    `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int    `json:"TotalCount" xml:"TotalCount"`
	BindList    []Bind `json:"BindList" xml:"BindList"`
}

// CreateDescribeWebLockBindListRequest creates a request to invoke DescribeWebLockBindList API
func CreateDescribeWebLockBindListRequest() (request *DescribeWebLockBindListRequest) {
	request = &DescribeWebLockBindListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeWebLockBindList", "vipaegis", "openAPI")
	return
}

// CreateDescribeWebLockBindListResponse creates a response to parse from DescribeWebLockBindList response
func CreateDescribeWebLockBindListResponse() (response *DescribeWebLockBindListResponse) {
	response = &DescribeWebLockBindListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
