package schedulerx2

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

// ListNamespaces invokes the schedulerx2.ListNamespaces API synchronously
func (client *Client) ListNamespaces(request *ListNamespacesRequest) (response *ListNamespacesResponse, err error) {
	response = CreateListNamespacesResponse()
	err = client.DoAction(request, response)
	return
}

// ListNamespacesWithChan invokes the schedulerx2.ListNamespaces API asynchronously
func (client *Client) ListNamespacesWithChan(request *ListNamespacesRequest) (<-chan *ListNamespacesResponse, <-chan error) {
	responseChan := make(chan *ListNamespacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNamespaces(request)
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

// ListNamespacesWithCallback invokes the schedulerx2.ListNamespaces API asynchronously
func (client *Client) ListNamespacesWithCallback(request *ListNamespacesRequest, callback func(response *ListNamespacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNamespacesResponse
		var err error
		defer close(result)
		response, err = client.ListNamespaces(request)
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

// ListNamespacesRequest is the request struct for api ListNamespaces
type ListNamespacesRequest struct {
	*requests.RpcRequest
	NamespaceName string `position:"Query" name:"NamespaceName"`
	Namespace     string `position:"Query" name:"Namespace"`
}

// ListNamespacesResponse is the response struct for api ListNamespaces
type ListNamespacesResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListNamespacesRequest creates a request to invoke ListNamespaces API
func CreateListNamespacesRequest() (request *ListNamespacesRequest) {
	request = &ListNamespacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "ListNamespaces", "", "")
	request.Method = requests.POST
	return
}

// CreateListNamespacesResponse creates a response to parse from ListNamespaces response
func CreateListNamespacesResponse() (response *ListNamespacesResponse) {
	response = &ListNamespacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
