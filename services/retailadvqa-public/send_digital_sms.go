package retailadvqa_public

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

// SendDigitalSms invokes the retailadvqa_public.SendDigitalSms API synchronously
func (client *Client) SendDigitalSms(request *SendDigitalSmsRequest) (response *SendDigitalSmsResponse, err error) {
	response = CreateSendDigitalSmsResponse()
	err = client.DoAction(request, response)
	return
}

// SendDigitalSmsWithChan invokes the retailadvqa_public.SendDigitalSms API asynchronously
func (client *Client) SendDigitalSmsWithChan(request *SendDigitalSmsRequest) (<-chan *SendDigitalSmsResponse, <-chan error) {
	responseChan := make(chan *SendDigitalSmsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendDigitalSms(request)
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

// SendDigitalSmsWithCallback invokes the retailadvqa_public.SendDigitalSms API asynchronously
func (client *Client) SendDigitalSmsWithCallback(request *SendDigitalSmsRequest, callback func(response *SendDigitalSmsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendDigitalSmsResponse
		var err error
		defer close(result)
		response, err = client.SendDigitalSms(request)
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

// SendDigitalSmsRequest is the request struct for api SendDigitalSms
type SendDigitalSmsRequest struct {
	*requests.RpcRequest
	AccessId      string           `position:"Query" name:"AccessId"`
	PhoneNumbers  string           `position:"Query" name:"PhoneNumbers"`
	ChannelType   requests.Integer `position:"Query" name:"ChannelType"`
	TenantId      string           `position:"Query" name:"TenantId"`
	TaskName      string           `position:"Query" name:"TaskName"`
	OutId         string           `position:"Query" name:"OutId"`
	PlatformId    string           `position:"Query" name:"PlatformId"`
	SmsTemplateId string           `position:"Query" name:"SmsTemplateId"`
	WorkspaceId   string           `position:"Query" name:"WorkspaceId"`
}

// SendDigitalSmsResponse is the response struct for api SendDigitalSms
type SendDigitalSmsResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	Success   bool   `json:"Success" xml:"Success"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSendDigitalSmsRequest creates a request to invoke SendDigitalSms API
func CreateSendDigitalSmsRequest() (request *SendDigitalSmsRequest) {
	request = &SendDigitalSmsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "SendDigitalSms", "", "")
	request.Method = requests.GET
	return
}

// CreateSendDigitalSmsResponse creates a response to parse from SendDigitalSms response
func CreateSendDigitalSmsResponse() (response *SendDigitalSmsResponse) {
	response = &SendDigitalSmsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
