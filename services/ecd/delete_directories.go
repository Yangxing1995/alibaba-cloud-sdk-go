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

// DeleteDirectories invokes the ecd.DeleteDirectories API synchronously
func (client *Client) DeleteDirectories(request *DeleteDirectoriesRequest) (response *DeleteDirectoriesResponse, err error) {
	response = CreateDeleteDirectoriesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDirectoriesWithChan invokes the ecd.DeleteDirectories API asynchronously
func (client *Client) DeleteDirectoriesWithChan(request *DeleteDirectoriesRequest) (<-chan *DeleteDirectoriesResponse, <-chan error) {
	responseChan := make(chan *DeleteDirectoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDirectories(request)
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

// DeleteDirectoriesWithCallback invokes the ecd.DeleteDirectories API asynchronously
func (client *Client) DeleteDirectoriesWithCallback(request *DeleteDirectoriesRequest, callback func(response *DeleteDirectoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDirectoriesResponse
		var err error
		defer close(result)
		response, err = client.DeleteDirectories(request)
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

// DeleteDirectoriesRequest is the request struct for api DeleteDirectories
type DeleteDirectoriesRequest struct {
	*requests.RpcRequest
	DirectoryId *[]string `position:"Query" name:"DirectoryId"  type:"Repeated"`
}

// DeleteDirectoriesResponse is the response struct for api DeleteDirectories
type DeleteDirectoriesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDirectoriesRequest creates a request to invoke DeleteDirectories API
func CreateDeleteDirectoriesRequest() (request *DeleteDirectoriesRequest) {
	request = &DeleteDirectoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DeleteDirectories", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteDirectoriesResponse creates a response to parse from DeleteDirectories response
func CreateDeleteDirectoriesResponse() (response *DeleteDirectoriesResponse) {
	response = &DeleteDirectoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
