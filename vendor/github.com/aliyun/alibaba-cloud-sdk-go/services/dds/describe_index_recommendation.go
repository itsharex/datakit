package dds

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

// DescribeIndexRecommendation invokes the dds.DescribeIndexRecommendation API synchronously
// api document: https://help.aliyun.com/api/dds/describeindexrecommendation.html
func (client *Client) DescribeIndexRecommendation(request *DescribeIndexRecommendationRequest) (response *DescribeIndexRecommendationResponse, err error) {
	response = CreateDescribeIndexRecommendationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIndexRecommendationWithChan invokes the dds.DescribeIndexRecommendation API asynchronously
// api document: https://help.aliyun.com/api/dds/describeindexrecommendation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeIndexRecommendationWithChan(request *DescribeIndexRecommendationRequest) (<-chan *DescribeIndexRecommendationResponse, <-chan error) {
	responseChan := make(chan *DescribeIndexRecommendationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIndexRecommendation(request)
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

// DescribeIndexRecommendationWithCallback invokes the dds.DescribeIndexRecommendation API asynchronously
// api document: https://help.aliyun.com/api/dds/describeindexrecommendation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeIndexRecommendationWithCallback(request *DescribeIndexRecommendationRequest, callback func(response *DescribeIndexRecommendationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIndexRecommendationResponse
		var err error
		defer close(result)
		response, err = client.DescribeIndexRecommendation(request)
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

// DescribeIndexRecommendationRequest is the request struct for api DescribeIndexRecommendation
type DescribeIndexRecommendationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	Database             string           `position:"Query" name:"Database"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	NodeId               string           `position:"Query" name:"NodeId"`
	TaskId               string           `position:"Query" name:"TaskId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	Collection           string           `position:"Query" name:"Collection"`
	OperationType        string           `position:"Query" name:"OperationType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeIndexRecommendationResponse is the response struct for api DescribeIndexRecommendation
type DescribeIndexRecommendationResponse struct {
	*responses.BaseResponse
	RequestId        string       `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int          `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int          `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int          `json:"PageRecordCount" xml:"PageRecordCount"`
	Analyzations     Analyzations `json:"Analyzations" xml:"Analyzations"`
}

// CreateDescribeIndexRecommendationRequest creates a request to invoke DescribeIndexRecommendation API
func CreateDescribeIndexRecommendationRequest() (request *DescribeIndexRecommendationRequest) {
	request = &DescribeIndexRecommendationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeIndexRecommendation", "Dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeIndexRecommendationResponse creates a response to parse from DescribeIndexRecommendation response
func CreateDescribeIndexRecommendationResponse() (response *DescribeIndexRecommendationResponse) {
	response = &DescribeIndexRecommendationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
