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

// SaveEffectiveDays invokes the outboundbot.SaveEffectiveDays API synchronously
func (client *Client) SaveEffectiveDays(request *SaveEffectiveDaysRequest) (response *SaveEffectiveDaysResponse, err error) {
	response = CreateSaveEffectiveDaysResponse()
	err = client.DoAction(request, response)
	return
}

// SaveEffectiveDaysWithChan invokes the outboundbot.SaveEffectiveDays API asynchronously
func (client *Client) SaveEffectiveDaysWithChan(request *SaveEffectiveDaysRequest) (<-chan *SaveEffectiveDaysResponse, <-chan error) {
	responseChan := make(chan *SaveEffectiveDaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveEffectiveDays(request)
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

// SaveEffectiveDaysWithCallback invokes the outboundbot.SaveEffectiveDays API asynchronously
func (client *Client) SaveEffectiveDaysWithCallback(request *SaveEffectiveDaysRequest, callback func(response *SaveEffectiveDaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveEffectiveDaysResponse
		var err error
		defer close(result)
		response, err = client.SaveEffectiveDays(request)
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

// SaveEffectiveDaysRequest is the request struct for api SaveEffectiveDays
type SaveEffectiveDaysRequest struct {
	*requests.RpcRequest
	EffectiveDays requests.Integer `position:"Query" name:"EffectiveDays"`
	StrategyLevel requests.Integer `position:"Query" name:"StrategyLevel"`
	EntryId       string           `position:"Query" name:"EntryId"`
}

// SaveEffectiveDaysResponse is the response struct for api SaveEffectiveDays
type SaveEffectiveDaysResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateSaveEffectiveDaysRequest creates a request to invoke SaveEffectiveDays API
func CreateSaveEffectiveDaysRequest() (request *SaveEffectiveDaysRequest) {
	request = &SaveEffectiveDaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "SaveEffectiveDays", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveEffectiveDaysResponse creates a response to parse from SaveEffectiveDays response
func CreateSaveEffectiveDaysResponse() (response *SaveEffectiveDaysResponse) {
	response = &SaveEffectiveDaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
