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

// SetGWSClusterPolicy invokes the ehpc.SetGWSClusterPolicy API synchronously
func (client *Client) SetGWSClusterPolicy(request *SetGWSClusterPolicyRequest) (response *SetGWSClusterPolicyResponse, err error) {
	response = CreateSetGWSClusterPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// SetGWSClusterPolicyWithChan invokes the ehpc.SetGWSClusterPolicy API asynchronously
func (client *Client) SetGWSClusterPolicyWithChan(request *SetGWSClusterPolicyRequest) (<-chan *SetGWSClusterPolicyResponse, <-chan error) {
	responseChan := make(chan *SetGWSClusterPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetGWSClusterPolicy(request)
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

// SetGWSClusterPolicyWithCallback invokes the ehpc.SetGWSClusterPolicy API asynchronously
func (client *Client) SetGWSClusterPolicyWithCallback(request *SetGWSClusterPolicyRequest, callback func(response *SetGWSClusterPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetGWSClusterPolicyResponse
		var err error
		defer close(result)
		response, err = client.SetGWSClusterPolicy(request)
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

// SetGWSClusterPolicyRequest is the request struct for api SetGWSClusterPolicy
type SetGWSClusterPolicyRequest struct {
	*requests.RpcRequest
	Watermark   string           `position:"Query" name:"Watermark"`
	LocalDrive  string           `position:"Query" name:"LocalDrive"`
	ClusterId   string           `position:"Query" name:"ClusterId"`
	Clipboard   string           `position:"Query" name:"Clipboard"`
	UsbRedirect string           `position:"Query" name:"UsbRedirect"`
	AsyncMode   requests.Boolean `position:"Query" name:"AsyncMode"`
	UdpPort     string           `position:"Query" name:"UdpPort"`
}

// SetGWSClusterPolicyResponse is the response struct for api SetGWSClusterPolicy
type SetGWSClusterPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetGWSClusterPolicyRequest creates a request to invoke SetGWSClusterPolicy API
func CreateSetGWSClusterPolicyRequest() (request *SetGWSClusterPolicyRequest) {
	request = &SetGWSClusterPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "SetGWSClusterPolicy", "ehs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetGWSClusterPolicyResponse creates a response to parse from SetGWSClusterPolicy response
func CreateSetGWSClusterPolicyResponse() (response *SetGWSClusterPolicyResponse) {
	response = &SetGWSClusterPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
