package cloud_siem

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

// ListAllProds invokes the cloud_siem.ListAllProds API synchronously
func (client *Client) ListAllProds(request *ListAllProdsRequest) (response *ListAllProdsResponse, err error) {
	response = CreateListAllProdsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAllProdsWithChan invokes the cloud_siem.ListAllProds API asynchronously
func (client *Client) ListAllProdsWithChan(request *ListAllProdsRequest) (<-chan *ListAllProdsResponse, <-chan error) {
	responseChan := make(chan *ListAllProdsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAllProds(request)
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

// ListAllProdsWithCallback invokes the cloud_siem.ListAllProds API asynchronously
func (client *Client) ListAllProdsWithCallback(request *ListAllProdsRequest, callback func(response *ListAllProdsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAllProdsResponse
		var err error
		defer close(result)
		response, err = client.ListAllProds(request)
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

// ListAllProdsRequest is the request struct for api ListAllProds
type ListAllProdsRequest struct {
	*requests.RpcRequest
}

// ListAllProdsResponse is the response struct for api ListAllProds
type ListAllProdsResponse struct {
	*responses.BaseResponse
	RequestId string             `json:"RequestId" xml:"RequestId"`
	Data      DataInListAllProds `json:"Data" xml:"Data"`
}

// CreateListAllProdsRequest creates a request to invoke ListAllProds API
func CreateListAllProdsRequest() (request *ListAllProdsRequest) {
	request = &ListAllProdsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "ListAllProds", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAllProdsResponse creates a response to parse from ListAllProds response
func CreateListAllProdsResponse() (response *ListAllProdsResponse) {
	response = &ListAllProdsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
