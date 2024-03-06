package amqp_open

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

// UpdateInstanceName invokes the amqp_open.UpdateInstanceName API synchronously
func (client *Client) UpdateInstanceName(request *UpdateInstanceNameRequest) (response *UpdateInstanceNameResponse, err error) {
	response = CreateUpdateInstanceNameResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateInstanceNameWithChan invokes the amqp_open.UpdateInstanceName API asynchronously
func (client *Client) UpdateInstanceNameWithChan(request *UpdateInstanceNameRequest) (<-chan *UpdateInstanceNameResponse, <-chan error) {
	responseChan := make(chan *UpdateInstanceNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateInstanceName(request)
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

// UpdateInstanceNameWithCallback invokes the amqp_open.UpdateInstanceName API asynchronously
func (client *Client) UpdateInstanceNameWithCallback(request *UpdateInstanceNameRequest, callback func(response *UpdateInstanceNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateInstanceNameResponse
		var err error
		defer close(result)
		response, err = client.UpdateInstanceName(request)
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

// UpdateInstanceNameRequest is the request struct for api UpdateInstanceName
type UpdateInstanceNameRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceName string `position:"Query" name:"InstanceName"`
}

// UpdateInstanceNameResponse is the response struct for api UpdateInstanceName
type UpdateInstanceNameResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateInstanceNameRequest creates a request to invoke UpdateInstanceName API
func CreateUpdateInstanceNameRequest() (request *UpdateInstanceNameRequest) {
	request = &UpdateInstanceNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("amqp-open", "2019-12-12", "UpdateInstanceName", "onsproxy", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateInstanceNameResponse creates a response to parse from UpdateInstanceName response
func CreateUpdateInstanceNameResponse() (response *UpdateInstanceNameResponse) {
	response = &UpdateInstanceNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
