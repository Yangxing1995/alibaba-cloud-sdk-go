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

// DescribeAutoCcBlacklist invokes the ddoscoo.DescribeAutoCcBlacklist API synchronously
func (client *Client) DescribeAutoCcBlacklist(request *DescribeAutoCcBlacklistRequest) (response *DescribeAutoCcBlacklistResponse, err error) {
	response = CreateDescribeAutoCcBlacklistResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAutoCcBlacklistWithChan invokes the ddoscoo.DescribeAutoCcBlacklist API asynchronously
func (client *Client) DescribeAutoCcBlacklistWithChan(request *DescribeAutoCcBlacklistRequest) (<-chan *DescribeAutoCcBlacklistResponse, <-chan error) {
	responseChan := make(chan *DescribeAutoCcBlacklistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAutoCcBlacklist(request)
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

// DescribeAutoCcBlacklistWithCallback invokes the ddoscoo.DescribeAutoCcBlacklist API asynchronously
func (client *Client) DescribeAutoCcBlacklistWithCallback(request *DescribeAutoCcBlacklistRequest, callback func(response *DescribeAutoCcBlacklistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAutoCcBlacklistResponse
		var err error
		defer close(result)
		response, err = client.DescribeAutoCcBlacklist(request)
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

// DescribeAutoCcBlacklistRequest is the request struct for api DescribeAutoCcBlacklist
type DescribeAutoCcBlacklistRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	SourceIp   string           `position:"Query" name:"SourceIp"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	KeyWord    string           `position:"Query" name:"KeyWord"`
}

// DescribeAutoCcBlacklistResponse is the response struct for api DescribeAutoCcBlacklist
type DescribeAutoCcBlacklistResponse struct {
	*responses.BaseResponse
	TotalCount      int64                 `json:"TotalCount" xml:"TotalCount"`
	RequestId       string                `json:"RequestId" xml:"RequestId"`
	AutoCcBlacklist []AutoCcBlacklistItem `json:"AutoCcBlacklist" xml:"AutoCcBlacklist"`
}

// CreateDescribeAutoCcBlacklistRequest creates a request to invoke DescribeAutoCcBlacklist API
func CreateDescribeAutoCcBlacklistRequest() (request *DescribeAutoCcBlacklistRequest) {
	request = &DescribeAutoCcBlacklistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeAutoCcBlacklist", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAutoCcBlacklistResponse creates a response to parse from DescribeAutoCcBlacklist response
func CreateDescribeAutoCcBlacklistResponse() (response *DescribeAutoCcBlacklistResponse) {
	response = &DescribeAutoCcBlacklistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
