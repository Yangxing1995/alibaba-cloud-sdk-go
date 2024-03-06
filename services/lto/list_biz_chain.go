package lto

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

// ListBizChain invokes the lto.ListBizChain API synchronously
func (client *Client) ListBizChain(request *ListBizChainRequest) (response *ListBizChainResponse, err error) {
	response = CreateListBizChainResponse()
	err = client.DoAction(request, response)
	return
}

// ListBizChainWithChan invokes the lto.ListBizChain API asynchronously
func (client *Client) ListBizChainWithChan(request *ListBizChainRequest) (<-chan *ListBizChainResponse, <-chan error) {
	responseChan := make(chan *ListBizChainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBizChain(request)
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

// ListBizChainWithCallback invokes the lto.ListBizChain API asynchronously
func (client *Client) ListBizChainWithCallback(request *ListBizChainRequest, callback func(response *ListBizChainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBizChainResponse
		var err error
		defer close(result)
		response, err = client.ListBizChain(request)
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

// ListBizChainRequest is the request struct for api ListBizChain
type ListBizChainRequest struct {
	*requests.RpcRequest
	Num        requests.Integer `position:"Query" name:"Num"`
	BizChainId string           `position:"Query" name:"BizChainId"`
	Size       requests.Integer `position:"Query" name:"Size"`
	Name       string           `position:"Query" name:"Name"`
}

// ListBizChainResponse is the response struct for api ListBizChain
type ListBizChainResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateListBizChainRequest creates a request to invoke ListBizChain API
func CreateListBizChainRequest() (request *ListBizChainRequest) {
	request = &ListBizChainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "ListBizChain", "", "")
	request.Method = requests.POST
	return
}

// CreateListBizChainResponse creates a response to parse from ListBizChain response
func CreateListBizChainResponse() (response *ListBizChainResponse) {
	response = &ListBizChainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
