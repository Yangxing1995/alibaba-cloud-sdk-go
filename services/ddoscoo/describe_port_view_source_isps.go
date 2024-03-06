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

// DescribePortViewSourceIsps invokes the ddoscoo.DescribePortViewSourceIsps API synchronously
func (client *Client) DescribePortViewSourceIsps(request *DescribePortViewSourceIspsRequest) (response *DescribePortViewSourceIspsResponse, err error) {
	response = CreateDescribePortViewSourceIspsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePortViewSourceIspsWithChan invokes the ddoscoo.DescribePortViewSourceIsps API asynchronously
func (client *Client) DescribePortViewSourceIspsWithChan(request *DescribePortViewSourceIspsRequest) (<-chan *DescribePortViewSourceIspsResponse, <-chan error) {
	responseChan := make(chan *DescribePortViewSourceIspsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePortViewSourceIsps(request)
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

// DescribePortViewSourceIspsWithCallback invokes the ddoscoo.DescribePortViewSourceIsps API asynchronously
func (client *Client) DescribePortViewSourceIspsWithCallback(request *DescribePortViewSourceIspsRequest, callback func(response *DescribePortViewSourceIspsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePortViewSourceIspsResponse
		var err error
		defer close(result)
		response, err = client.DescribePortViewSourceIsps(request)
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

// DescribePortViewSourceIspsRequest is the request struct for api DescribePortViewSourceIsps
type DescribePortViewSourceIspsRequest struct {
	*requests.RpcRequest
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	InstanceIds     *[]string        `position:"Query" name:"InstanceIds"  type:"Repeated"`
}

// DescribePortViewSourceIspsResponse is the response struct for api DescribePortViewSourceIsps
type DescribePortViewSourceIspsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Isps      []Isp  `json:"Isps" xml:"Isps"`
}

// CreateDescribePortViewSourceIspsRequest creates a request to invoke DescribePortViewSourceIsps API
func CreateDescribePortViewSourceIspsRequest() (request *DescribePortViewSourceIspsRequest) {
	request = &DescribePortViewSourceIspsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribePortViewSourceIsps", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePortViewSourceIspsResponse creates a response to parse from DescribePortViewSourceIsps response
func CreateDescribePortViewSourceIspsResponse() (response *DescribePortViewSourceIspsResponse) {
	response = &DescribePortViewSourceIspsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
