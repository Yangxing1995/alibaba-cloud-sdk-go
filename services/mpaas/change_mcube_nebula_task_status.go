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

// ChangeMcubeNebulaTaskStatus invokes the mpaas.ChangeMcubeNebulaTaskStatus API synchronously
func (client *Client) ChangeMcubeNebulaTaskStatus(request *ChangeMcubeNebulaTaskStatusRequest) (response *ChangeMcubeNebulaTaskStatusResponse, err error) {
	response = CreateChangeMcubeNebulaTaskStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeMcubeNebulaTaskStatusWithChan invokes the mpaas.ChangeMcubeNebulaTaskStatus API asynchronously
func (client *Client) ChangeMcubeNebulaTaskStatusWithChan(request *ChangeMcubeNebulaTaskStatusRequest) (<-chan *ChangeMcubeNebulaTaskStatusResponse, <-chan error) {
	responseChan := make(chan *ChangeMcubeNebulaTaskStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeMcubeNebulaTaskStatus(request)
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

// ChangeMcubeNebulaTaskStatusWithCallback invokes the mpaas.ChangeMcubeNebulaTaskStatus API asynchronously
func (client *Client) ChangeMcubeNebulaTaskStatusWithCallback(request *ChangeMcubeNebulaTaskStatusRequest, callback func(response *ChangeMcubeNebulaTaskStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeMcubeNebulaTaskStatusResponse
		var err error
		defer close(result)
		response, err = client.ChangeMcubeNebulaTaskStatus(request)
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

// ChangeMcubeNebulaTaskStatusRequest is the request struct for api ChangeMcubeNebulaTaskStatus
type ChangeMcubeNebulaTaskStatusRequest struct {
	*requests.RpcRequest
	PackageId   string           `position:"Body" name:"PackageId"`
	TaskStatus  requests.Integer `position:"Body" name:"TaskStatus"`
	TenantId    string           `position:"Body" name:"TenantId"`
	TaskId      string           `position:"Body" name:"TaskId"`
	BizType     string           `position:"Body" name:"BizType"`
	AppId       string           `position:"Body" name:"AppId"`
	WorkspaceId string           `position:"Body" name:"WorkspaceId"`
}

// ChangeMcubeNebulaTaskStatusResponse is the response struct for api ChangeMcubeNebulaTaskStatus
type ChangeMcubeNebulaTaskStatusResponse struct {
	*responses.BaseResponse
	ResultMessage                     string                            `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode                        string                            `json:"ResultCode" xml:"ResultCode"`
	RequestId                         string                            `json:"RequestId" xml:"RequestId"`
	ChangeMcubeNebulaTaskStatusResult ChangeMcubeNebulaTaskStatusResult `json:"ChangeMcubeNebulaTaskStatusResult" xml:"ChangeMcubeNebulaTaskStatusResult"`
}

// CreateChangeMcubeNebulaTaskStatusRequest creates a request to invoke ChangeMcubeNebulaTaskStatus API
func CreateChangeMcubeNebulaTaskStatusRequest() (request *ChangeMcubeNebulaTaskStatusRequest) {
	request = &ChangeMcubeNebulaTaskStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "ChangeMcubeNebulaTaskStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateChangeMcubeNebulaTaskStatusResponse creates a response to parse from ChangeMcubeNebulaTaskStatus response
func CreateChangeMcubeNebulaTaskStatusResponse() (response *ChangeMcubeNebulaTaskStatusResponse) {
	response = &ChangeMcubeNebulaTaskStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
