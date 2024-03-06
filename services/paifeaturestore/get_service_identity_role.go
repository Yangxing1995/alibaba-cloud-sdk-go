package paifeaturestore

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

// GetServiceIdentityRole invokes the paifeaturestore.GetServiceIdentityRole API synchronously
func (client *Client) GetServiceIdentityRole(request *GetServiceIdentityRoleRequest) (response *GetServiceIdentityRoleResponse, err error) {
	response = CreateGetServiceIdentityRoleResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceIdentityRoleWithChan invokes the paifeaturestore.GetServiceIdentityRole API asynchronously
func (client *Client) GetServiceIdentityRoleWithChan(request *GetServiceIdentityRoleRequest) (<-chan *GetServiceIdentityRoleResponse, <-chan error) {
	responseChan := make(chan *GetServiceIdentityRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceIdentityRole(request)
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

// GetServiceIdentityRoleWithCallback invokes the paifeaturestore.GetServiceIdentityRole API asynchronously
func (client *Client) GetServiceIdentityRoleWithCallback(request *GetServiceIdentityRoleRequest, callback func(response *GetServiceIdentityRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceIdentityRoleResponse
		var err error
		defer close(result)
		response, err = client.GetServiceIdentityRole(request)
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

// GetServiceIdentityRoleRequest is the request struct for api GetServiceIdentityRole
type GetServiceIdentityRoleRequest struct {
	*requests.RoaRequest
	RoleName string `position:"Path" name:"RoleName"`
}

// GetServiceIdentityRoleResponse is the response struct for api GetServiceIdentityRole
type GetServiceIdentityRoleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	RoleName  string `json:"RoleName" xml:"RoleName"`
	Policy    string `json:"Policy" xml:"Policy"`
}

// CreateGetServiceIdentityRoleRequest creates a request to invoke GetServiceIdentityRole API
func CreateGetServiceIdentityRoleRequest() (request *GetServiceIdentityRoleRequest) {
	request = &GetServiceIdentityRoleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "GetServiceIdentityRole", "/api/v1/serviceidentityroles/[RoleName]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetServiceIdentityRoleResponse creates a response to parse from GetServiceIdentityRole response
func CreateGetServiceIdentityRoleResponse() (response *GetServiceIdentityRoleResponse) {
	response = &GetServiceIdentityRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
