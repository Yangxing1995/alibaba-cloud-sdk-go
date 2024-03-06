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

// ListBindDataSources invokes the cloud_siem.ListBindDataSources API synchronously
func (client *Client) ListBindDataSources(request *ListBindDataSourcesRequest) (response *ListBindDataSourcesResponse, err error) {
	response = CreateListBindDataSourcesResponse()
	err = client.DoAction(request, response)
	return
}

// ListBindDataSourcesWithChan invokes the cloud_siem.ListBindDataSources API asynchronously
func (client *Client) ListBindDataSourcesWithChan(request *ListBindDataSourcesRequest) (<-chan *ListBindDataSourcesResponse, <-chan error) {
	responseChan := make(chan *ListBindDataSourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBindDataSources(request)
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

// ListBindDataSourcesWithCallback invokes the cloud_siem.ListBindDataSources API asynchronously
func (client *Client) ListBindDataSourcesWithCallback(request *ListBindDataSourcesRequest, callback func(response *ListBindDataSourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBindDataSourcesResponse
		var err error
		defer close(result)
		response, err = client.ListBindDataSources(request)
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

// ListBindDataSourcesRequest is the request struct for api ListBindDataSources
type ListBindDataSourcesRequest struct {
	*requests.RpcRequest
	CloudCode string `position:"Body" name:"CloudCode"`
	AccountId string `position:"Body" name:"AccountId"`
}

// ListBindDataSourcesResponse is the response struct for api ListBindDataSources
type ListBindDataSourcesResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateListBindDataSourcesRequest creates a request to invoke ListBindDataSources API
func CreateListBindDataSourcesRequest() (request *ListBindDataSourcesRequest) {
	request = &ListBindDataSourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "ListBindDataSources", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListBindDataSourcesResponse creates a response to parse from ListBindDataSources response
func CreateListBindDataSourcesResponse() (response *ListBindDataSourcesResponse) {
	response = &ListBindDataSourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
