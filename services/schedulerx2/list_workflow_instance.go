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

// ListWorkflowInstance invokes the schedulerx2.ListWorkflowInstance API synchronously
func (client *Client) ListWorkflowInstance(request *ListWorkflowInstanceRequest) (response *ListWorkflowInstanceResponse, err error) {
	response = CreateListWorkflowInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ListWorkflowInstanceWithChan invokes the schedulerx2.ListWorkflowInstance API asynchronously
func (client *Client) ListWorkflowInstanceWithChan(request *ListWorkflowInstanceRequest) (<-chan *ListWorkflowInstanceResponse, <-chan error) {
	responseChan := make(chan *ListWorkflowInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListWorkflowInstance(request)
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

// ListWorkflowInstanceWithCallback invokes the schedulerx2.ListWorkflowInstance API asynchronously
func (client *Client) ListWorkflowInstanceWithCallback(request *ListWorkflowInstanceRequest, callback func(response *ListWorkflowInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListWorkflowInstanceResponse
		var err error
		defer close(result)
		response, err = client.ListWorkflowInstance(request)
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

// ListWorkflowInstanceRequest is the request struct for api ListWorkflowInstance
type ListWorkflowInstanceRequest struct {
	*requests.RpcRequest
	NamespaceSource string `position:"Query" name:"NamespaceSource"`
	GroupId         string `position:"Query" name:"GroupId"`
	Namespace       string `position:"Query" name:"Namespace"`
	WorkflowId      string `position:"Query" name:"WorkflowId"`
}

// ListWorkflowInstanceResponse is the response struct for api ListWorkflowInstance
type ListWorkflowInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListWorkflowInstanceRequest creates a request to invoke ListWorkflowInstance API
func CreateListWorkflowInstanceRequest() (request *ListWorkflowInstanceRequest) {
	request = &ListWorkflowInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "ListWorkflowInstance", "", "")
	request.Method = requests.GET
	return
}

// CreateListWorkflowInstanceResponse creates a response to parse from ListWorkflowInstance response
func CreateListWorkflowInstanceResponse() (response *ListWorkflowInstanceResponse) {
	response = &ListWorkflowInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
