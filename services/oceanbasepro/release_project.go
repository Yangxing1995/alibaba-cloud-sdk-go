package oceanbasepro

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

// ReleaseProject invokes the oceanbasepro.ReleaseProject API synchronously
func (client *Client) ReleaseProject(request *ReleaseProjectRequest) (response *ReleaseProjectResponse, err error) {
	response = CreateReleaseProjectResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseProjectWithChan invokes the oceanbasepro.ReleaseProject API asynchronously
func (client *Client) ReleaseProjectWithChan(request *ReleaseProjectRequest) (<-chan *ReleaseProjectResponse, <-chan error) {
	responseChan := make(chan *ReleaseProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseProject(request)
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

// ReleaseProjectWithCallback invokes the oceanbasepro.ReleaseProject API asynchronously
func (client *Client) ReleaseProjectWithCallback(request *ReleaseProjectRequest, callback func(response *ReleaseProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseProjectResponse
		var err error
		defer close(result)
		response, err = client.ReleaseProject(request)
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

// ReleaseProjectRequest is the request struct for api ReleaseProject
type ReleaseProjectRequest struct {
	*requests.RpcRequest
	Id string `position:"Body" name:"Id"`
}

// ReleaseProjectResponse is the response struct for api ReleaseProject
type ReleaseProjectResponse struct {
	*responses.BaseResponse
}

// CreateReleaseProjectRequest creates a request to invoke ReleaseProject API
func CreateReleaseProjectRequest() (request *ReleaseProjectRequest) {
	request = &ReleaseProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "ReleaseProject", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReleaseProjectResponse creates a response to parse from ReleaseProject response
func CreateReleaseProjectResponse() (response *ReleaseProjectResponse) {
	response = &ReleaseProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
