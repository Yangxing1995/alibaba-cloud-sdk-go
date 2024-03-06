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

// ChangeMcubeMiniTaskStatus invokes the mpaas.ChangeMcubeMiniTaskStatus API synchronously
func (client *Client) ChangeMcubeMiniTaskStatus(request *ChangeMcubeMiniTaskStatusRequest) (response *ChangeMcubeMiniTaskStatusResponse, err error) {
	response = CreateChangeMcubeMiniTaskStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeMcubeMiniTaskStatusWithChan invokes the mpaas.ChangeMcubeMiniTaskStatus API asynchronously
func (client *Client) ChangeMcubeMiniTaskStatusWithChan(request *ChangeMcubeMiniTaskStatusRequest) (<-chan *ChangeMcubeMiniTaskStatusResponse, <-chan error) {
	responseChan := make(chan *ChangeMcubeMiniTaskStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeMcubeMiniTaskStatus(request)
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

// ChangeMcubeMiniTaskStatusWithCallback invokes the mpaas.ChangeMcubeMiniTaskStatus API asynchronously
func (client *Client) ChangeMcubeMiniTaskStatusWithCallback(request *ChangeMcubeMiniTaskStatusRequest, callback func(response *ChangeMcubeMiniTaskStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeMcubeMiniTaskStatusResponse
		var err error
		defer close(result)
		response, err = client.ChangeMcubeMiniTaskStatus(request)
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

// ChangeMcubeMiniTaskStatusRequest is the request struct for api ChangeMcubeMiniTaskStatus
type ChangeMcubeMiniTaskStatusRequest struct {
	*requests.RpcRequest
	PackageId   requests.Integer `position:"Body" name:"PackageId"`
	TaskStatus  requests.Integer `position:"Body" name:"TaskStatus"`
	TenantId    string           `position:"Body" name:"TenantId"`
	TaskId      requests.Integer `position:"Body" name:"TaskId"`
	BizType     string           `position:"Body" name:"BizType"`
	AppId       string           `position:"Body" name:"AppId"`
	WorkspaceId string           `position:"Body" name:"WorkspaceId"`
}

// ChangeMcubeMiniTaskStatusResponse is the response struct for api ChangeMcubeMiniTaskStatus
type ChangeMcubeMiniTaskStatusResponse struct {
	*responses.BaseResponse
	ResultMessage              string                     `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode                 string                     `json:"ResultCode" xml:"ResultCode"`
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	ChangeMiniTaskStatusResult ChangeMiniTaskStatusResult `json:"ChangeMiniTaskStatusResult" xml:"ChangeMiniTaskStatusResult"`
}

// CreateChangeMcubeMiniTaskStatusRequest creates a request to invoke ChangeMcubeMiniTaskStatus API
func CreateChangeMcubeMiniTaskStatusRequest() (request *ChangeMcubeMiniTaskStatusRequest) {
	request = &ChangeMcubeMiniTaskStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "ChangeMcubeMiniTaskStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateChangeMcubeMiniTaskStatusResponse creates a response to parse from ChangeMcubeMiniTaskStatus response
func CreateChangeMcubeMiniTaskStatusResponse() (response *ChangeMcubeMiniTaskStatusResponse) {
	response = &ChangeMcubeMiniTaskStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
