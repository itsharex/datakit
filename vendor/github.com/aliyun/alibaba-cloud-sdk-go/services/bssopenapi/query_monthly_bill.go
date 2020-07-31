package bssopenapi

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

// QueryMonthlyBill invokes the bssopenapi.QueryMonthlyBill API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/querymonthlybill.html
func (client *Client) QueryMonthlyBill(request *QueryMonthlyBillRequest) (response *QueryMonthlyBillResponse, err error) {
	response = CreateQueryMonthlyBillResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMonthlyBillWithChan invokes the bssopenapi.QueryMonthlyBill API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querymonthlybill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMonthlyBillWithChan(request *QueryMonthlyBillRequest) (<-chan *QueryMonthlyBillResponse, <-chan error) {
	responseChan := make(chan *QueryMonthlyBillResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMonthlyBill(request)
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

// QueryMonthlyBillWithCallback invokes the bssopenapi.QueryMonthlyBill API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querymonthlybill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMonthlyBillWithCallback(request *QueryMonthlyBillRequest, callback func(response *QueryMonthlyBillResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMonthlyBillResponse
		var err error
		defer close(result)
		response, err = client.QueryMonthlyBill(request)
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

// QueryMonthlyBillRequest is the request struct for api QueryMonthlyBill
type QueryMonthlyBillRequest struct {
	*requests.RpcRequest
	BillingCycle string `position:"Query" name:"BillingCycle"`
}

// QueryMonthlyBillResponse is the response struct for api QueryMonthlyBill
type QueryMonthlyBillResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryMonthlyBillRequest creates a request to invoke QueryMonthlyBill API
func CreateQueryMonthlyBillRequest() (request *QueryMonthlyBillRequest) {
	request = &QueryMonthlyBillRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryMonthlyBill", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryMonthlyBillResponse creates a response to parse from QueryMonthlyBill response
func CreateQueryMonthlyBillResponse() (response *QueryMonthlyBillResponse) {
	response = &QueryMonthlyBillResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
