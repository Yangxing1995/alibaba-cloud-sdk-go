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

// UpdateGatewayAuthConsumer invokes the mse.UpdateGatewayAuthConsumer API synchronously
func (client *Client) UpdateGatewayAuthConsumer(request *UpdateGatewayAuthConsumerRequest) (response *UpdateGatewayAuthConsumerResponse, err error) {
	response = CreateUpdateGatewayAuthConsumerResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewayAuthConsumerWithChan invokes the mse.UpdateGatewayAuthConsumer API asynchronously
func (client *Client) UpdateGatewayAuthConsumerWithChan(request *UpdateGatewayAuthConsumerRequest) (<-chan *UpdateGatewayAuthConsumerResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewayAuthConsumerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewayAuthConsumer(request)
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

// UpdateGatewayAuthConsumerWithCallback invokes the mse.UpdateGatewayAuthConsumer API asynchronously
func (client *Client) UpdateGatewayAuthConsumerWithCallback(request *UpdateGatewayAuthConsumerRequest, callback func(response *UpdateGatewayAuthConsumerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewayAuthConsumerResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewayAuthConsumer(request)
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

// UpdateGatewayAuthConsumerRequest is the request struct for api UpdateGatewayAuthConsumer
type UpdateGatewayAuthConsumerRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	KeyName         string           `position:"Query" name:"KeyName"`
	Id              requests.Integer `position:"Query" name:"Id"`
	TokenPrefix     string           `position:"Query" name:"TokenPrefix"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	Jwks            string           `position:"Query" name:"Jwks"`
	Description     string           `position:"Query" name:"Description"`
	TokenPosition   string           `position:"Query" name:"TokenPosition"`
	EncodeType      string           `position:"Query" name:"EncodeType"`
	KeyValue        string           `position:"Query" name:"KeyValue"`
	TokenPass       requests.Boolean `position:"Query" name:"TokenPass"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
	TokenName       string           `position:"Query" name:"TokenName"`
}

// UpdateGatewayAuthConsumerResponse is the response struct for api UpdateGatewayAuthConsumer
type UpdateGatewayAuthConsumerResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           int    `json:"Code" xml:"Code"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           int64  `json:"Data" xml:"Data"`
}

// CreateUpdateGatewayAuthConsumerRequest creates a request to invoke UpdateGatewayAuthConsumer API
func CreateUpdateGatewayAuthConsumerRequest() (request *UpdateGatewayAuthConsumerRequest) {
	request = &UpdateGatewayAuthConsumerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateGatewayAuthConsumer", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewayAuthConsumerResponse creates a response to parse from UpdateGatewayAuthConsumer response
func CreateUpdateGatewayAuthConsumerResponse() (response *UpdateGatewayAuthConsumerResponse) {
	response = &UpdateGatewayAuthConsumerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
