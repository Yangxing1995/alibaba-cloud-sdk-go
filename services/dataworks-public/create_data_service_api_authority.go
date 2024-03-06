package dataworks_public

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

// CreateDataServiceApiAuthority invokes the dataworks_public.CreateDataServiceApiAuthority API synchronously
func (client *Client) CreateDataServiceApiAuthority(request *CreateDataServiceApiAuthorityRequest) (response *CreateDataServiceApiAuthorityResponse, err error) {
	response = CreateCreateDataServiceApiAuthorityResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataServiceApiAuthorityWithChan invokes the dataworks_public.CreateDataServiceApiAuthority API asynchronously
func (client *Client) CreateDataServiceApiAuthorityWithChan(request *CreateDataServiceApiAuthorityRequest) (<-chan *CreateDataServiceApiAuthorityResponse, <-chan error) {
	responseChan := make(chan *CreateDataServiceApiAuthorityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataServiceApiAuthority(request)
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

// CreateDataServiceApiAuthorityWithCallback invokes the dataworks_public.CreateDataServiceApiAuthority API asynchronously
func (client *Client) CreateDataServiceApiAuthorityWithCallback(request *CreateDataServiceApiAuthorityRequest, callback func(response *CreateDataServiceApiAuthorityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataServiceApiAuthorityResponse
		var err error
		defer close(result)
		response, err = client.CreateDataServiceApiAuthority(request)
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

// CreateDataServiceApiAuthorityRequest is the request struct for api CreateDataServiceApiAuthority
type CreateDataServiceApiAuthorityRequest struct {
	*requests.RpcRequest
	AuthorizedProjectId requests.Integer `position:"Body" name:"AuthorizedProjectId"`
	EndTime             requests.Integer `position:"Body" name:"EndTime"`
	TenantId            requests.Integer `position:"Body" name:"TenantId"`
	ProjectId           requests.Integer `position:"Body" name:"ProjectId"`
	ApiId               requests.Integer `position:"Body" name:"ApiId"`
}

// CreateDataServiceApiAuthorityResponse is the response struct for api CreateDataServiceApiAuthority
type CreateDataServiceApiAuthorityResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDataServiceApiAuthorityRequest creates a request to invoke CreateDataServiceApiAuthority API
func CreateCreateDataServiceApiAuthorityRequest() (request *CreateDataServiceApiAuthorityRequest) {
	request = &CreateDataServiceApiAuthorityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateDataServiceApiAuthority", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDataServiceApiAuthorityResponse creates a response to parse from CreateDataServiceApiAuthority response
func CreateCreateDataServiceApiAuthorityResponse() (response *CreateDataServiceApiAuthorityResponse) {
	response = &CreateDataServiceApiAuthorityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
