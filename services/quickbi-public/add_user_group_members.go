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

// AddUserGroupMembers invokes the quickbi_public.AddUserGroupMembers API synchronously
func (client *Client) AddUserGroupMembers(request *AddUserGroupMembersRequest) (response *AddUserGroupMembersResponse, err error) {
	response = CreateAddUserGroupMembersResponse()
	err = client.DoAction(request, response)
	return
}

// AddUserGroupMembersWithChan invokes the quickbi_public.AddUserGroupMembers API asynchronously
func (client *Client) AddUserGroupMembersWithChan(request *AddUserGroupMembersRequest) (<-chan *AddUserGroupMembersResponse, <-chan error) {
	responseChan := make(chan *AddUserGroupMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddUserGroupMembers(request)
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

// AddUserGroupMembersWithCallback invokes the quickbi_public.AddUserGroupMembers API asynchronously
func (client *Client) AddUserGroupMembersWithCallback(request *AddUserGroupMembersRequest, callback func(response *AddUserGroupMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddUserGroupMembersResponse
		var err error
		defer close(result)
		response, err = client.AddUserGroupMembers(request)
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

// AddUserGroupMembersRequest is the request struct for api AddUserGroupMembers
type AddUserGroupMembersRequest struct {
	*requests.RpcRequest
	AccessPoint  string `position:"Query" name:"AccessPoint"`
	SignType     string `position:"Query" name:"SignType"`
	UserGroupIds string `position:"Query" name:"UserGroupIds"`
	UserId       string `position:"Query" name:"UserId"`
}

// AddUserGroupMembersResponse is the response struct for api AddUserGroupMembers
type AddUserGroupMembersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateAddUserGroupMembersRequest creates a request to invoke AddUserGroupMembers API
func CreateAddUserGroupMembersRequest() (request *AddUserGroupMembersRequest) {
	request = &AddUserGroupMembersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "AddUserGroupMembers", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddUserGroupMembersResponse creates a response to parse from AddUserGroupMembers response
func CreateAddUserGroupMembersResponse() (response *AddUserGroupMembersResponse) {
	response = &AddUserGroupMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
