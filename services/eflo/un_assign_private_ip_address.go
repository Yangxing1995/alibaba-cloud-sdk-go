package eflo

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

// UnAssignPrivateIpAddress invokes the eflo.UnAssignPrivateIpAddress API synchronously
func (client *Client) UnAssignPrivateIpAddress(request *UnAssignPrivateIpAddressRequest) (response *UnAssignPrivateIpAddressResponse, err error) {
	response = CreateUnAssignPrivateIpAddressResponse()
	err = client.DoAction(request, response)
	return
}

// UnAssignPrivateIpAddressWithChan invokes the eflo.UnAssignPrivateIpAddress API asynchronously
func (client *Client) UnAssignPrivateIpAddressWithChan(request *UnAssignPrivateIpAddressRequest) (<-chan *UnAssignPrivateIpAddressResponse, <-chan error) {
	responseChan := make(chan *UnAssignPrivateIpAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnAssignPrivateIpAddress(request)
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

// UnAssignPrivateIpAddressWithCallback invokes the eflo.UnAssignPrivateIpAddress API asynchronously
func (client *Client) UnAssignPrivateIpAddressWithCallback(request *UnAssignPrivateIpAddressRequest, callback func(response *UnAssignPrivateIpAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnAssignPrivateIpAddressResponse
		var err error
		defer close(result)
		response, err = client.UnAssignPrivateIpAddress(request)
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

// UnAssignPrivateIpAddressRequest is the request struct for api UnAssignPrivateIpAddress
type UnAssignPrivateIpAddressRequest struct {
	*requests.RpcRequest
	SubnetId           string `position:"Body" name:"SubnetId"`
	ClientToken        string `position:"Body" name:"ClientToken"`
	PrivateIpAddress   string `position:"Body" name:"PrivateIpAddress"`
	IpName             string `position:"Body" name:"IpName"`
	NetworkInterfaceId string `position:"Body" name:"NetworkInterfaceId"`
}

// UnAssignPrivateIpAddressResponse is the response struct for api UnAssignPrivateIpAddress
type UnAssignPrivateIpAddressResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateUnAssignPrivateIpAddressRequest creates a request to invoke UnAssignPrivateIpAddress API
func CreateUnAssignPrivateIpAddressRequest() (request *UnAssignPrivateIpAddressRequest) {
	request = &UnAssignPrivateIpAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "UnAssignPrivateIpAddress", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnAssignPrivateIpAddressResponse creates a response to parse from UnAssignPrivateIpAddress response
func CreateUnAssignPrivateIpAddressResponse() (response *UnAssignPrivateIpAddressResponse) {
	response = &UnAssignPrivateIpAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
