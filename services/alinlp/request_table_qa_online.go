package alinlp

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

// RequestTableQAOnline invokes the alinlp.RequestTableQAOnline API synchronously
func (client *Client) RequestTableQAOnline(request *RequestTableQAOnlineRequest) (response *RequestTableQAOnlineResponse, err error) {
	response = CreateRequestTableQAOnlineResponse()
	err = client.DoAction(request, response)
	return
}

// RequestTableQAOnlineWithChan invokes the alinlp.RequestTableQAOnline API asynchronously
func (client *Client) RequestTableQAOnlineWithChan(request *RequestTableQAOnlineRequest) (<-chan *RequestTableQAOnlineResponse, <-chan error) {
	responseChan := make(chan *RequestTableQAOnlineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RequestTableQAOnline(request)
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

// RequestTableQAOnlineWithCallback invokes the alinlp.RequestTableQAOnline API asynchronously
func (client *Client) RequestTableQAOnlineWithCallback(request *RequestTableQAOnlineRequest, callback func(response *RequestTableQAOnlineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RequestTableQAOnlineResponse
		var err error
		defer close(result)
		response, err = client.RequestTableQAOnline(request)
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

// RequestTableQAOnlineRequest is the request struct for api RequestTableQAOnline
type RequestTableQAOnlineRequest struct {
	*requests.RpcRequest
	Question    string `position:"Body" name:"Question"`
	Params      string `position:"Body" name:"Params"`
	ServiceCode string `position:"Body" name:"ServiceCode"`
	BotId       string `position:"Body" name:"BotId"`
}

// RequestTableQAOnlineResponse is the response struct for api RequestTableQAOnline
type RequestTableQAOnlineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateRequestTableQAOnlineRequest creates a request to invoke RequestTableQAOnline API
func CreateRequestTableQAOnlineRequest() (request *RequestTableQAOnlineRequest) {
	request = &RequestTableQAOnlineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "RequestTableQAOnline", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRequestTableQAOnlineResponse creates a response to parse from RequestTableQAOnline response
func CreateRequestTableQAOnlineResponse() (response *RequestTableQAOnlineResponse) {
	response = &RequestTableQAOnlineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
