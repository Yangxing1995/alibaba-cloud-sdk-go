package ehpc

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

// ListTasks invokes the ehpc.ListTasks API synchronously
func (client *Client) ListTasks(request *ListTasksRequest) (response *ListTasksResponse, err error) {
	response = CreateListTasksResponse()
	err = client.DoAction(request, response)
	return
}

// ListTasksWithChan invokes the ehpc.ListTasks API asynchronously
func (client *Client) ListTasksWithChan(request *ListTasksRequest) (<-chan *ListTasksResponse, <-chan error) {
	responseChan := make(chan *ListTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTasks(request)
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

// ListTasksWithCallback invokes the ehpc.ListTasks API asynchronously
func (client *Client) ListTasksWithCallback(request *ListTasksRequest, callback func(response *ListTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTasksResponse
		var err error
		defer close(result)
		response, err = client.ListTasks(request)
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

// ListTasksRequest is the request struct for api ListTasks
type ListTasksRequest struct {
	*requests.RpcRequest
	ClusterId  string           `position:"Query" name:"ClusterId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Archived   requests.Boolean `position:"Query" name:"Archived"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	TaskId     string           `position:"Query" name:"TaskId"`
}

// ListTasksResponse is the response struct for api ListTasks
type ListTasksResponse struct {
	*responses.BaseResponse
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	Tasks      []TaskInfo `json:"Tasks" xml:"Tasks"`
}

// CreateListTasksRequest creates a request to invoke ListTasks API
func CreateListTasksRequest() (request *ListTasksRequest) {
	request = &ListTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListTasks", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListTasksResponse creates a response to parse from ListTasks response
func CreateListTasksResponse() (response *ListTasksResponse) {
	response = &ListTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
