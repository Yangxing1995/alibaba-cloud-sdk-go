package live

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

// DeletePlaylistItems invokes the live.DeletePlaylistItems API synchronously
func (client *Client) DeletePlaylistItems(request *DeletePlaylistItemsRequest) (response *DeletePlaylistItemsResponse, err error) {
	response = CreateDeletePlaylistItemsResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePlaylistItemsWithChan invokes the live.DeletePlaylistItems API asynchronously
func (client *Client) DeletePlaylistItemsWithChan(request *DeletePlaylistItemsRequest) (<-chan *DeletePlaylistItemsResponse, <-chan error) {
	responseChan := make(chan *DeletePlaylistItemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePlaylistItems(request)
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

// DeletePlaylistItemsWithCallback invokes the live.DeletePlaylistItems API asynchronously
func (client *Client) DeletePlaylistItemsWithCallback(request *DeletePlaylistItemsRequest, callback func(response *DeletePlaylistItemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePlaylistItemsResponse
		var err error
		defer close(result)
		response, err = client.DeletePlaylistItems(request)
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

// DeletePlaylistItemsRequest is the request struct for api DeletePlaylistItems
type DeletePlaylistItemsRequest struct {
	*requests.RpcRequest
	ProgramItemIds string           `position:"Query" name:"ProgramItemIds"`
	ProgramId      string           `position:"Query" name:"ProgramId"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
}

// DeletePlaylistItemsResponse is the response struct for api DeletePlaylistItems
type DeletePlaylistItemsResponse struct {
	*responses.BaseResponse
	ProgramId string `json:"ProgramId" xml:"ProgramId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeletePlaylistItemsRequest creates a request to invoke DeletePlaylistItems API
func CreateDeletePlaylistItemsRequest() (request *DeletePlaylistItemsRequest) {
	request = &DeletePlaylistItemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeletePlaylistItems", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeletePlaylistItemsResponse creates a response to parse from DeletePlaylistItems response
func CreateDeletePlaylistItemsResponse() (response *DeletePlaylistItemsResponse) {
	response = &DeletePlaylistItemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
