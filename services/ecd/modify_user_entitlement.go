package ecd

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

// ModifyUserEntitlement invokes the ecd.ModifyUserEntitlement API synchronously
func (client *Client) ModifyUserEntitlement(request *ModifyUserEntitlementRequest) (response *ModifyUserEntitlementResponse, err error) {
	response = CreateModifyUserEntitlementResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyUserEntitlementWithChan invokes the ecd.ModifyUserEntitlement API asynchronously
func (client *Client) ModifyUserEntitlementWithChan(request *ModifyUserEntitlementRequest) (<-chan *ModifyUserEntitlementResponse, <-chan error) {
	responseChan := make(chan *ModifyUserEntitlementResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyUserEntitlement(request)
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

// ModifyUserEntitlementWithCallback invokes the ecd.ModifyUserEntitlement API asynchronously
func (client *Client) ModifyUserEntitlementWithCallback(request *ModifyUserEntitlementRequest, callback func(response *ModifyUserEntitlementResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyUserEntitlementResponse
		var err error
		defer close(result)
		response, err = client.ModifyUserEntitlement(request)
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

// ModifyUserEntitlementRequest is the request struct for api ModifyUserEntitlement
type ModifyUserEntitlementRequest struct {
	*requests.RpcRequest
	AuthorizeDesktopId *[]string `position:"Query" name:"AuthorizeDesktopId"  type:"Repeated"`
	RevokeDesktopId    *[]string `position:"Query" name:"RevokeDesktopId"  type:"Repeated"`
	EndUserId          *[]string `position:"Query" name:"EndUserId"  type:"Repeated"`
}

// ModifyUserEntitlementResponse is the response struct for api ModifyUserEntitlement
type ModifyUserEntitlementResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyUserEntitlementRequest creates a request to invoke ModifyUserEntitlement API
func CreateModifyUserEntitlementRequest() (request *ModifyUserEntitlementRequest) {
	request = &ModifyUserEntitlementRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ModifyUserEntitlement", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyUserEntitlementResponse creates a response to parse from ModifyUserEntitlement response
func CreateModifyUserEntitlementResponse() (response *ModifyUserEntitlementResponse) {
	response = &ModifyUserEntitlementResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
