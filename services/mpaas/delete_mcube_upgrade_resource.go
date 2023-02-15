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

// DeleteMcubeUpgradeResource invokes the mpaas.DeleteMcubeUpgradeResource API synchronously
func (client *Client) DeleteMcubeUpgradeResource(request *DeleteMcubeUpgradeResourceRequest) (response *DeleteMcubeUpgradeResourceResponse, err error) {
	response = CreateDeleteMcubeUpgradeResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMcubeUpgradeResourceWithChan invokes the mpaas.DeleteMcubeUpgradeResource API asynchronously
func (client *Client) DeleteMcubeUpgradeResourceWithChan(request *DeleteMcubeUpgradeResourceRequest) (<-chan *DeleteMcubeUpgradeResourceResponse, <-chan error) {
	responseChan := make(chan *DeleteMcubeUpgradeResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMcubeUpgradeResource(request)
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

// DeleteMcubeUpgradeResourceWithCallback invokes the mpaas.DeleteMcubeUpgradeResource API asynchronously
func (client *Client) DeleteMcubeUpgradeResourceWithCallback(request *DeleteMcubeUpgradeResourceRequest, callback func(response *DeleteMcubeUpgradeResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMcubeUpgradeResourceResponse
		var err error
		defer close(result)
		response, err = client.DeleteMcubeUpgradeResource(request)
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

// DeleteMcubeUpgradeResourceRequest is the request struct for api DeleteMcubeUpgradeResource
type DeleteMcubeUpgradeResourceRequest struct {
	*requests.RpcRequest
	Platform    string `position:"Body" name:"Platform"`
	TenantId    string `position:"Body" name:"TenantId"`
	Id          string `position:"Body" name:"Id"`
	AppId       string `position:"Body" name:"AppId"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// DeleteMcubeUpgradeResourceResponse is the response struct for api DeleteMcubeUpgradeResource
type DeleteMcubeUpgradeResourceResponse struct {
	*responses.BaseResponse
	ResultMessage string       `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string       `json:"ResultCode" xml:"ResultCode"`
	RequestId     string       `json:"RequestId" xml:"RequestId"`
	DeleteResult  DeleteResult `json:"DeleteResult" xml:"DeleteResult"`
}

// CreateDeleteMcubeUpgradeResourceRequest creates a request to invoke DeleteMcubeUpgradeResource API
func CreateDeleteMcubeUpgradeResourceRequest() (request *DeleteMcubeUpgradeResourceRequest) {
	request = &DeleteMcubeUpgradeResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "DeleteMcubeUpgradeResource", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteMcubeUpgradeResourceResponse creates a response to parse from DeleteMcubeUpgradeResource response
func CreateDeleteMcubeUpgradeResourceResponse() (response *DeleteMcubeUpgradeResourceResponse) {
	response = &DeleteMcubeUpgradeResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
