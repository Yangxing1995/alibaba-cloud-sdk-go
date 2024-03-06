package swas_open

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

// DescribeCloudAssistantStatus invokes the swas_open.DescribeCloudAssistantStatus API synchronously
func (client *Client) DescribeCloudAssistantStatus(request *DescribeCloudAssistantStatusRequest) (response *DescribeCloudAssistantStatusResponse, err error) {
	response = CreateDescribeCloudAssistantStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudAssistantStatusWithChan invokes the swas_open.DescribeCloudAssistantStatus API asynchronously
func (client *Client) DescribeCloudAssistantStatusWithChan(request *DescribeCloudAssistantStatusRequest) (<-chan *DescribeCloudAssistantStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudAssistantStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudAssistantStatus(request)
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

// DescribeCloudAssistantStatusWithCallback invokes the swas_open.DescribeCloudAssistantStatus API asynchronously
func (client *Client) DescribeCloudAssistantStatusWithCallback(request *DescribeCloudAssistantStatusRequest, callback func(response *DescribeCloudAssistantStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudAssistantStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudAssistantStatus(request)
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

// DescribeCloudAssistantStatusRequest is the request struct for api DescribeCloudAssistantStatus
type DescribeCloudAssistantStatusRequest struct {
	*requests.RpcRequest
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	InstanceIds *[]string        `position:"Query" name:"InstanceIds"  type:"Json"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeCloudAssistantStatusResponse is the response struct for api DescribeCloudAssistantStatus
type DescribeCloudAssistantStatusResponse struct {
	*responses.BaseResponse
	RequestId            string   `json:"RequestId" xml:"RequestId"`
	TotalCount           int      `json:"TotalCount" xml:"TotalCount"`
	PageSize             int      `json:"PageSize" xml:"PageSize"`
	PageNumber           int      `json:"PageNumber" xml:"PageNumber"`
	CloudAssistantStatus []Status `json:"CloudAssistantStatus" xml:"CloudAssistantStatus"`
}

// CreateDescribeCloudAssistantStatusRequest creates a request to invoke DescribeCloudAssistantStatus API
func CreateDescribeCloudAssistantStatusRequest() (request *DescribeCloudAssistantStatusRequest) {
	request = &DescribeCloudAssistantStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DescribeCloudAssistantStatus", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCloudAssistantStatusResponse creates a response to parse from DescribeCloudAssistantStatus response
func CreateDescribeCloudAssistantStatusResponse() (response *DescribeCloudAssistantStatusResponse) {
	response = &DescribeCloudAssistantStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
