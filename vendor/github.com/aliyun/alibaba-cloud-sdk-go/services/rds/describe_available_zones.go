package rds

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

// DescribeAvailableZones invokes the rds.DescribeAvailableZones API synchronously
// api document: https://help.aliyun.com/api/rds/describeavailablezones.html
func (client *Client) DescribeAvailableZones(request *DescribeAvailableZonesRequest) (response *DescribeAvailableZonesResponse, err error) {
	response = CreateDescribeAvailableZonesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAvailableZonesWithChan invokes the rds.DescribeAvailableZones API asynchronously
// api document: https://help.aliyun.com/api/rds/describeavailablezones.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableZonesWithChan(request *DescribeAvailableZonesRequest) (<-chan *DescribeAvailableZonesResponse, <-chan error) {
	responseChan := make(chan *DescribeAvailableZonesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAvailableZones(request)
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

// DescribeAvailableZonesWithCallback invokes the rds.DescribeAvailableZones API asynchronously
// api document: https://help.aliyun.com/api/rds/describeavailablezones.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableZonesWithCallback(request *DescribeAvailableZonesRequest, callback func(response *DescribeAvailableZonesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAvailableZonesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAvailableZones(request)
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

// DescribeAvailableZonesRequest is the request struct for api DescribeAvailableZones
type DescribeAvailableZonesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	Engine               string           `position:"Query" name:"Engine"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DescribeAvailableZonesResponse is the response struct for api DescribeAvailableZones
type DescribeAvailableZonesResponse struct {
	*responses.BaseResponse
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	AvailableZones []AvailableZone `json:"AvailableZones" xml:"AvailableZones"`
}

// CreateDescribeAvailableZonesRequest creates a request to invoke DescribeAvailableZones API
func CreateDescribeAvailableZonesRequest() (request *DescribeAvailableZonesRequest) {
	request = &DescribeAvailableZonesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeAvailableZones", "rds", "openAPI")
	return
}

// CreateDescribeAvailableZonesResponse creates a response to parse from DescribeAvailableZones response
func CreateDescribeAvailableZonesResponse() (response *DescribeAvailableZonesResponse) {
	response = &DescribeAvailableZonesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
