package cloud_siem

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

// DescribeLogSource invokes the cloud_siem.DescribeLogSource API synchronously
func (client *Client) DescribeLogSource(request *DescribeLogSourceRequest) (response *DescribeLogSourceResponse, err error) {
	response = CreateDescribeLogSourceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLogSourceWithChan invokes the cloud_siem.DescribeLogSource API asynchronously
func (client *Client) DescribeLogSourceWithChan(request *DescribeLogSourceRequest) (<-chan *DescribeLogSourceResponse, <-chan error) {
	responseChan := make(chan *DescribeLogSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLogSource(request)
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

// DescribeLogSourceWithCallback invokes the cloud_siem.DescribeLogSource API asynchronously
func (client *Client) DescribeLogSourceWithCallback(request *DescribeLogSourceRequest, callback func(response *DescribeLogSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLogSourceResponse
		var err error
		defer close(result)
		response, err = client.DescribeLogSource(request)
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

// DescribeLogSourceRequest is the request struct for api DescribeLogSource
type DescribeLogSourceRequest struct {
	*requests.RpcRequest
	LogType string `position:"Body" name:"LogType"`
}

// DescribeLogSourceResponse is the response struct for api DescribeLogSource
type DescribeLogSourceResponse struct {
	*responses.BaseResponse
	Success   bool       `json:"Success" xml:"Success"`
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeLogSourceRequest creates a request to invoke DescribeLogSource API
func CreateDescribeLogSourceRequest() (request *DescribeLogSourceRequest) {
	request = &DescribeLogSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeLogSource", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLogSourceResponse creates a response to parse from DescribeLogSource response
func CreateDescribeLogSourceResponse() (response *DescribeLogSourceResponse) {
	response = &DescribeLogSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
