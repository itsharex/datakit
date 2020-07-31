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

// AddLiveAppRecordConfig invokes the cdn.AddLiveAppRecordConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/addliveapprecordconfig.html
func (client *Client) AddLiveAppRecordConfig(request *AddLiveAppRecordConfigRequest) (response *AddLiveAppRecordConfigResponse, err error) {
	response = CreateAddLiveAppRecordConfigResponse()
	err = client.DoAction(request, response)
	return
}

// AddLiveAppRecordConfigWithChan invokes the cdn.AddLiveAppRecordConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/addliveapprecordconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLiveAppRecordConfigWithChan(request *AddLiveAppRecordConfigRequest) (<-chan *AddLiveAppRecordConfigResponse, <-chan error) {
	responseChan := make(chan *AddLiveAppRecordConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLiveAppRecordConfig(request)
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

// AddLiveAppRecordConfigWithCallback invokes the cdn.AddLiveAppRecordConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/addliveapprecordconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLiveAppRecordConfigWithCallback(request *AddLiveAppRecordConfigRequest, callback func(response *AddLiveAppRecordConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLiveAppRecordConfigResponse
		var err error
		defer close(result)
		response, err = client.AddLiveAppRecordConfig(request)
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

// AddLiveAppRecordConfigRequest is the request struct for api AddLiveAppRecordConfig
type AddLiveAppRecordConfigRequest struct {
	*requests.RpcRequest
	OssEndpoint     string           `position:"Query" name:"OssEndpoint"`
	AppName         string           `position:"Query" name:"AppName"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	OssBucket       string           `position:"Query" name:"OssBucket"`
	DomainName      string           `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	OssObjectPrefix string           `position:"Query" name:"OssObjectPrefix"`
}

// AddLiveAppRecordConfigResponse is the response struct for api AddLiveAppRecordConfig
type AddLiveAppRecordConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLiveAppRecordConfigRequest creates a request to invoke AddLiveAppRecordConfig API
func CreateAddLiveAppRecordConfigRequest() (request *AddLiveAppRecordConfigRequest) {
	request = &AddLiveAppRecordConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "AddLiveAppRecordConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateAddLiveAppRecordConfigResponse creates a response to parse from AddLiveAppRecordConfig response
func CreateAddLiveAppRecordConfigResponse() (response *AddLiveAppRecordConfigResponse) {
	response = &AddLiveAppRecordConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
