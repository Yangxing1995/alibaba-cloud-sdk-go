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

// ModifyUserGroups invokes the ehpc.ModifyUserGroups API synchronously
func (client *Client) ModifyUserGroups(request *ModifyUserGroupsRequest) (response *ModifyUserGroupsResponse, err error) {
	response = CreateModifyUserGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyUserGroupsWithChan invokes the ehpc.ModifyUserGroups API asynchronously
func (client *Client) ModifyUserGroupsWithChan(request *ModifyUserGroupsRequest) (<-chan *ModifyUserGroupsResponse, <-chan error) {
	responseChan := make(chan *ModifyUserGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyUserGroups(request)
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

// ModifyUserGroupsWithCallback invokes the ehpc.ModifyUserGroups API asynchronously
func (client *Client) ModifyUserGroupsWithCallback(request *ModifyUserGroupsRequest, callback func(response *ModifyUserGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyUserGroupsResponse
		var err error
		defer close(result)
		response, err = client.ModifyUserGroups(request)
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

// ModifyUserGroupsRequest is the request struct for api ModifyUserGroups
type ModifyUserGroupsRequest struct {
	*requests.RpcRequest
	ClusterId string                  `position:"Query" name:"ClusterId"`
	Async     requests.Boolean        `position:"Query" name:"Async"`
	User      *[]ModifyUserGroupsUser `position:"Query" name:"User"  type:"Repeated"`
}

// ModifyUserGroupsUser is a repeated param struct in ModifyUserGroupsRequest
type ModifyUserGroupsUser struct {
	Name  string `name:"Name"`
	Group string `name:"Group"`
}

// ModifyUserGroupsResponse is the response struct for api ModifyUserGroups
type ModifyUserGroupsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyUserGroupsRequest creates a request to invoke ModifyUserGroups API
func CreateModifyUserGroupsRequest() (request *ModifyUserGroupsRequest) {
	request = &ModifyUserGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ModifyUserGroups", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateModifyUserGroupsResponse creates a response to parse from ModifyUserGroups response
func CreateModifyUserGroupsResponse() (response *ModifyUserGroupsResponse) {
	response = &ModifyUserGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
