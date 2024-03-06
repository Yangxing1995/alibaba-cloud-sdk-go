package market

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

// DescribeProjectNodes invokes the market.DescribeProjectNodes API synchronously
func (client *Client) DescribeProjectNodes(request *DescribeProjectNodesRequest) (response *DescribeProjectNodesResponse, err error) {
	response = CreateDescribeProjectNodesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProjectNodesWithChan invokes the market.DescribeProjectNodes API asynchronously
func (client *Client) DescribeProjectNodesWithChan(request *DescribeProjectNodesRequest) (<-chan *DescribeProjectNodesResponse, <-chan error) {
	responseChan := make(chan *DescribeProjectNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProjectNodes(request)
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

// DescribeProjectNodesWithCallback invokes the market.DescribeProjectNodes API asynchronously
func (client *Client) DescribeProjectNodesWithCallback(request *DescribeProjectNodesRequest, callback func(response *DescribeProjectNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProjectNodesResponse
		var err error
		defer close(result)
		response, err = client.DescribeProjectNodes(request)
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

// DescribeProjectNodesRequest is the request struct for api DescribeProjectNodes
type DescribeProjectNodesRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeProjectNodesResponse is the response struct for api DescribeProjectNodes
type DescribeProjectNodesResponse struct {
	*responses.BaseResponse
	Success   bool          `json:"Success" xml:"Success"`
	RequestId string        `json:"RequestId" xml:"RequestId"`
	Result    []ProjectNode `json:"Result" xml:"Result"`
}

// CreateDescribeProjectNodesRequest creates a request to invoke DescribeProjectNodes API
func CreateDescribeProjectNodesRequest() (request *DescribeProjectNodesRequest) {
	request = &DescribeProjectNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Market", "2015-11-01", "DescribeProjectNodes", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeProjectNodesResponse creates a response to parse from DescribeProjectNodes response
func CreateDescribeProjectNodesResponse() (response *DescribeProjectNodesResponse) {
	response = &DescribeProjectNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
