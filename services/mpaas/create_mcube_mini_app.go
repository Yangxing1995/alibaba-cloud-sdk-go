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

// CreateMcubeMiniApp invokes the mpaas.CreateMcubeMiniApp API synchronously
func (client *Client) CreateMcubeMiniApp(request *CreateMcubeMiniAppRequest) (response *CreateMcubeMiniAppResponse, err error) {
	response = CreateCreateMcubeMiniAppResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMcubeMiniAppWithChan invokes the mpaas.CreateMcubeMiniApp API asynchronously
func (client *Client) CreateMcubeMiniAppWithChan(request *CreateMcubeMiniAppRequest) (<-chan *CreateMcubeMiniAppResponse, <-chan error) {
	responseChan := make(chan *CreateMcubeMiniAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMcubeMiniApp(request)
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

// CreateMcubeMiniAppWithCallback invokes the mpaas.CreateMcubeMiniApp API asynchronously
func (client *Client) CreateMcubeMiniAppWithCallback(request *CreateMcubeMiniAppRequest, callback func(response *CreateMcubeMiniAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMcubeMiniAppResponse
		var err error
		defer close(result)
		response, err = client.CreateMcubeMiniApp(request)
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

// CreateMcubeMiniAppRequest is the request struct for api CreateMcubeMiniApp
type CreateMcubeMiniAppRequest struct {
	*requests.RpcRequest
	H5Name      string `position:"Body" name:"H5Name"`
	H5Id        string `position:"Body" name:"H5Id"`
	TenantId    string `position:"Body" name:"TenantId"`
	AppId       string `position:"Body" name:"AppId"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// CreateMcubeMiniAppResponse is the response struct for api CreateMcubeMiniApp
type CreateMcubeMiniAppResponse struct {
	*responses.BaseResponse
	ResultMessage    string           `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode       string           `json:"ResultCode" xml:"ResultCode"`
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	CreateMiniResult CreateMiniResult `json:"CreateMiniResult" xml:"CreateMiniResult"`
}

// CreateCreateMcubeMiniAppRequest creates a request to invoke CreateMcubeMiniApp API
func CreateCreateMcubeMiniAppRequest() (request *CreateMcubeMiniAppRequest) {
	request = &CreateMcubeMiniAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "CreateMcubeMiniApp", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateMcubeMiniAppResponse creates a response to parse from CreateMcubeMiniApp response
func CreateCreateMcubeMiniAppResponse() (response *CreateMcubeMiniAppResponse) {
	response = &CreateMcubeMiniAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
