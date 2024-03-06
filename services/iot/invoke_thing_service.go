package iot

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

// InvokeThingService invokes the iot.InvokeThingService API synchronously
func (client *Client) InvokeThingService(request *InvokeThingServiceRequest) (response *InvokeThingServiceResponse, err error) {
	response = CreateInvokeThingServiceResponse()
	err = client.DoAction(request, response)
	return
}

// InvokeThingServiceWithChan invokes the iot.InvokeThingService API asynchronously
func (client *Client) InvokeThingServiceWithChan(request *InvokeThingServiceRequest) (<-chan *InvokeThingServiceResponse, <-chan error) {
	responseChan := make(chan *InvokeThingServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InvokeThingService(request)
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

// InvokeThingServiceWithCallback invokes the iot.InvokeThingService API asynchronously
func (client *Client) InvokeThingServiceWithCallback(request *InvokeThingServiceRequest, callback func(response *InvokeThingServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InvokeThingServiceResponse
		var err error
		defer close(result)
		response, err = client.InvokeThingService(request)
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

// InvokeThingServiceRequest is the request struct for api InvokeThingService
type InvokeThingServiceRequest struct {
	*requests.RpcRequest
	RealTenantId      string           `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string           `position:"Query" name:"RealTripartiteKey"`
	IotId             string           `position:"Query" name:"IotId"`
	Qos               requests.Integer `position:"Query" name:"Qos"`
	IotInstanceId     string           `position:"Query" name:"IotInstanceId"`
	Identifier        string           `position:"Query" name:"Identifier"`
	ProductKey        string           `position:"Query" name:"ProductKey"`
	Args              string           `position:"Query" name:"Args"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
	DeviceName        string           `position:"Query" name:"DeviceName"`
}

// InvokeThingServiceResponse is the response struct for api InvokeThingService
type InvokeThingServiceResponse struct {
	*responses.BaseResponse
	RequestId    string                   `json:"RequestId" xml:"RequestId"`
	Success      bool                     `json:"Success" xml:"Success"`
	ErrorMessage string                   `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string                   `json:"Code" xml:"Code"`
	Data         DataInInvokeThingService `json:"Data" xml:"Data"`
}

// CreateInvokeThingServiceRequest creates a request to invoke InvokeThingService API
func CreateInvokeThingServiceRequest() (request *InvokeThingServiceRequest) {
	request = &InvokeThingServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "InvokeThingService", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInvokeThingServiceResponse creates a response to parse from InvokeThingService response
func CreateInvokeThingServiceResponse() (response *InvokeThingServiceResponse) {
	response = &InvokeThingServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
