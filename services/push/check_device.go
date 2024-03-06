package push

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

// CheckDevice invokes the push.CheckDevice API synchronously
func (client *Client) CheckDevice(request *CheckDeviceRequest) (response *CheckDeviceResponse, err error) {
	response = CreateCheckDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// CheckDeviceWithChan invokes the push.CheckDevice API asynchronously
func (client *Client) CheckDeviceWithChan(request *CheckDeviceRequest) (<-chan *CheckDeviceResponse, <-chan error) {
	responseChan := make(chan *CheckDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckDevice(request)
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

// CheckDeviceWithCallback invokes the push.CheckDevice API asynchronously
func (client *Client) CheckDeviceWithCallback(request *CheckDeviceRequest, callback func(response *CheckDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckDeviceResponse
		var err error
		defer close(result)
		response, err = client.CheckDevice(request)
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

// CheckDeviceRequest is the request struct for api CheckDevice
type CheckDeviceRequest struct {
	*requests.RpcRequest
	DeviceId string           `position:"Query" name:"DeviceId"`
	AppKey   requests.Integer `position:"Query" name:"AppKey"`
}

// CheckDeviceResponse is the response struct for api CheckDevice
type CheckDeviceResponse struct {
	*responses.BaseResponse
	Available bool   `json:"Available" xml:"Available"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCheckDeviceRequest creates a request to invoke CheckDevice API
func CreateCheckDeviceRequest() (request *CheckDeviceRequest) {
	request = &CheckDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "CheckDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckDeviceResponse creates a response to parse from CheckDevice response
func CreateCheckDeviceResponse() (response *CheckDeviceResponse) {
	response = &CheckDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
