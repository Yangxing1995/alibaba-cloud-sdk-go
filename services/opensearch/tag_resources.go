package opensearch

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

// TagResources invokes the opensearch.TagResources API synchronously
func (client *Client) TagResources(request *TagResourcesRequest) (response *TagResourcesResponse, err error) {
	response = CreateTagResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// TagResourcesWithChan invokes the opensearch.TagResources API asynchronously
func (client *Client) TagResourcesWithChan(request *TagResourcesRequest) (<-chan *TagResourcesResponse, <-chan error) {
	responseChan := make(chan *TagResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TagResources(request)
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

// TagResourcesWithCallback invokes the opensearch.TagResources API asynchronously
func (client *Client) TagResourcesWithCallback(request *TagResourcesRequest, callback func(response *TagResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TagResourcesResponse
		var err error
		defer close(result)
		response, err = client.TagResources(request)
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

// TagResourcesRequest is the request struct for api TagResources
type TagResourcesRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// TagResourcesResponse is the response struct for api TagResources
type TagResourcesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateTagResourcesRequest creates a request to invoke TagResources API
func CreateTagResourcesRequest() (request *TagResourcesRequest) {
	request = &TagResourcesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "TagResources", "/v4/openapi/resource-tags", "", "")
	request.Method = requests.POST
	return
}

// CreateTagResourcesResponse creates a response to parse from TagResources response
func CreateTagResourcesResponse() (response *TagResourcesResponse) {
	response = &TagResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
