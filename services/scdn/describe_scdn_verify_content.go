package scdn

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

// DescribeScdnVerifyContent invokes the scdn.DescribeScdnVerifyContent API synchronously
func (client *Client) DescribeScdnVerifyContent(request *DescribeScdnVerifyContentRequest) (response *DescribeScdnVerifyContentResponse, err error) {
	response = CreateDescribeScdnVerifyContentResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnVerifyContentWithChan invokes the scdn.DescribeScdnVerifyContent API asynchronously
func (client *Client) DescribeScdnVerifyContentWithChan(request *DescribeScdnVerifyContentRequest) (<-chan *DescribeScdnVerifyContentResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnVerifyContentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnVerifyContent(request)
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

// DescribeScdnVerifyContentWithCallback invokes the scdn.DescribeScdnVerifyContent API asynchronously
func (client *Client) DescribeScdnVerifyContentWithCallback(request *DescribeScdnVerifyContentRequest, callback func(response *DescribeScdnVerifyContentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnVerifyContentResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnVerifyContent(request)
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

// DescribeScdnVerifyContentRequest is the request struct for api DescribeScdnVerifyContent
type DescribeScdnVerifyContentRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
}

// DescribeScdnVerifyContentResponse is the response struct for api DescribeScdnVerifyContent
type DescribeScdnVerifyContentResponse struct {
	*responses.BaseResponse
	Content   string `json:"Content" xml:"Content"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeScdnVerifyContentRequest creates a request to invoke DescribeScdnVerifyContent API
func CreateDescribeScdnVerifyContentRequest() (request *DescribeScdnVerifyContentRequest) {
	request = &DescribeScdnVerifyContentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnVerifyContent", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnVerifyContentResponse creates a response to parse from DescribeScdnVerifyContent response
func CreateDescribeScdnVerifyContentResponse() (response *DescribeScdnVerifyContentResponse) {
	response = &DescribeScdnVerifyContentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
