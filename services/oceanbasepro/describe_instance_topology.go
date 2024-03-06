package oceanbasepro

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

// DescribeInstanceTopology invokes the oceanbasepro.DescribeInstanceTopology API synchronously
func (client *Client) DescribeInstanceTopology(request *DescribeInstanceTopologyRequest) (response *DescribeInstanceTopologyResponse, err error) {
	response = CreateDescribeInstanceTopologyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceTopologyWithChan invokes the oceanbasepro.DescribeInstanceTopology API asynchronously
func (client *Client) DescribeInstanceTopologyWithChan(request *DescribeInstanceTopologyRequest) (<-chan *DescribeInstanceTopologyResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceTopologyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceTopology(request)
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

// DescribeInstanceTopologyWithCallback invokes the oceanbasepro.DescribeInstanceTopology API asynchronously
func (client *Client) DescribeInstanceTopologyWithCallback(request *DescribeInstanceTopologyRequest, callback func(response *DescribeInstanceTopologyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceTopologyResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceTopology(request)
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

// DescribeInstanceTopologyRequest is the request struct for api DescribeInstanceTopology
type DescribeInstanceTopologyRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Body" name:"InstanceId"`
	ReplicaTypes string `position:"Body" name:"ReplicaTypes"`
}

// DescribeInstanceTopologyResponse is the response struct for api DescribeInstanceTopology
type DescribeInstanceTopologyResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	InstanceTopology InstanceTopology `json:"InstanceTopology" xml:"InstanceTopology"`
}

// CreateDescribeInstanceTopologyRequest creates a request to invoke DescribeInstanceTopology API
func CreateDescribeInstanceTopologyRequest() (request *DescribeInstanceTopologyRequest) {
	request = &DescribeInstanceTopologyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeInstanceTopology", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceTopologyResponse creates a response to parse from DescribeInstanceTopology response
func CreateDescribeInstanceTopologyResponse() (response *DescribeInstanceTopologyResponse) {
	response = &DescribeInstanceTopologyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
