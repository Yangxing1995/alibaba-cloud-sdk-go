package dms_enterprise

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

// ListEffectiveOrders invokes the dms_enterprise.ListEffectiveOrders API synchronously
func (client *Client) ListEffectiveOrders(request *ListEffectiveOrdersRequest) (response *ListEffectiveOrdersResponse, err error) {
	response = CreateListEffectiveOrdersResponse()
	err = client.DoAction(request, response)
	return
}

// ListEffectiveOrdersWithChan invokes the dms_enterprise.ListEffectiveOrders API asynchronously
func (client *Client) ListEffectiveOrdersWithChan(request *ListEffectiveOrdersRequest) (<-chan *ListEffectiveOrdersResponse, <-chan error) {
	responseChan := make(chan *ListEffectiveOrdersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEffectiveOrders(request)
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

// ListEffectiveOrdersWithCallback invokes the dms_enterprise.ListEffectiveOrders API asynchronously
func (client *Client) ListEffectiveOrdersWithCallback(request *ListEffectiveOrdersRequest, callback func(response *ListEffectiveOrdersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEffectiveOrdersResponse
		var err error
		defer close(result)
		response, err = client.ListEffectiveOrders(request)
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

// ListEffectiveOrdersRequest is the request struct for api ListEffectiveOrders
type ListEffectiveOrdersRequest struct {
	*requests.RpcRequest
	Tid requests.Integer `position:"Query" name:"Tid"`
}

// ListEffectiveOrdersResponse is the response struct for api ListEffectiveOrders
type ListEffectiveOrdersResponse struct {
	*responses.BaseResponse
	RequestId    string             `json:"RequestId" xml:"RequestId"`
	Success      bool               `json:"Success" xml:"Success"`
	ErrorMessage string             `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string             `json:"ErrorCode" xml:"ErrorCode"`
	OrderSummary []OrderSummaryItem `json:"OrderSummary" xml:"OrderSummary"`
}

// CreateListEffectiveOrdersRequest creates a request to invoke ListEffectiveOrders API
func CreateListEffectiveOrdersRequest() (request *ListEffectiveOrdersRequest) {
	request = &ListEffectiveOrdersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListEffectiveOrders", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListEffectiveOrdersResponse creates a response to parse from ListEffectiveOrders response
func CreateListEffectiveOrdersResponse() (response *ListEffectiveOrdersResponse) {
	response = &ListEffectiveOrdersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
