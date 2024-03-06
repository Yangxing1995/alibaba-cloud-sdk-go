package oms

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

// PutDomainPartByProxy invokes the oms.PutDomainPartByProxy API synchronously
func (client *Client) PutDomainPartByProxy(request *PutDomainPartByProxyRequest) (response *PutDomainPartByProxyResponse, err error) {
	response = CreatePutDomainPartByProxyResponse()
	err = client.DoAction(request, response)
	return
}

// PutDomainPartByProxyWithChan invokes the oms.PutDomainPartByProxy API asynchronously
func (client *Client) PutDomainPartByProxyWithChan(request *PutDomainPartByProxyRequest) (<-chan *PutDomainPartByProxyResponse, <-chan error) {
	responseChan := make(chan *PutDomainPartByProxyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutDomainPartByProxy(request)
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

// PutDomainPartByProxyWithCallback invokes the oms.PutDomainPartByProxy API asynchronously
func (client *Client) PutDomainPartByProxyWithCallback(request *PutDomainPartByProxyRequest, callback func(response *PutDomainPartByProxyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutDomainPartByProxyResponse
		var err error
		defer close(result)
		response, err = client.PutDomainPartByProxy(request)
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

// PutDomainPartByProxyRequest is the request struct for api PutDomainPartByProxy
type PutDomainPartByProxyRequest struct {
	*requests.RpcRequest
	DomainCode string           `position:"Query" name:"DomainCode"`
	Data       string           `position:"Query" name:"Data"`
	DataType   string           `position:"Query" name:"DataType"`
	Compressed requests.Boolean `position:"Query" name:"Compressed"`
}

// PutDomainPartByProxyResponse is the response struct for api PutDomainPartByProxy
type PutDomainPartByProxyResponse struct {
	*responses.BaseResponse
	DataType   string `json:"DataType" xml:"DataType"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainCode string `json:"DomainCode" xml:"DomainCode"`
}

// CreatePutDomainPartByProxyRequest creates a request to invoke PutDomainPartByProxy API
func CreatePutDomainPartByProxyRequest() (request *PutDomainPartByProxyRequest) {
	request = &PutDomainPartByProxyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oms", "2016-06-15", "PutDomainPartByProxy", "", "")
	request.Method = requests.POST
	return
}

// CreatePutDomainPartByProxyResponse creates a response to parse from PutDomainPartByProxy response
func CreatePutDomainPartByProxyResponse() (response *PutDomainPartByProxyResponse) {
	response = &PutDomainPartByProxyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
