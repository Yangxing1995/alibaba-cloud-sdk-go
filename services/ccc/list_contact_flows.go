package ccc

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

// ListContactFlows invokes the ccc.ListContactFlows API synchronously
func (client *Client) ListContactFlows(request *ListContactFlowsRequest) (response *ListContactFlowsResponse, err error) {
	response = CreateListContactFlowsResponse()
	err = client.DoAction(request, response)
	return
}

// ListContactFlowsWithChan invokes the ccc.ListContactFlows API asynchronously
func (client *Client) ListContactFlowsWithChan(request *ListContactFlowsRequest) (<-chan *ListContactFlowsResponse, <-chan error) {
	responseChan := make(chan *ListContactFlowsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListContactFlows(request)
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

// ListContactFlowsWithCallback invokes the ccc.ListContactFlows API asynchronously
func (client *Client) ListContactFlowsWithCallback(request *ListContactFlowsRequest, callback func(response *ListContactFlowsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListContactFlowsResponse
		var err error
		defer close(result)
		response, err = client.ListContactFlows(request)
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

// ListContactFlowsRequest is the request struct for api ListContactFlows
type ListContactFlowsRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListContactFlowsResponse is the response struct for api ListContactFlows
type ListContactFlowsResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Success        bool         `json:"Success" xml:"Success"`
	Code           string       `json:"Code" xml:"Code"`
	Message        string       `json:"Message" xml:"Message"`
	HttpStatusCode int          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ContactFlows   ContactFlows `json:"ContactFlows" xml:"ContactFlows"`
}

// CreateListContactFlowsRequest creates a request to invoke ListContactFlows API
func CreateListContactFlowsRequest() (request *ListContactFlowsRequest) {
	request = &ListContactFlowsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListContactFlows", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListContactFlowsResponse creates a response to parse from ListContactFlows response
func CreateListContactFlowsResponse() (response *ListContactFlowsResponse) {
	response = &ListContactFlowsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
