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

// ListDataCollections invokes the opensearch.ListDataCollections API synchronously
func (client *Client) ListDataCollections(request *ListDataCollectionsRequest) (response *ListDataCollectionsResponse, err error) {
	response = CreateListDataCollectionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataCollectionsWithChan invokes the opensearch.ListDataCollections API asynchronously
func (client *Client) ListDataCollectionsWithChan(request *ListDataCollectionsRequest) (<-chan *ListDataCollectionsResponse, <-chan error) {
	responseChan := make(chan *ListDataCollectionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataCollections(request)
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

// ListDataCollectionsWithCallback invokes the opensearch.ListDataCollections API asynchronously
func (client *Client) ListDataCollectionsWithCallback(request *ListDataCollectionsRequest, callback func(response *ListDataCollectionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataCollectionsResponse
		var err error
		defer close(result)
		response, err = client.ListDataCollections(request)
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

// ListDataCollectionsRequest is the request struct for api ListDataCollections
type ListDataCollectionsRequest struct {
	*requests.RoaRequest
	PageSize         requests.Integer `position:"Query" name:"pageSize"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
	PageNumber       requests.Integer `position:"Query" name:"pageNumber"`
}

// ListDataCollectionsResponse is the response struct for api ListDataCollections
type ListDataCollectionsResponse struct {
	*responses.BaseResponse
	TotalCount int          `json:"totalCount" xml:"totalCount"`
	RequestId  string       `json:"requestId" xml:"requestId"`
	Result     []ResultItem `json:"result" xml:"result"`
}

// CreateListDataCollectionsRequest creates a request to invoke ListDataCollections API
func CreateListDataCollectionsRequest() (request *ListDataCollectionsRequest) {
	request = &ListDataCollectionsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListDataCollections", "/v4/openapi/app-groups/[appGroupIdentity]/data-collections", "", "")
	request.Method = requests.GET
	return
}

// CreateListDataCollectionsResponse creates a response to parse from ListDataCollections response
func CreateListDataCollectionsResponse() (response *ListDataCollectionsResponse) {
	response = &ListDataCollectionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
