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

// DescribeWebshell invokes the aegis.DescribeWebshell API synchronously
// api document: https://help.aliyun.com/api/aegis/describewebshell.html
func (client *Client) DescribeWebshell(request *DescribeWebshellRequest) (response *DescribeWebshellResponse, err error) {
	response = CreateDescribeWebshellResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebshellWithChan invokes the aegis.DescribeWebshell API asynchronously
// api document: https://help.aliyun.com/api/aegis/describewebshell.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebshellWithChan(request *DescribeWebshellRequest) (<-chan *DescribeWebshellResponse, <-chan error) {
	responseChan := make(chan *DescribeWebshellResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebshell(request)
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

// DescribeWebshellWithCallback invokes the aegis.DescribeWebshell API asynchronously
// api document: https://help.aliyun.com/api/aegis/describewebshell.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebshellWithCallback(request *DescribeWebshellRequest, callback func(response *DescribeWebshellResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebshellResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebshell(request)
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

// DescribeWebshellRequest is the request struct for api DescribeWebshell
type DescribeWebshellRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	GroupId  requests.Integer `position:"Query" name:"GroupId"`
	Remark   string           `position:"Query" name:"Remark"`
	Dealed   string           `position:"Query" name:"Dealed"`
	Tag      requests.Integer `position:"Query" name:"Tag"`
}

// DescribeWebshellResponse is the response struct for api DescribeWebshell
type DescribeWebshellResponse struct {
	*responses.BaseResponse
	RequestId    string             `json:"RequestId" xml:"RequestId"`
	WebshellList []WebshellListItem `json:"WebshellList" xml:"WebshellList"`
}

// CreateDescribeWebshellRequest creates a request to invoke DescribeWebshell API
func CreateDescribeWebshellRequest() (request *DescribeWebshellRequest) {
	request = &DescribeWebshellRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeWebshell", "vipaegis", "openAPI")
	return
}

// CreateDescribeWebshellResponse creates a response to parse from DescribeWebshell response
func CreateDescribeWebshellResponse() (response *DescribeWebshellResponse) {
	response = &DescribeWebshellResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
