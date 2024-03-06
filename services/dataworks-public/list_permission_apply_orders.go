package dataworks_public

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

// ListPermissionApplyOrders invokes the dataworks_public.ListPermissionApplyOrders API synchronously
func (client *Client) ListPermissionApplyOrders(request *ListPermissionApplyOrdersRequest) (response *ListPermissionApplyOrdersResponse, err error) {
	response = CreateListPermissionApplyOrdersResponse()
	err = client.DoAction(request, response)
	return
}

// ListPermissionApplyOrdersWithChan invokes the dataworks_public.ListPermissionApplyOrders API asynchronously
func (client *Client) ListPermissionApplyOrdersWithChan(request *ListPermissionApplyOrdersRequest) (<-chan *ListPermissionApplyOrdersResponse, <-chan error) {
	responseChan := make(chan *ListPermissionApplyOrdersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPermissionApplyOrders(request)
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

// ListPermissionApplyOrdersWithCallback invokes the dataworks_public.ListPermissionApplyOrders API asynchronously
func (client *Client) ListPermissionApplyOrdersWithCallback(request *ListPermissionApplyOrdersRequest, callback func(response *ListPermissionApplyOrdersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPermissionApplyOrdersResponse
		var err error
		defer close(result)
		response, err = client.ListPermissionApplyOrders(request)
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

// ListPermissionApplyOrdersRequest is the request struct for api ListPermissionApplyOrders
type ListPermissionApplyOrdersRequest struct {
	*requests.RpcRequest
	StartTime             requests.Integer `position:"Query" name:"StartTime"`
	PageNum               requests.Integer `position:"Query" name:"PageNum"`
	PageSize              requests.Integer `position:"Query" name:"PageSize"`
	TableName             string           `position:"Query" name:"TableName"`
	QueryType             requests.Integer `position:"Query" name:"QueryType"`
	EngineType            string           `position:"Query" name:"EngineType"`
	MaxComputeProjectName string           `position:"Query" name:"MaxComputeProjectName"`
	EndTime               requests.Integer `position:"Query" name:"EndTime"`
	FlowStatus            requests.Integer `position:"Query" name:"FlowStatus"`
	WorkspaceId           requests.Integer `position:"Query" name:"WorkspaceId"`
	OrderType             requests.Integer `position:"Query" name:"OrderType"`
}

// ListPermissionApplyOrdersResponse is the response struct for api ListPermissionApplyOrders
type ListPermissionApplyOrdersResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	ApplyOrders ApplyOrders `json:"ApplyOrders" xml:"ApplyOrders"`
}

// CreateListPermissionApplyOrdersRequest creates a request to invoke ListPermissionApplyOrders API
func CreateListPermissionApplyOrdersRequest() (request *ListPermissionApplyOrdersRequest) {
	request = &ListPermissionApplyOrdersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListPermissionApplyOrders", "", "")
	request.Method = requests.POST
	return
}

// CreateListPermissionApplyOrdersResponse creates a response to parse from ListPermissionApplyOrders response
func CreateListPermissionApplyOrdersResponse() (response *ListPermissionApplyOrdersResponse) {
	response = &ListPermissionApplyOrdersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
