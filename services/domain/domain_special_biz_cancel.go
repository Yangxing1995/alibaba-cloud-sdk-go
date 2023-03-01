package domain

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

// DomainSpecialBizCancel invokes the domain.DomainSpecialBizCancel API synchronously
func (client *Client) DomainSpecialBizCancel(request *DomainSpecialBizCancelRequest) (response *DomainSpecialBizCancelResponse, err error) {
	response = CreateDomainSpecialBizCancelResponse()
	err = client.DoAction(request, response)
	return
}

// DomainSpecialBizCancelWithChan invokes the domain.DomainSpecialBizCancel API asynchronously
func (client *Client) DomainSpecialBizCancelWithChan(request *DomainSpecialBizCancelRequest) (<-chan *DomainSpecialBizCancelResponse, <-chan error) {
	responseChan := make(chan *DomainSpecialBizCancelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DomainSpecialBizCancel(request)
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

// DomainSpecialBizCancelWithCallback invokes the domain.DomainSpecialBizCancel API asynchronously
func (client *Client) DomainSpecialBizCancelWithCallback(request *DomainSpecialBizCancelRequest, callback func(response *DomainSpecialBizCancelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DomainSpecialBizCancelResponse
		var err error
		defer close(result)
		response, err = client.DomainSpecialBizCancel(request)
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

// DomainSpecialBizCancelRequest is the request struct for api DomainSpecialBizCancel
type DomainSpecialBizCancelRequest struct {
	*requests.RpcRequest
	BizId requests.Integer `position:"Body" name:"BizId"`
}

// DomainSpecialBizCancelResponse is the response struct for api DomainSpecialBizCancel
type DomainSpecialBizCancelResponse struct {
	*responses.BaseResponse
}

// CreateDomainSpecialBizCancelRequest creates a request to invoke DomainSpecialBizCancel API
func CreateDomainSpecialBizCancelRequest() (request *DomainSpecialBizCancelRequest) {
	request = &DomainSpecialBizCancelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "DomainSpecialBizCancel", "", "")
	request.Method = requests.POST
	return
}

// CreateDomainSpecialBizCancelResponse creates a response to parse from DomainSpecialBizCancel response
func CreateDomainSpecialBizCancelResponse() (response *DomainSpecialBizCancelResponse) {
	response = &DomainSpecialBizCancelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
