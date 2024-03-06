package quickbi_public

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

// DeleteUserFromWorkspace invokes the quickbi_public.DeleteUserFromWorkspace API synchronously
func (client *Client) DeleteUserFromWorkspace(request *DeleteUserFromWorkspaceRequest) (response *DeleteUserFromWorkspaceResponse, err error) {
	response = CreateDeleteUserFromWorkspaceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUserFromWorkspaceWithChan invokes the quickbi_public.DeleteUserFromWorkspace API asynchronously
func (client *Client) DeleteUserFromWorkspaceWithChan(request *DeleteUserFromWorkspaceRequest) (<-chan *DeleteUserFromWorkspaceResponse, <-chan error) {
	responseChan := make(chan *DeleteUserFromWorkspaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUserFromWorkspace(request)
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

// DeleteUserFromWorkspaceWithCallback invokes the quickbi_public.DeleteUserFromWorkspace API asynchronously
func (client *Client) DeleteUserFromWorkspaceWithCallback(request *DeleteUserFromWorkspaceRequest, callback func(response *DeleteUserFromWorkspaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUserFromWorkspaceResponse
		var err error
		defer close(result)
		response, err = client.DeleteUserFromWorkspace(request)
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

// DeleteUserFromWorkspaceRequest is the request struct for api DeleteUserFromWorkspace
type DeleteUserFromWorkspaceRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	UserId      string `position:"Query" name:"UserId"`
	WorkspaceId string `position:"Query" name:"WorkspaceId"`
}

// DeleteUserFromWorkspaceResponse is the response struct for api DeleteUserFromWorkspace
type DeleteUserFromWorkspaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteUserFromWorkspaceRequest creates a request to invoke DeleteUserFromWorkspace API
func CreateDeleteUserFromWorkspaceRequest() (request *DeleteUserFromWorkspaceRequest) {
	request = &DeleteUserFromWorkspaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "DeleteUserFromWorkspace", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteUserFromWorkspaceResponse creates a response to parse from DeleteUserFromWorkspace response
func CreateDeleteUserFromWorkspaceResponse() (response *DeleteUserFromWorkspaceResponse) {
	response = &DeleteUserFromWorkspaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
