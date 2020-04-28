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

// DeleteMetricRuleTemplate invokes the cms.DeleteMetricRuleTemplate API synchronously
// api document: https://help.aliyun.com/api/cms/deletemetricruletemplate.html
func (client *Client) DeleteMetricRuleTemplate(request *DeleteMetricRuleTemplateRequest) (response *DeleteMetricRuleTemplateResponse, err error) {
	response = CreateDeleteMetricRuleTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMetricRuleTemplateWithChan invokes the cms.DeleteMetricRuleTemplate API asynchronously
// api document: https://help.aliyun.com/api/cms/deletemetricruletemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMetricRuleTemplateWithChan(request *DeleteMetricRuleTemplateRequest) (<-chan *DeleteMetricRuleTemplateResponse, <-chan error) {
	responseChan := make(chan *DeleteMetricRuleTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMetricRuleTemplate(request)
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

// DeleteMetricRuleTemplateWithCallback invokes the cms.DeleteMetricRuleTemplate API asynchronously
// api document: https://help.aliyun.com/api/cms/deletemetricruletemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMetricRuleTemplateWithCallback(request *DeleteMetricRuleTemplateRequest, callback func(response *DeleteMetricRuleTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMetricRuleTemplateResponse
		var err error
		defer close(result)
		response, err = client.DeleteMetricRuleTemplate(request)
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

// DeleteMetricRuleTemplateRequest is the request struct for api DeleteMetricRuleTemplate
type DeleteMetricRuleTemplateRequest struct {
	*requests.RpcRequest
	TemplateId string `position:"Query" name:"TemplateId"`
}

// DeleteMetricRuleTemplateResponse is the response struct for api DeleteMetricRuleTemplate
type DeleteMetricRuleTemplateResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Success   bool     `json:"Success" xml:"Success"`
	Code      int      `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	Resource  Resource `json:"Resource" xml:"Resource"`
}

// CreateDeleteMetricRuleTemplateRequest creates a request to invoke DeleteMetricRuleTemplate API
func CreateDeleteMetricRuleTemplateRequest() (request *DeleteMetricRuleTemplateRequest) {
	request = &DeleteMetricRuleTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteMetricRuleTemplate", "cms", "openAPI")
	return
}

// CreateDeleteMetricRuleTemplateResponse creates a response to parse from DeleteMetricRuleTemplate response
func CreateDeleteMetricRuleTemplateResponse() (response *DeleteMetricRuleTemplateResponse) {
	response = &DeleteMetricRuleTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
