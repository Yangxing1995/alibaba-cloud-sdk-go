package mpaas

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

// RevokePushTask invokes the mpaas.RevokePushTask API synchronously
func (client *Client) RevokePushTask(request *RevokePushTaskRequest) (response *RevokePushTaskResponse, err error) {
	response = CreateRevokePushTaskResponse()
	err = client.DoAction(request, response)
	return
}

// RevokePushTaskWithChan invokes the mpaas.RevokePushTask API asynchronously
func (client *Client) RevokePushTaskWithChan(request *RevokePushTaskRequest) (<-chan *RevokePushTaskResponse, <-chan error) {
	responseChan := make(chan *RevokePushTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokePushTask(request)
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

// RevokePushTaskWithCallback invokes the mpaas.RevokePushTask API asynchronously
func (client *Client) RevokePushTaskWithCallback(request *RevokePushTaskRequest, callback func(response *RevokePushTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokePushTaskResponse
		var err error
		defer close(result)
		response, err = client.RevokePushTask(request)
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

// RevokePushTaskRequest is the request struct for api RevokePushTask
type RevokePushTaskRequest struct {
	*requests.RpcRequest
	TaskId      string `position:"Body" name:"TaskId"`
	AppId       string `position:"Body" name:"AppId"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// RevokePushTaskResponse is the response struct for api RevokePushTask
type RevokePushTaskResponse struct {
	*responses.BaseResponse
	ResultMessage string     `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string     `json:"ResultCode" xml:"ResultCode"`
	RequestId     string     `json:"RequestId" xml:"RequestId"`
	PushResult    PushResult `json:"PushResult" xml:"PushResult"`
}

// CreateRevokePushTaskRequest creates a request to invoke RevokePushTask API
func CreateRevokePushTaskRequest() (request *RevokePushTaskRequest) {
	request = &RevokePushTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "RevokePushTask", "", "")
	request.Method = requests.POST
	return
}

// CreateRevokePushTaskResponse creates a response to parse from RevokePushTask response
func CreateRevokePushTaskResponse() (response *RevokePushTaskResponse) {
	response = &RevokePushTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
