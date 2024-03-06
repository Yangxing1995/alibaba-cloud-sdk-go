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

// CreateGlobalQuestion invokes the outboundbot.CreateGlobalQuestion API synchronously
func (client *Client) CreateGlobalQuestion(request *CreateGlobalQuestionRequest) (response *CreateGlobalQuestionResponse, err error) {
	response = CreateCreateGlobalQuestionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGlobalQuestionWithChan invokes the outboundbot.CreateGlobalQuestion API asynchronously
func (client *Client) CreateGlobalQuestionWithChan(request *CreateGlobalQuestionRequest) (<-chan *CreateGlobalQuestionResponse, <-chan error) {
	responseChan := make(chan *CreateGlobalQuestionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGlobalQuestion(request)
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

// CreateGlobalQuestionWithCallback invokes the outboundbot.CreateGlobalQuestion API asynchronously
func (client *Client) CreateGlobalQuestionWithCallback(request *CreateGlobalQuestionRequest, callback func(response *CreateGlobalQuestionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGlobalQuestionResponse
		var err error
		defer close(result)
		response, err = client.CreateGlobalQuestion(request)
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

// CreateGlobalQuestionRequest is the request struct for api CreateGlobalQuestion
type CreateGlobalQuestionRequest struct {
	*requests.RpcRequest
	GlobalQuestionName string `position:"Query" name:"GlobalQuestionName"`
	Questions          string `position:"Query" name:"Questions"`
	Answers            string `position:"Query" name:"Answers"`
	ScriptId           string `position:"Query" name:"ScriptId"`
	InstanceId         string `position:"Query" name:"InstanceId"`
	GlobalQuestionType string `position:"Query" name:"GlobalQuestionType"`
}

// CreateGlobalQuestionResponse is the response struct for api CreateGlobalQuestion
type CreateGlobalQuestionResponse struct {
	*responses.BaseResponse
	HttpStatusCode   int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	GlobalQuestionId string `json:"GlobalQuestionId" xml:"GlobalQuestionId"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Success          bool   `json:"Success" xml:"Success"`
	Code             string `json:"Code" xml:"Code"`
	Message          string `json:"Message" xml:"Message"`
}

// CreateCreateGlobalQuestionRequest creates a request to invoke CreateGlobalQuestion API
func CreateCreateGlobalQuestionRequest() (request *CreateGlobalQuestionRequest) {
	request = &CreateGlobalQuestionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "CreateGlobalQuestion", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateGlobalQuestionResponse creates a response to parse from CreateGlobalQuestion response
func CreateCreateGlobalQuestionResponse() (response *CreateGlobalQuestionResponse) {
	response = &CreateGlobalQuestionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
