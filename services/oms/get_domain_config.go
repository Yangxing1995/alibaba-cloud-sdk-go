package oms

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

// GetDomainConfig invokes the oms.GetDomainConfig API synchronously
func (client *Client) GetDomainConfig(request *GetDomainConfigRequest) (response *GetDomainConfigResponse, err error) {
	response = CreateGetDomainConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetDomainConfigWithChan invokes the oms.GetDomainConfig API asynchronously
func (client *Client) GetDomainConfigWithChan(request *GetDomainConfigRequest) (<-chan *GetDomainConfigResponse, <-chan error) {
	responseChan := make(chan *GetDomainConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDomainConfig(request)
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

// GetDomainConfigWithCallback invokes the oms.GetDomainConfig API asynchronously
func (client *Client) GetDomainConfigWithCallback(request *GetDomainConfigRequest, callback func(response *GetDomainConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDomainConfigResponse
		var err error
		defer close(result)
		response, err = client.GetDomainConfig(request)
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

// GetDomainConfigRequest is the request struct for api GetDomainConfig
type GetDomainConfigRequest struct {
	*requests.RpcRequest
	DomainCode     string           `position:"Query" name:"DomainCode"`
	CompressEnable requests.Boolean `position:"Query" name:"CompressEnable"`
}

// GetDomainConfigResponse is the response struct for api GetDomainConfig
type GetDomainConfigResponse struct {
	*responses.BaseResponse
	Compressed bool   `json:"Compressed" xml:"Compressed"`
	Data       string `json:"Data" xml:"Data"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainCode string `json:"DomainCode" xml:"DomainCode"`
}

// CreateGetDomainConfigRequest creates a request to invoke GetDomainConfig API
func CreateGetDomainConfigRequest() (request *GetDomainConfigRequest) {
	request = &GetDomainConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oms", "2016-06-15", "GetDomainConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateGetDomainConfigResponse creates a response to parse from GetDomainConfig response
func CreateGetDomainConfigResponse() (response *GetDomainConfigResponse) {
	response = &GetDomainConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
