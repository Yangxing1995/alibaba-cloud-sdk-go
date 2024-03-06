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

// DescribeLogFields invokes the cloud_siem.DescribeLogFields API synchronously
func (client *Client) DescribeLogFields(request *DescribeLogFieldsRequest) (response *DescribeLogFieldsResponse, err error) {
	response = CreateDescribeLogFieldsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLogFieldsWithChan invokes the cloud_siem.DescribeLogFields API asynchronously
func (client *Client) DescribeLogFieldsWithChan(request *DescribeLogFieldsRequest) (<-chan *DescribeLogFieldsResponse, <-chan error) {
	responseChan := make(chan *DescribeLogFieldsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLogFields(request)
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

// DescribeLogFieldsWithCallback invokes the cloud_siem.DescribeLogFields API asynchronously
func (client *Client) DescribeLogFieldsWithCallback(request *DescribeLogFieldsRequest, callback func(response *DescribeLogFieldsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLogFieldsResponse
		var err error
		defer close(result)
		response, err = client.DescribeLogFields(request)
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

// DescribeLogFieldsRequest is the request struct for api DescribeLogFields
type DescribeLogFieldsRequest struct {
	*requests.RpcRequest
	LogType   string `position:"Body" name:"LogType"`
	LogSource string `position:"Body" name:"LogSource"`
}

// DescribeLogFieldsResponse is the response struct for api DescribeLogFields
type DescribeLogFieldsResponse struct {
	*responses.BaseResponse
	Success   bool       `json:"Success" xml:"Success"`
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeLogFieldsRequest creates a request to invoke DescribeLogFields API
func CreateDescribeLogFieldsRequest() (request *DescribeLogFieldsRequest) {
	request = &DescribeLogFieldsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeLogFields", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLogFieldsResponse creates a response to parse from DescribeLogFields response
func CreateDescribeLogFieldsResponse() (response *DescribeLogFieldsResponse) {
	response = &DescribeLogFieldsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
