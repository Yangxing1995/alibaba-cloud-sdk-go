package resourcesharing

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

// AssociateResourceSharePermission invokes the resourcesharing.AssociateResourceSharePermission API synchronously
func (client *Client) AssociateResourceSharePermission(request *AssociateResourceSharePermissionRequest) (response *AssociateResourceSharePermissionResponse, err error) {
	response = CreateAssociateResourceSharePermissionResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateResourceSharePermissionWithChan invokes the resourcesharing.AssociateResourceSharePermission API asynchronously
func (client *Client) AssociateResourceSharePermissionWithChan(request *AssociateResourceSharePermissionRequest) (<-chan *AssociateResourceSharePermissionResponse, <-chan error) {
	responseChan := make(chan *AssociateResourceSharePermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateResourceSharePermission(request)
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

// AssociateResourceSharePermissionWithCallback invokes the resourcesharing.AssociateResourceSharePermission API asynchronously
func (client *Client) AssociateResourceSharePermissionWithCallback(request *AssociateResourceSharePermissionRequest, callback func(response *AssociateResourceSharePermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateResourceSharePermissionResponse
		var err error
		defer close(result)
		response, err = client.AssociateResourceSharePermission(request)
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

// AssociateResourceSharePermissionRequest is the request struct for api AssociateResourceSharePermission
type AssociateResourceSharePermissionRequest struct {
	*requests.RpcRequest
	Replace         requests.Boolean `position:"Query" name:"Replace"`
	PermissionName  string           `position:"Query" name:"PermissionName"`
	ResourceShareId string           `position:"Query" name:"ResourceShareId"`
}

// AssociateResourceSharePermissionResponse is the response struct for api AssociateResourceSharePermission
type AssociateResourceSharePermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssociateResourceSharePermissionRequest creates a request to invoke AssociateResourceSharePermission API
func CreateAssociateResourceSharePermissionRequest() (request *AssociateResourceSharePermissionRequest) {
	request = &AssociateResourceSharePermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceSharing", "2020-01-10", "AssociateResourceSharePermission", "ressharing", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssociateResourceSharePermissionResponse creates a response to parse from AssociateResourceSharePermission response
func CreateAssociateResourceSharePermissionResponse() (response *AssociateResourceSharePermissionResponse) {
	response = &AssociateResourceSharePermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
