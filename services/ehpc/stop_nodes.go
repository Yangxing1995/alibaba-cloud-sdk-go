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

// StopNodes invokes the ehpc.StopNodes API synchronously
func (client *Client) StopNodes(request *StopNodesRequest) (response *StopNodesResponse, err error) {
	response = CreateStopNodesResponse()
	err = client.DoAction(request, response)
	return
}

// StopNodesWithChan invokes the ehpc.StopNodes API asynchronously
func (client *Client) StopNodesWithChan(request *StopNodesRequest) (<-chan *StopNodesResponse, <-chan error) {
	responseChan := make(chan *StopNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopNodes(request)
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

// StopNodesWithCallback invokes the ehpc.StopNodes API asynchronously
func (client *Client) StopNodesWithCallback(request *StopNodesRequest, callback func(response *StopNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopNodesResponse
		var err error
		defer close(result)
		response, err = client.StopNodes(request)
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

// StopNodesRequest is the request struct for api StopNodes
type StopNodesRequest struct {
	*requests.RpcRequest
	Role      string               `position:"Query" name:"Role"`
	Instance  *[]StopNodesInstance `position:"Query" name:"Instance"  type:"Repeated"`
	ClusterId string               `position:"Query" name:"ClusterId"`
}

// StopNodesInstance is a repeated param struct in StopNodesRequest
type StopNodesInstance struct {
	Id string `name:"Id"`
}

// StopNodesResponse is the response struct for api StopNodes
type StopNodesResponse struct {
	*responses.BaseResponse
	TaskId    string `json:"TaskId" xml:"TaskId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopNodesRequest creates a request to invoke StopNodes API
func CreateStopNodesRequest() (request *StopNodesRequest) {
	request = &StopNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "StopNodes", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateStopNodesResponse creates a response to parse from StopNodes response
func CreateStopNodesResponse() (response *StopNodesResponse) {
	response = &StopNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
