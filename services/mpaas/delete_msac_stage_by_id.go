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

// DeleteMsacStageById invokes the mpaas.DeleteMsacStageById API synchronously
func (client *Client) DeleteMsacStageById(request *DeleteMsacStageByIdRequest) (response *DeleteMsacStageByIdResponse, err error) {
	response = CreateDeleteMsacStageByIdResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMsacStageByIdWithChan invokes the mpaas.DeleteMsacStageById API asynchronously
func (client *Client) DeleteMsacStageByIdWithChan(request *DeleteMsacStageByIdRequest) (<-chan *DeleteMsacStageByIdResponse, <-chan error) {
	responseChan := make(chan *DeleteMsacStageByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMsacStageById(request)
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

// DeleteMsacStageByIdWithCallback invokes the mpaas.DeleteMsacStageById API asynchronously
func (client *Client) DeleteMsacStageByIdWithCallback(request *DeleteMsacStageByIdRequest, callback func(response *DeleteMsacStageByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMsacStageByIdResponse
		var err error
		defer close(result)
		response, err = client.DeleteMsacStageById(request)
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

// DeleteMsacStageByIdRequest is the request struct for api DeleteMsacStageById
type DeleteMsacStageByIdRequest struct {
	*requests.RpcRequest
	MpaasMappcenterMsacDeleteStageJsonStr string `position:"Body" name:"MpaasMappcenterMsacDeleteStageJsonStr"`
	TenantId                              string `position:"Body" name:"TenantId"`
	AppId                                 string `position:"Body" name:"AppId"`
	WorkspaceId                           string `position:"Body" name:"WorkspaceId"`
}

// DeleteMsacStageByIdResponse is the response struct for api DeleteMsacStageById
type DeleteMsacStageByIdResponse struct {
	*responses.BaseResponse
	ResultMessage string                             `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string                             `json:"ResultCode" xml:"ResultCode"`
	RequestId     string                             `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContentInDeleteMsacStageById `json:"ResultContent" xml:"ResultContent"`
}

// CreateDeleteMsacStageByIdRequest creates a request to invoke DeleteMsacStageById API
func CreateDeleteMsacStageByIdRequest() (request *DeleteMsacStageByIdRequest) {
	request = &DeleteMsacStageByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "DeleteMsacStageById", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteMsacStageByIdResponse creates a response to parse from DeleteMsacStageById response
func CreateDeleteMsacStageByIdResponse() (response *DeleteMsacStageByIdResponse) {
	response = &DeleteMsacStageByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
