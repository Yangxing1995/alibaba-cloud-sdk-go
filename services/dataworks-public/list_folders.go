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

// ListFolders invokes the dataworks_public.ListFolders API synchronously
func (client *Client) ListFolders(request *ListFoldersRequest) (response *ListFoldersResponse, err error) {
	response = CreateListFoldersResponse()
	err = client.DoAction(request, response)
	return
}

// ListFoldersWithChan invokes the dataworks_public.ListFolders API asynchronously
func (client *Client) ListFoldersWithChan(request *ListFoldersRequest) (<-chan *ListFoldersResponse, <-chan error) {
	responseChan := make(chan *ListFoldersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFolders(request)
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

// ListFoldersWithCallback invokes the dataworks_public.ListFolders API asynchronously
func (client *Client) ListFoldersWithCallback(request *ListFoldersRequest, callback func(response *ListFoldersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFoldersResponse
		var err error
		defer close(result)
		response, err = client.ListFolders(request)
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

// ListFoldersRequest is the request struct for api ListFolders
type ListFoldersRequest struct {
	*requests.RpcRequest
	PageSize          requests.Integer `position:"Body" name:"PageSize"`
	ParentFolderPath  string           `position:"Body" name:"ParentFolderPath"`
	ProjectId         requests.Integer `position:"Body" name:"ProjectId"`
	ProjectIdentifier string           `position:"Body" name:"ProjectIdentifier"`
	PageNumber        requests.Integer `position:"Body" name:"PageNumber"`
}

// ListFoldersResponse is the response struct for api ListFolders
type ListFoldersResponse struct {
	*responses.BaseResponse
	RequestId      string            `json:"RequestId" xml:"RequestId"`
	Success        bool              `json:"Success" xml:"Success"`
	ErrorCode      string            `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string            `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInListFolders `json:"Data" xml:"Data"`
}

// CreateListFoldersRequest creates a request to invoke ListFolders API
func CreateListFoldersRequest() (request *ListFoldersRequest) {
	request = &ListFoldersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListFolders", "", "")
	request.Method = requests.POST
	return
}

// CreateListFoldersResponse creates a response to parse from ListFolders response
func CreateListFoldersResponse() (response *ListFoldersResponse) {
	response = &ListFoldersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
