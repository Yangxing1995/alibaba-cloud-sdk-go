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

// WithdrawAllUserGroups invokes the quickbi_public.WithdrawAllUserGroups API synchronously
func (client *Client) WithdrawAllUserGroups(request *WithdrawAllUserGroupsRequest) (response *WithdrawAllUserGroupsResponse, err error) {
	response = CreateWithdrawAllUserGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// WithdrawAllUserGroupsWithChan invokes the quickbi_public.WithdrawAllUserGroups API asynchronously
func (client *Client) WithdrawAllUserGroupsWithChan(request *WithdrawAllUserGroupsRequest) (<-chan *WithdrawAllUserGroupsResponse, <-chan error) {
	responseChan := make(chan *WithdrawAllUserGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.WithdrawAllUserGroups(request)
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

// WithdrawAllUserGroupsWithCallback invokes the quickbi_public.WithdrawAllUserGroups API asynchronously
func (client *Client) WithdrawAllUserGroupsWithCallback(request *WithdrawAllUserGroupsRequest, callback func(response *WithdrawAllUserGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *WithdrawAllUserGroupsResponse
		var err error
		defer close(result)
		response, err = client.WithdrawAllUserGroups(request)
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

// WithdrawAllUserGroupsRequest is the request struct for api WithdrawAllUserGroups
type WithdrawAllUserGroupsRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	UserId      string `position:"Query" name:"UserId"`
}

// WithdrawAllUserGroupsResponse is the response struct for api WithdrawAllUserGroups
type WithdrawAllUserGroupsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateWithdrawAllUserGroupsRequest creates a request to invoke WithdrawAllUserGroups API
func CreateWithdrawAllUserGroupsRequest() (request *WithdrawAllUserGroupsRequest) {
	request = &WithdrawAllUserGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "WithdrawAllUserGroups", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateWithdrawAllUserGroupsResponse creates a response to parse from WithdrawAllUserGroups response
func CreateWithdrawAllUserGroupsResponse() (response *WithdrawAllUserGroupsResponse) {
	response = &WithdrawAllUserGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
