package ddoscoo

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

// ModifyPort invokes the ddoscoo.ModifyPort API synchronously
func (client *Client) ModifyPort(request *ModifyPortRequest) (response *ModifyPortResponse, err error) {
	response = CreateModifyPortResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPortWithChan invokes the ddoscoo.ModifyPort API asynchronously
func (client *Client) ModifyPortWithChan(request *ModifyPortRequest) (<-chan *ModifyPortResponse, <-chan error) {
	responseChan := make(chan *ModifyPortResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPort(request)
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

// ModifyPortWithCallback invokes the ddoscoo.ModifyPort API asynchronously
func (client *Client) ModifyPortWithCallback(request *ModifyPortRequest, callback func(response *ModifyPortResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPortResponse
		var err error
		defer close(result)
		response, err = client.ModifyPort(request)
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

// ModifyPortRequest is the request struct for api ModifyPort
type ModifyPortRequest struct {
	*requests.RpcRequest
	SourceIp         string    `position:"Query" name:"SourceIp"`
	BackendPort      string    `position:"Query" name:"BackendPort"`
	FrontendProtocol string    `position:"Query" name:"FrontendProtocol"`
	InstanceId       string    `position:"Query" name:"InstanceId"`
	RealServers      *[]string `position:"Query" name:"RealServers"  type:"Repeated"`
	FrontendPort     string    `position:"Query" name:"FrontendPort"`
}

// ModifyPortResponse is the response struct for api ModifyPort
type ModifyPortResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyPortRequest creates a request to invoke ModifyPort API
func CreateModifyPortRequest() (request *ModifyPortRequest) {
	request = &ModifyPortRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ModifyPort", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyPortResponse creates a response to parse from ModifyPort response
func CreateModifyPortResponse() (response *ModifyPortResponse) {
	response = &ModifyPortResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
