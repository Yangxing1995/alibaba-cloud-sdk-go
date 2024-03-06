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

// DescribeTTSConfig invokes the outboundbot.DescribeTTSConfig API synchronously
func (client *Client) DescribeTTSConfig(request *DescribeTTSConfigRequest) (response *DescribeTTSConfigResponse, err error) {
	response = CreateDescribeTTSConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTTSConfigWithChan invokes the outboundbot.DescribeTTSConfig API asynchronously
func (client *Client) DescribeTTSConfigWithChan(request *DescribeTTSConfigRequest) (<-chan *DescribeTTSConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeTTSConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTTSConfig(request)
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

// DescribeTTSConfigWithCallback invokes the outboundbot.DescribeTTSConfig API asynchronously
func (client *Client) DescribeTTSConfigWithCallback(request *DescribeTTSConfigRequest, callback func(response *DescribeTTSConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTTSConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeTTSConfig(request)
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

// DescribeTTSConfigRequest is the request struct for api DescribeTTSConfig
type DescribeTTSConfigRequest struct {
	*requests.RpcRequest
	ScriptId   string `position:"Query" name:"ScriptId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeTTSConfigResponse is the response struct for api DescribeTTSConfig
type DescribeTTSConfigResponse struct {
	*responses.BaseResponse
	HttpStatusCode int       `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string    `json:"Code" xml:"Code"`
	Message        string    `json:"Message" xml:"Message"`
	RequestId      string    `json:"RequestId" xml:"RequestId"`
	Success        bool      `json:"Success" xml:"Success"`
	TTSConfig      TTSConfig `json:"TTSConfig" xml:"TTSConfig"`
}

// CreateDescribeTTSConfigRequest creates a request to invoke DescribeTTSConfig API
func CreateDescribeTTSConfigRequest() (request *DescribeTTSConfigRequest) {
	request = &DescribeTTSConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DescribeTTSConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeTTSConfigResponse creates a response to parse from DescribeTTSConfig response
func CreateDescribeTTSConfigResponse() (response *DescribeTTSConfigResponse) {
	response = &DescribeTTSConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
