package cloudapi

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

// DescribeInstanceHttpCode invokes the cloudapi.DescribeInstanceHttpCode API synchronously
func (client *Client) DescribeInstanceHttpCode(request *DescribeInstanceHttpCodeRequest) (response *DescribeInstanceHttpCodeResponse, err error) {
	response = CreateDescribeInstanceHttpCodeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceHttpCodeWithChan invokes the cloudapi.DescribeInstanceHttpCode API asynchronously
func (client *Client) DescribeInstanceHttpCodeWithChan(request *DescribeInstanceHttpCodeRequest) (<-chan *DescribeInstanceHttpCodeResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceHttpCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceHttpCode(request)
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

// DescribeInstanceHttpCodeWithCallback invokes the cloudapi.DescribeInstanceHttpCode API asynchronously
func (client *Client) DescribeInstanceHttpCodeWithCallback(request *DescribeInstanceHttpCodeRequest, callback func(response *DescribeInstanceHttpCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceHttpCodeResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceHttpCode(request)
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

// DescribeInstanceHttpCodeRequest is the request struct for api DescribeInstanceHttpCode
type DescribeInstanceHttpCodeRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DescribeInstanceHttpCodeResponse is the response struct for api DescribeInstanceHttpCode
type DescribeInstanceHttpCodeResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	InstanceHttpCode InstanceHttpCode `json:"InstanceHttpCode" xml:"InstanceHttpCode"`
}

// CreateDescribeInstanceHttpCodeRequest creates a request to invoke DescribeInstanceHttpCode API
func CreateDescribeInstanceHttpCodeRequest() (request *DescribeInstanceHttpCodeRequest) {
	request = &DescribeInstanceHttpCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeInstanceHttpCode", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceHttpCodeResponse creates a response to parse from DescribeInstanceHttpCode response
func CreateDescribeInstanceHttpCodeResponse() (response *DescribeInstanceHttpCodeResponse) {
	response = &DescribeInstanceHttpCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
