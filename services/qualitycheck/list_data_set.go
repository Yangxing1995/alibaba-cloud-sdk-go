package qualitycheck

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

// ListDataSet invokes the qualitycheck.ListDataSet API synchronously
func (client *Client) ListDataSet(request *ListDataSetRequest) (response *ListDataSetResponse, err error) {
	response = CreateListDataSetResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataSetWithChan invokes the qualitycheck.ListDataSet API asynchronously
func (client *Client) ListDataSetWithChan(request *ListDataSetRequest) (<-chan *ListDataSetResponse, <-chan error) {
	responseChan := make(chan *ListDataSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataSet(request)
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

// ListDataSetWithCallback invokes the qualitycheck.ListDataSet API asynchronously
func (client *Client) ListDataSetWithCallback(request *ListDataSetRequest, callback func(response *ListDataSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataSetResponse
		var err error
		defer close(result)
		response, err = client.ListDataSet(request)
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

// ListDataSetRequest is the request struct for api ListDataSet
type ListDataSetRequest struct {
	*requests.RpcRequest
	JsonStr       string           `position:"Query" name:"JsonStr"`
	BaseMeAgentId requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// ListDataSetResponse is the response struct for api ListDataSet
type ListDataSetResponse struct {
	*responses.BaseResponse
	Count          int                   `json:"Count" xml:"Count"`
	CurrentPage    int                   `json:"CurrentPage" xml:"CurrentPage"`
	PageSize       int                   `json:"PageSize" xml:"PageSize"`
	PageNumber     int                   `json:"PageNumber" xml:"PageNumber"`
	RequestId      string                `json:"RequestId" xml:"RequestId"`
	Success        bool                  `json:"Success" xml:"Success"`
	Code           string                `json:"Code" xml:"Code"`
	Message        string                `json:"Message" xml:"Message"`
	HttpStatusCode int                   `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Messages       MessagesInListDataSet `json:"Messages" xml:"Messages"`
	Data           DataInListDataSet     `json:"Data" xml:"Data"`
}

// CreateListDataSetRequest creates a request to invoke ListDataSet API
func CreateListDataSetRequest() (request *ListDataSetRequest) {
	request = &ListDataSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "ListDataSet", "", "")
	request.Method = requests.POST
	return
}

// CreateListDataSetResponse creates a response to parse from ListDataSet response
func CreateListDataSetResponse() (response *ListDataSetResponse) {
	response = &ListDataSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
