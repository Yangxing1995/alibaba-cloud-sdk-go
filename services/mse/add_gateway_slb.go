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

// AddGatewaySlb invokes the mse.AddGatewaySlb API synchronously
func (client *Client) AddGatewaySlb(request *AddGatewaySlbRequest) (response *AddGatewaySlbResponse, err error) {
	response = CreateAddGatewaySlbResponse()
	err = client.DoAction(request, response)
	return
}

// AddGatewaySlbWithChan invokes the mse.AddGatewaySlb API asynchronously
func (client *Client) AddGatewaySlbWithChan(request *AddGatewaySlbRequest) (<-chan *AddGatewaySlbResponse, <-chan error) {
	responseChan := make(chan *AddGatewaySlbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddGatewaySlb(request)
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

// AddGatewaySlbWithCallback invokes the mse.AddGatewaySlb API asynchronously
func (client *Client) AddGatewaySlbWithCallback(request *AddGatewaySlbRequest, callback func(response *AddGatewaySlbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddGatewaySlbResponse
		var err error
		defer close(result)
		response, err = client.AddGatewaySlb(request)
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

// AddGatewaySlbRequest is the request struct for api AddGatewaySlb
type AddGatewaySlbRequest struct {
	*requests.RpcRequest
	MseSessionId        string                       `position:"Query" name:"MseSessionId"`
	SlbId               string                       `position:"Query" name:"SlbId"`
	GatewayUniqueId     string                       `position:"Query" name:"GatewayUniqueId"`
	Type                string                       `position:"Query" name:"Type"`
	HttpPort            requests.Integer             `position:"Query" name:"HttpPort"`
	ServiceWeight       requests.Integer             `position:"Query" name:"ServiceWeight"`
	VServerGroupId      string                       `position:"Query" name:"VServerGroupId"`
	VServiceList        *[]AddGatewaySlbVServiceList `position:"Query" name:"VServiceList"  type:"Json"`
	HttpsVServerGroupId string                       `position:"Query" name:"HttpsVServerGroupId"`
	AcceptLanguage      string                       `position:"Query" name:"AcceptLanguage"`
	HttpsPort           requests.Integer             `position:"Query" name:"HttpsPort"`
}

// AddGatewaySlbVServiceList is a repeated param struct in AddGatewaySlbRequest
type AddGatewaySlbVServiceList struct {
	VServerGroupId   string `name:"VServerGroupId"`
	Protocol         string `name:"Protocol"`
	Port             string `name:"Port"`
	VServerGroupName string `name:"VServerGroupName"`
}

// AddGatewaySlbResponse is the response struct for api AddGatewaySlb
type AddGatewaySlbResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           string `json:"Data" xml:"Data"`
}

// CreateAddGatewaySlbRequest creates a request to invoke AddGatewaySlb API
func CreateAddGatewaySlbRequest() (request *AddGatewaySlbRequest) {
	request = &AddGatewaySlbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "AddGatewaySlb", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddGatewaySlbResponse creates a response to parse from AddGatewaySlb response
func CreateAddGatewaySlbResponse() (response *AddGatewaySlbResponse) {
	response = &AddGatewaySlbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
