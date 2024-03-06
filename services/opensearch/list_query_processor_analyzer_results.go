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

// ListQueryProcessorAnalyzerResults invokes the opensearch.ListQueryProcessorAnalyzerResults API synchronously
func (client *Client) ListQueryProcessorAnalyzerResults(request *ListQueryProcessorAnalyzerResultsRequest) (response *ListQueryProcessorAnalyzerResultsResponse, err error) {
	response = CreateListQueryProcessorAnalyzerResultsResponse()
	err = client.DoAction(request, response)
	return
}

// ListQueryProcessorAnalyzerResultsWithChan invokes the opensearch.ListQueryProcessorAnalyzerResults API asynchronously
func (client *Client) ListQueryProcessorAnalyzerResultsWithChan(request *ListQueryProcessorAnalyzerResultsRequest) (<-chan *ListQueryProcessorAnalyzerResultsResponse, <-chan error) {
	responseChan := make(chan *ListQueryProcessorAnalyzerResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListQueryProcessorAnalyzerResults(request)
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

// ListQueryProcessorAnalyzerResultsWithCallback invokes the opensearch.ListQueryProcessorAnalyzerResults API asynchronously
func (client *Client) ListQueryProcessorAnalyzerResultsWithCallback(request *ListQueryProcessorAnalyzerResultsRequest, callback func(response *ListQueryProcessorAnalyzerResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListQueryProcessorAnalyzerResultsResponse
		var err error
		defer close(result)
		response, err = client.ListQueryProcessorAnalyzerResults(request)
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

// ListQueryProcessorAnalyzerResultsRequest is the request struct for api ListQueryProcessorAnalyzerResults
type ListQueryProcessorAnalyzerResultsRequest struct {
	*requests.RoaRequest
	AppId            string `position:"Path" name:"appId"`
	Name             string `position:"Path" name:"name"`
	Text             string `position:"Query" name:"text"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// ListQueryProcessorAnalyzerResultsResponse is the response struct for api ListQueryProcessorAnalyzerResults
type ListQueryProcessorAnalyzerResultsResponse struct {
	*responses.BaseResponse
	Result    map[string]interface{} `json:"result" xml:"result"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateListQueryProcessorAnalyzerResultsRequest creates a request to invoke ListQueryProcessorAnalyzerResults API
func CreateListQueryProcessorAnalyzerResultsRequest() (request *ListQueryProcessorAnalyzerResultsRequest) {
	request = &ListQueryProcessorAnalyzerResultsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListQueryProcessorAnalyzerResults", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/query-processors/[name]/analyze", "", "")
	request.Method = requests.GET
	return
}

// CreateListQueryProcessorAnalyzerResultsResponse creates a response to parse from ListQueryProcessorAnalyzerResults response
func CreateListQueryProcessorAnalyzerResultsResponse() (response *ListQueryProcessorAnalyzerResultsResponse) {
	response = &ListQueryProcessorAnalyzerResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
