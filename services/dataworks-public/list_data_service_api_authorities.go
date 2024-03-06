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

// ListDataServiceApiAuthorities invokes the dataworks_public.ListDataServiceApiAuthorities API synchronously
func (client *Client) ListDataServiceApiAuthorities(request *ListDataServiceApiAuthoritiesRequest) (response *ListDataServiceApiAuthoritiesResponse, err error) {
	response = CreateListDataServiceApiAuthoritiesResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataServiceApiAuthoritiesWithChan invokes the dataworks_public.ListDataServiceApiAuthorities API asynchronously
func (client *Client) ListDataServiceApiAuthoritiesWithChan(request *ListDataServiceApiAuthoritiesRequest) (<-chan *ListDataServiceApiAuthoritiesResponse, <-chan error) {
	responseChan := make(chan *ListDataServiceApiAuthoritiesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataServiceApiAuthorities(request)
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

// ListDataServiceApiAuthoritiesWithCallback invokes the dataworks_public.ListDataServiceApiAuthorities API asynchronously
func (client *Client) ListDataServiceApiAuthoritiesWithCallback(request *ListDataServiceApiAuthoritiesRequest, callback func(response *ListDataServiceApiAuthoritiesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataServiceApiAuthoritiesResponse
		var err error
		defer close(result)
		response, err = client.ListDataServiceApiAuthorities(request)
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

// ListDataServiceApiAuthoritiesRequest is the request struct for api ListDataServiceApiAuthorities
type ListDataServiceApiAuthoritiesRequest struct {
	*requests.RpcRequest
	ApiNameKeyword string           `position:"Body" name:"ApiNameKeyword"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	TenantId       requests.Integer `position:"Body" name:"TenantId"`
	ProjectId      requests.Integer `position:"Body" name:"ProjectId"`
}

// ListDataServiceApiAuthoritiesResponse is the response struct for api ListDataServiceApiAuthorities
type ListDataServiceApiAuthoritiesResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateListDataServiceApiAuthoritiesRequest creates a request to invoke ListDataServiceApiAuthorities API
func CreateListDataServiceApiAuthoritiesRequest() (request *ListDataServiceApiAuthoritiesRequest) {
	request = &ListDataServiceApiAuthoritiesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListDataServiceApiAuthorities", "", "")
	request.Method = requests.POST
	return
}

// CreateListDataServiceApiAuthoritiesResponse creates a response to parse from ListDataServiceApiAuthorities response
func CreateListDataServiceApiAuthoritiesResponse() (response *ListDataServiceApiAuthoritiesResponse) {
	response = &ListDataServiceApiAuthoritiesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
