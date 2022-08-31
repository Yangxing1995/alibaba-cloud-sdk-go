package cdn

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

// DescribeRealtimeLogAuthorized invokes the cdn.DescribeRealtimeLogAuthorized API synchronously
func (client *Client) DescribeRealtimeLogAuthorized(request *DescribeRealtimeLogAuthorizedRequest) (response *DescribeRealtimeLogAuthorizedResponse, err error) {
	response = CreateDescribeRealtimeLogAuthorizedResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRealtimeLogAuthorizedWithChan invokes the cdn.DescribeRealtimeLogAuthorized API asynchronously
func (client *Client) DescribeRealtimeLogAuthorizedWithChan(request *DescribeRealtimeLogAuthorizedRequest) (<-chan *DescribeRealtimeLogAuthorizedResponse, <-chan error) {
	responseChan := make(chan *DescribeRealtimeLogAuthorizedResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRealtimeLogAuthorized(request)
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

// DescribeRealtimeLogAuthorizedWithCallback invokes the cdn.DescribeRealtimeLogAuthorized API asynchronously
func (client *Client) DescribeRealtimeLogAuthorizedWithCallback(request *DescribeRealtimeLogAuthorizedRequest, callback func(response *DescribeRealtimeLogAuthorizedResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRealtimeLogAuthorizedResponse
		var err error
		defer close(result)
		response, err = client.DescribeRealtimeLogAuthorized(request)
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

// DescribeRealtimeLogAuthorizedRequest is the request struct for api DescribeRealtimeLogAuthorized
type DescribeRealtimeLogAuthorizedRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeRealtimeLogAuthorizedResponse is the response struct for api DescribeRealtimeLogAuthorized
type DescribeRealtimeLogAuthorizedResponse struct {
	*responses.BaseResponse
	AuthorizedStatus string `json:"AuthorizedStatus" xml:"AuthorizedStatus"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeRealtimeLogAuthorizedRequest creates a request to invoke DescribeRealtimeLogAuthorized API
func CreateDescribeRealtimeLogAuthorizedRequest() (request *DescribeRealtimeLogAuthorizedRequest) {
	request = &DescribeRealtimeLogAuthorizedRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeRealtimeLogAuthorized", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeRealtimeLogAuthorizedResponse creates a response to parse from DescribeRealtimeLogAuthorized response
func CreateDescribeRealtimeLogAuthorizedResponse() (response *DescribeRealtimeLogAuthorizedResponse) {
	response = &DescribeRealtimeLogAuthorizedResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
