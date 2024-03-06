package ehpc

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

// GetIfEcsTypeSupportHtConfig invokes the ehpc.GetIfEcsTypeSupportHtConfig API synchronously
func (client *Client) GetIfEcsTypeSupportHtConfig(request *GetIfEcsTypeSupportHtConfigRequest) (response *GetIfEcsTypeSupportHtConfigResponse, err error) {
	response = CreateGetIfEcsTypeSupportHtConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetIfEcsTypeSupportHtConfigWithChan invokes the ehpc.GetIfEcsTypeSupportHtConfig API asynchronously
func (client *Client) GetIfEcsTypeSupportHtConfigWithChan(request *GetIfEcsTypeSupportHtConfigRequest) (<-chan *GetIfEcsTypeSupportHtConfigResponse, <-chan error) {
	responseChan := make(chan *GetIfEcsTypeSupportHtConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetIfEcsTypeSupportHtConfig(request)
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

// GetIfEcsTypeSupportHtConfigWithCallback invokes the ehpc.GetIfEcsTypeSupportHtConfig API asynchronously
func (client *Client) GetIfEcsTypeSupportHtConfigWithCallback(request *GetIfEcsTypeSupportHtConfigRequest, callback func(response *GetIfEcsTypeSupportHtConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetIfEcsTypeSupportHtConfigResponse
		var err error
		defer close(result)
		response, err = client.GetIfEcsTypeSupportHtConfig(request)
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

// GetIfEcsTypeSupportHtConfigRequest is the request struct for api GetIfEcsTypeSupportHtConfig
type GetIfEcsTypeSupportHtConfigRequest struct {
	*requests.RpcRequest
	InstanceType string `position:"Query" name:"InstanceType"`
}

// GetIfEcsTypeSupportHtConfigResponse is the response struct for api GetIfEcsTypeSupportHtConfig
type GetIfEcsTypeSupportHtConfigResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DefaultHtEnabled bool   `json:"DefaultHtEnabled" xml:"DefaultHtEnabled"`
	InstanceType     string `json:"InstanceType" xml:"InstanceType"`
	SupportHtConfig  bool   `json:"SupportHtConfig" xml:"SupportHtConfig"`
}

// CreateGetIfEcsTypeSupportHtConfigRequest creates a request to invoke GetIfEcsTypeSupportHtConfig API
func CreateGetIfEcsTypeSupportHtConfigRequest() (request *GetIfEcsTypeSupportHtConfigRequest) {
	request = &GetIfEcsTypeSupportHtConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "GetIfEcsTypeSupportHtConfig", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetIfEcsTypeSupportHtConfigResponse creates a response to parse from GetIfEcsTypeSupportHtConfig response
func CreateGetIfEcsTypeSupportHtConfigResponse() (response *GetIfEcsTypeSupportHtConfigResponse) {
	response = &GetIfEcsTypeSupportHtConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
