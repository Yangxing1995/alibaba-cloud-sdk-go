package fnf

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

// StartSyncExecution invokes the fnf.StartSyncExecution API synchronously
func (client *Client) StartSyncExecution(request *StartSyncExecutionRequest) (response *StartSyncExecutionResponse, err error) {
	response = CreateStartSyncExecutionResponse()
	err = client.DoAction(request, response)
	return
}

// StartSyncExecutionWithChan invokes the fnf.StartSyncExecution API asynchronously
func (client *Client) StartSyncExecutionWithChan(request *StartSyncExecutionRequest) (<-chan *StartSyncExecutionResponse, <-chan error) {
	responseChan := make(chan *StartSyncExecutionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartSyncExecution(request)
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

// StartSyncExecutionWithCallback invokes the fnf.StartSyncExecution API asynchronously
func (client *Client) StartSyncExecutionWithCallback(request *StartSyncExecutionRequest, callback func(response *StartSyncExecutionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartSyncExecutionResponse
		var err error
		defer close(result)
		response, err = client.StartSyncExecution(request)
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

// StartSyncExecutionRequest is the request struct for api StartSyncExecution
type StartSyncExecutionRequest struct {
	*requests.RpcRequest
	ExecutionName string `position:"Body" name:"ExecutionName"`
	Input         string `position:"Body" name:"Input"`
	RequestId     string `position:"Query" name:"RequestId"`
	FlowName      string `position:"Body" name:"FlowName"`
}

// StartSyncExecutionResponse is the response struct for api StartSyncExecution
type StartSyncExecutionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	FlowName     string `json:"FlowName" xml:"FlowName"`
	Name         string `json:"Name" xml:"Name"`
	Status       string `json:"Status" xml:"Status"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Output       string `json:"Output" xml:"Output"`
	StartedTime  string `json:"StartedTime" xml:"StartedTime"`
	StoppedTime  string `json:"StoppedTime" xml:"StoppedTime"`
}

// CreateStartSyncExecutionRequest creates a request to invoke StartSyncExecution API
func CreateStartSyncExecutionRequest() (request *StartSyncExecutionRequest) {
	request = &StartSyncExecutionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("fnf", "2019-03-15", "StartSyncExecution", "fnf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartSyncExecutionResponse creates a response to parse from StartSyncExecution response
func CreateStartSyncExecutionResponse() (response *StartSyncExecutionResponse) {
	response = &StartSyncExecutionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
