package mse

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

// GetMseFeatureSwitch invokes the mse.GetMseFeatureSwitch API synchronously
func (client *Client) GetMseFeatureSwitch(request *GetMseFeatureSwitchRequest) (response *GetMseFeatureSwitchResponse, err error) {
	response = CreateGetMseFeatureSwitchResponse()
	err = client.DoAction(request, response)
	return
}

// GetMseFeatureSwitchWithChan invokes the mse.GetMseFeatureSwitch API asynchronously
func (client *Client) GetMseFeatureSwitchWithChan(request *GetMseFeatureSwitchRequest) (<-chan *GetMseFeatureSwitchResponse, <-chan error) {
	responseChan := make(chan *GetMseFeatureSwitchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMseFeatureSwitch(request)
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

// GetMseFeatureSwitchWithCallback invokes the mse.GetMseFeatureSwitch API asynchronously
func (client *Client) GetMseFeatureSwitchWithCallback(request *GetMseFeatureSwitchRequest, callback func(response *GetMseFeatureSwitchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMseFeatureSwitchResponse
		var err error
		defer close(result)
		response, err = client.GetMseFeatureSwitch(request)
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

// GetMseFeatureSwitchRequest is the request struct for api GetMseFeatureSwitch
type GetMseFeatureSwitchRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// GetMseFeatureSwitchResponse is the response struct for api GetMseFeatureSwitch
type GetMseFeatureSwitchResponse struct {
	*responses.BaseResponse
	Message   string                 `json:"Message" xml:"Message"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Result    map[string]interface{} `json:"Result" xml:"Result"`
	ErrorCode string                 `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool                   `json:"Success" xml:"Success"`
}

// CreateGetMseFeatureSwitchRequest creates a request to invoke GetMseFeatureSwitch API
func CreateGetMseFeatureSwitchRequest() (request *GetMseFeatureSwitchRequest) {
	request = &GetMseFeatureSwitchRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetMseFeatureSwitch", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMseFeatureSwitchResponse creates a response to parse from GetMseFeatureSwitch response
func CreateGetMseFeatureSwitchResponse() (response *GetMseFeatureSwitchResponse) {
	response = &GetMseFeatureSwitchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
