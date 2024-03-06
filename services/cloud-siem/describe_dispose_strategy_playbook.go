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

// DescribeDisposeStrategyPlaybook invokes the cloud_siem.DescribeDisposeStrategyPlaybook API synchronously
func (client *Client) DescribeDisposeStrategyPlaybook(request *DescribeDisposeStrategyPlaybookRequest) (response *DescribeDisposeStrategyPlaybookResponse, err error) {
	response = CreateDescribeDisposeStrategyPlaybookResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDisposeStrategyPlaybookWithChan invokes the cloud_siem.DescribeDisposeStrategyPlaybook API asynchronously
func (client *Client) DescribeDisposeStrategyPlaybookWithChan(request *DescribeDisposeStrategyPlaybookRequest) (<-chan *DescribeDisposeStrategyPlaybookResponse, <-chan error) {
	responseChan := make(chan *DescribeDisposeStrategyPlaybookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDisposeStrategyPlaybook(request)
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

// DescribeDisposeStrategyPlaybookWithCallback invokes the cloud_siem.DescribeDisposeStrategyPlaybook API asynchronously
func (client *Client) DescribeDisposeStrategyPlaybookWithCallback(request *DescribeDisposeStrategyPlaybookRequest, callback func(response *DescribeDisposeStrategyPlaybookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDisposeStrategyPlaybookResponse
		var err error
		defer close(result)
		response, err = client.DescribeDisposeStrategyPlaybook(request)
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

// DescribeDisposeStrategyPlaybookRequest is the request struct for api DescribeDisposeStrategyPlaybook
type DescribeDisposeStrategyPlaybookRequest struct {
	*requests.RpcRequest
	StartTime requests.Integer `position:"Body" name:"StartTime"`
	EndTime   requests.Integer `position:"Body" name:"EndTime"`
}

// DescribeDisposeStrategyPlaybookResponse is the response struct for api DescribeDisposeStrategyPlaybook
type DescribeDisposeStrategyPlaybookResponse struct {
	*responses.BaseResponse
	Success   bool       `json:"Success" xml:"Success"`
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeDisposeStrategyPlaybookRequest creates a request to invoke DescribeDisposeStrategyPlaybook API
func CreateDescribeDisposeStrategyPlaybookRequest() (request *DescribeDisposeStrategyPlaybookRequest) {
	request = &DescribeDisposeStrategyPlaybookRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeDisposeStrategyPlaybook", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDisposeStrategyPlaybookResponse creates a response to parse from DescribeDisposeStrategyPlaybook response
func CreateDescribeDisposeStrategyPlaybookResponse() (response *DescribeDisposeStrategyPlaybookResponse) {
	response = &DescribeDisposeStrategyPlaybookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
