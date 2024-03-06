package sae

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

// UpdateNamespaceVpc invokes the sae.UpdateNamespaceVpc API synchronously
func (client *Client) UpdateNamespaceVpc(request *UpdateNamespaceVpcRequest) (response *UpdateNamespaceVpcResponse, err error) {
	response = CreateUpdateNamespaceVpcResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateNamespaceVpcWithChan invokes the sae.UpdateNamespaceVpc API asynchronously
func (client *Client) UpdateNamespaceVpcWithChan(request *UpdateNamespaceVpcRequest) (<-chan *UpdateNamespaceVpcResponse, <-chan error) {
	responseChan := make(chan *UpdateNamespaceVpcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateNamespaceVpc(request)
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

// UpdateNamespaceVpcWithCallback invokes the sae.UpdateNamespaceVpc API asynchronously
func (client *Client) UpdateNamespaceVpcWithCallback(request *UpdateNamespaceVpcRequest, callback func(response *UpdateNamespaceVpcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateNamespaceVpcResponse
		var err error
		defer close(result)
		response, err = client.UpdateNamespaceVpc(request)
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

// UpdateNamespaceVpcRequest is the request struct for api UpdateNamespaceVpc
type UpdateNamespaceVpcRequest struct {
	*requests.RoaRequest
	NamespaceId      string `position:"Query" name:"NamespaceId"`
	VpcId            string `position:"Query" name:"VpcId"`
	NameSpaceShortId string `position:"Query" name:"NameSpaceShortId"`
}

// UpdateNamespaceVpcResponse is the response struct for api UpdateNamespaceVpc
type UpdateNamespaceVpcResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateNamespaceVpcRequest creates a request to invoke UpdateNamespaceVpc API
func CreateUpdateNamespaceVpcRequest() (request *UpdateNamespaceVpcRequest) {
	request = &UpdateNamespaceVpcRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "UpdateNamespaceVpc", "/pop/v1/sam/namespace/updateNamespaceVpc", "serverless", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateNamespaceVpcResponse creates a response to parse from UpdateNamespaceVpc response
func CreateUpdateNamespaceVpcResponse() (response *UpdateNamespaceVpcResponse) {
	response = &UpdateNamespaceVpcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
