package arms

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

// ConfigApp invokes the arms.ConfigApp API synchronously
func (client *Client) ConfigApp(request *ConfigAppRequest) (response *ConfigAppResponse, err error) {
	response = CreateConfigAppResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigAppWithChan invokes the arms.ConfigApp API asynchronously
func (client *Client) ConfigAppWithChan(request *ConfigAppRequest) (<-chan *ConfigAppResponse, <-chan error) {
	responseChan := make(chan *ConfigAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigApp(request)
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

// ConfigAppWithCallback invokes the arms.ConfigApp API asynchronously
func (client *Client) ConfigAppWithCallback(request *ConfigAppRequest, callback func(response *ConfigAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigAppResponse
		var err error
		defer close(result)
		response, err = client.ConfigApp(request)
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

// ConfigAppRequest is the request struct for api ConfigApp
type ConfigAppRequest struct {
	*requests.RpcRequest
	AppIds string `position:"Query" name:"AppIds"`
	Enable string `position:"Query" name:"Enable"`
	Type   string `position:"Query" name:"Type"`
}

// ConfigAppResponse is the response struct for api ConfigApp
type ConfigAppResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigAppRequest creates a request to invoke ConfigApp API
func CreateConfigAppRequest() (request *ConfigAppRequest) {
	request = &ConfigAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ConfigApp", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigAppResponse creates a response to parse from ConfigApp response
func CreateConfigAppResponse() (response *ConfigAppResponse) {
	response = &ConfigAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
