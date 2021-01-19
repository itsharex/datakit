package slb

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

// ModifyLoadBalancerInstanceSpec invokes the slb.ModifyLoadBalancerInstanceSpec API synchronously
// api document: https://help.aliyun.com/api/slb/modifyloadbalancerinstancespec.html
func (client *Client) ModifyLoadBalancerInstanceSpec(request *ModifyLoadBalancerInstanceSpecRequest) (response *ModifyLoadBalancerInstanceSpecResponse, err error) {
	response = CreateModifyLoadBalancerInstanceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLoadBalancerInstanceSpecWithChan invokes the slb.ModifyLoadBalancerInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/slb/modifyloadbalancerinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLoadBalancerInstanceSpecWithChan(request *ModifyLoadBalancerInstanceSpecRequest) (<-chan *ModifyLoadBalancerInstanceSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyLoadBalancerInstanceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLoadBalancerInstanceSpec(request)
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

// ModifyLoadBalancerInstanceSpecWithCallback invokes the slb.ModifyLoadBalancerInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/slb/modifyloadbalancerinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLoadBalancerInstanceSpecWithCallback(request *ModifyLoadBalancerInstanceSpecRequest, callback func(response *ModifyLoadBalancerInstanceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLoadBalancerInstanceSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyLoadBalancerInstanceSpec(request)
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

// ModifyLoadBalancerInstanceSpecRequest is the request struct for api ModifyLoadBalancerInstanceSpec
type ModifyLoadBalancerInstanceSpecRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerSpec     string           `position:"Query" name:"LoadBalancerSpec"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// ModifyLoadBalancerInstanceSpecResponse is the response struct for api ModifyLoadBalancerInstanceSpec
type ModifyLoadBalancerInstanceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   int64  `json:"OrderId" xml:"OrderId"`
}

// CreateModifyLoadBalancerInstanceSpecRequest creates a request to invoke ModifyLoadBalancerInstanceSpec API
func CreateModifyLoadBalancerInstanceSpecRequest() (request *ModifyLoadBalancerInstanceSpecRequest) {
	request = &ModifyLoadBalancerInstanceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "ModifyLoadBalancerInstanceSpec", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLoadBalancerInstanceSpecResponse creates a response to parse from ModifyLoadBalancerInstanceSpec response
func CreateModifyLoadBalancerInstanceSpecResponse() (response *ModifyLoadBalancerInstanceSpecResponse) {
	response = &ModifyLoadBalancerInstanceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
