package sddp

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

// DisableUserConfig invokes the sddp.DisableUserConfig API synchronously
func (client *Client) DisableUserConfig(request *DisableUserConfigRequest) (response *DisableUserConfigResponse, err error) {
	response = CreateDisableUserConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DisableUserConfigWithChan invokes the sddp.DisableUserConfig API asynchronously
func (client *Client) DisableUserConfigWithChan(request *DisableUserConfigRequest) (<-chan *DisableUserConfigResponse, <-chan error) {
	responseChan := make(chan *DisableUserConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableUserConfig(request)
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

// DisableUserConfigWithCallback invokes the sddp.DisableUserConfig API asynchronously
func (client *Client) DisableUserConfigWithCallback(request *DisableUserConfigRequest, callback func(response *DisableUserConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableUserConfigResponse
		var err error
		defer close(result)
		response, err = client.DisableUserConfig(request)
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

// DisableUserConfigRequest is the request struct for api DisableUserConfig
type DisableUserConfigRequest struct {
	*requests.RpcRequest
	Code        string           `position:"Query" name:"Code"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Lang        string           `position:"Query" name:"Lang"`
}

// DisableUserConfigResponse is the response struct for api DisableUserConfig
type DisableUserConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableUserConfigRequest creates a request to invoke DisableUserConfig API
func CreateDisableUserConfigRequest() (request *DisableUserConfigRequest) {
	request = &DisableUserConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DisableUserConfig", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableUserConfigResponse creates a response to parse from DisableUserConfig response
func CreateDisableUserConfigResponse() (response *DisableUserConfigResponse) {
	response = &DisableUserConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
