package outboundbot

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

// ListScripts invokes the outboundbot.ListScripts API synchronously
func (client *Client) ListScripts(request *ListScriptsRequest) (response *ListScriptsResponse, err error) {
	response = CreateListScriptsResponse()
	err = client.DoAction(request, response)
	return
}

// ListScriptsWithChan invokes the outboundbot.ListScripts API asynchronously
func (client *Client) ListScriptsWithChan(request *ListScriptsRequest) (<-chan *ListScriptsResponse, <-chan error) {
	responseChan := make(chan *ListScriptsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListScripts(request)
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

// ListScriptsWithCallback invokes the outboundbot.ListScripts API asynchronously
func (client *Client) ListScriptsWithCallback(request *ListScriptsRequest, callback func(response *ListScriptsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListScriptsResponse
		var err error
		defer close(result)
		response, err = client.ListScripts(request)
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

// ListScriptsRequest is the request struct for api ListScripts
type ListScriptsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListScriptsResponse is the response struct for api ListScripts
type ListScriptsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string  `json:"Code" xml:"Code"`
	Message        string  `json:"Message" xml:"Message"`
	RequestId      string  `json:"RequestId" xml:"RequestId"`
	Success        bool    `json:"Success" xml:"Success"`
	Scripts        Scripts `json:"Scripts" xml:"Scripts"`
}

// CreateListScriptsRequest creates a request to invoke ListScripts API
func CreateListScriptsRequest() (request *ListScriptsRequest) {
	request = &ListScriptsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ListScripts", "", "")
	request.Method = requests.POST
	return
}

// CreateListScriptsResponse creates a response to parse from ListScripts response
func CreateListScriptsResponse() (response *ListScriptsResponse) {
	response = &ListScriptsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
