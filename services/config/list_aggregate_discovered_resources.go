package config

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

// ListAggregateDiscoveredResources invokes the config.ListAggregateDiscoveredResources API synchronously
func (client *Client) ListAggregateDiscoveredResources(request *ListAggregateDiscoveredResourcesRequest) (response *ListAggregateDiscoveredResourcesResponse, err error) {
	response = CreateListAggregateDiscoveredResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// ListAggregateDiscoveredResourcesWithChan invokes the config.ListAggregateDiscoveredResources API asynchronously
func (client *Client) ListAggregateDiscoveredResourcesWithChan(request *ListAggregateDiscoveredResourcesRequest) (<-chan *ListAggregateDiscoveredResourcesResponse, <-chan error) {
	responseChan := make(chan *ListAggregateDiscoveredResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAggregateDiscoveredResources(request)
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

// ListAggregateDiscoveredResourcesWithCallback invokes the config.ListAggregateDiscoveredResources API asynchronously
func (client *Client) ListAggregateDiscoveredResourcesWithCallback(request *ListAggregateDiscoveredResourcesRequest, callback func(response *ListAggregateDiscoveredResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAggregateDiscoveredResourcesResponse
		var err error
		defer close(result)
		response, err = client.ListAggregateDiscoveredResources(request)
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

// ListAggregateDiscoveredResourcesRequest is the request struct for api ListAggregateDiscoveredResources
type ListAggregateDiscoveredResourcesRequest struct {
	*requests.RpcRequest
	ResourceDeleted requests.Integer `position:"Query" name:"ResourceDeleted"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Regions         string           `position:"Query" name:"Regions"`
	AggregatorId    string           `position:"Query" name:"AggregatorId"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	FolderId        string           `position:"Query" name:"FolderId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	ComplianceType  string           `position:"Query" name:"ComplianceType"`
	ResourceId      string           `position:"Query" name:"ResourceId"`
	ResourceTypes   string           `position:"Query" name:"ResourceTypes"`
}

// ListAggregateDiscoveredResourcesResponse is the response struct for api ListAggregateDiscoveredResources
type ListAggregateDiscoveredResourcesResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	DiscoveredResourceProfiles DiscoveredResourceProfiles `json:"DiscoveredResourceProfiles" xml:"DiscoveredResourceProfiles"`
}

// CreateListAggregateDiscoveredResourcesRequest creates a request to invoke ListAggregateDiscoveredResources API
func CreateListAggregateDiscoveredResourcesRequest() (request *ListAggregateDiscoveredResourcesRequest) {
	request = &ListAggregateDiscoveredResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "ListAggregateDiscoveredResources", "", "")
	request.Method = requests.POST
	return
}

// CreateListAggregateDiscoveredResourcesResponse creates a response to parse from ListAggregateDiscoveredResources response
func CreateListAggregateDiscoveredResourcesResponse() (response *ListAggregateDiscoveredResourcesResponse) {
	response = &ListAggregateDiscoveredResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
