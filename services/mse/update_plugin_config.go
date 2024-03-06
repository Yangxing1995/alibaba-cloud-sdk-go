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

// UpdatePluginConfig invokes the mse.UpdatePluginConfig API synchronously
func (client *Client) UpdatePluginConfig(request *UpdatePluginConfigRequest) (response *UpdatePluginConfigResponse, err error) {
	response = CreateUpdatePluginConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePluginConfigWithChan invokes the mse.UpdatePluginConfig API asynchronously
func (client *Client) UpdatePluginConfigWithChan(request *UpdatePluginConfigRequest) (<-chan *UpdatePluginConfigResponse, <-chan error) {
	responseChan := make(chan *UpdatePluginConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePluginConfig(request)
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

// UpdatePluginConfigWithCallback invokes the mse.UpdatePluginConfig API asynchronously
func (client *Client) UpdatePluginConfigWithCallback(request *UpdatePluginConfigRequest, callback func(response *UpdatePluginConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePluginConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdatePluginConfig(request)
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

// UpdatePluginConfigRequest is the request struct for api UpdatePluginConfig
type UpdatePluginConfigRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	GmtModified     string           `position:"Query" name:"GmtModified"`
	Enable          requests.Boolean `position:"Query" name:"Enable"`
	Id              requests.Integer `position:"Query" name:"Id"`
	GatewayId       requests.Integer `position:"Query" name:"GatewayId"`
	PluginId        requests.Integer `position:"Query" name:"PluginId"`
	GmtCreate       string           `position:"Query" name:"GmtCreate"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
	ConfigLevel     requests.Integer `position:"Query" name:"ConfigLevel"`
	Config          string           `position:"Query" name:"Config"`
}

// UpdatePluginConfigResponse is the response struct for api UpdatePluginConfig
type UpdatePluginConfigResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           int    `json:"Code" xml:"Code"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           int64  `json:"Data" xml:"Data"`
}

// CreateUpdatePluginConfigRequest creates a request to invoke UpdatePluginConfig API
func CreateUpdatePluginConfigRequest() (request *UpdatePluginConfigRequest) {
	request = &UpdatePluginConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdatePluginConfig", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdatePluginConfigResponse creates a response to parse from UpdatePluginConfig response
func CreateUpdatePluginConfigResponse() (response *UpdatePluginConfigResponse) {
	response = &UpdatePluginConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
