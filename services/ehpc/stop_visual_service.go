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

// StopVisualService invokes the ehpc.StopVisualService API synchronously
func (client *Client) StopVisualService(request *StopVisualServiceRequest) (response *StopVisualServiceResponse, err error) {
	response = CreateStopVisualServiceResponse()
	err = client.DoAction(request, response)
	return
}

// StopVisualServiceWithChan invokes the ehpc.StopVisualService API asynchronously
func (client *Client) StopVisualServiceWithChan(request *StopVisualServiceRequest) (<-chan *StopVisualServiceResponse, <-chan error) {
	responseChan := make(chan *StopVisualServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopVisualService(request)
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

// StopVisualServiceWithCallback invokes the ehpc.StopVisualService API asynchronously
func (client *Client) StopVisualServiceWithCallback(request *StopVisualServiceRequest, callback func(response *StopVisualServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopVisualServiceResponse
		var err error
		defer close(result)
		response, err = client.StopVisualService(request)
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

// StopVisualServiceRequest is the request struct for api StopVisualService
type StopVisualServiceRequest struct {
	*requests.RpcRequest
	ClusterId string           `position:"Query" name:"ClusterId"`
	Port      requests.Integer `position:"Query" name:"Port"`
	CidrIp    string           `position:"Query" name:"CidrIp"`
}

// StopVisualServiceResponse is the response struct for api StopVisualService
type StopVisualServiceResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopVisualServiceRequest creates a request to invoke StopVisualService API
func CreateStopVisualServiceRequest() (request *StopVisualServiceRequest) {
	request = &StopVisualServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "StopVisualService", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateStopVisualServiceResponse creates a response to parse from StopVisualService response
func CreateStopVisualServiceResponse() (response *StopVisualServiceResponse) {
	response = &StopVisualServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
