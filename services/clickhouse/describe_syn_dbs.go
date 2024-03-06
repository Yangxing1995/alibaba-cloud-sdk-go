package clickhouse

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

// DescribeSynDbs invokes the clickhouse.DescribeSynDbs API synchronously
func (client *Client) DescribeSynDbs(request *DescribeSynDbsRequest) (response *DescribeSynDbsResponse, err error) {
	response = CreateDescribeSynDbsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSynDbsWithChan invokes the clickhouse.DescribeSynDbs API asynchronously
func (client *Client) DescribeSynDbsWithChan(request *DescribeSynDbsRequest) (<-chan *DescribeSynDbsResponse, <-chan error) {
	responseChan := make(chan *DescribeSynDbsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSynDbs(request)
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

// DescribeSynDbsWithCallback invokes the clickhouse.DescribeSynDbs API asynchronously
func (client *Client) DescribeSynDbsWithCallback(request *DescribeSynDbsRequest, callback func(response *DescribeSynDbsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSynDbsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSynDbs(request)
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

// DescribeSynDbsRequest is the request struct for api DescribeSynDbs
type DescribeSynDbsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DbClusterId          string           `position:"Query" name:"DbClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSynDbsResponse is the response struct for api DescribeSynDbs
type DescribeSynDbsResponse struct {
	*responses.BaseResponse
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	TotalCount int     `json:"TotalCount" xml:"TotalCount"`
	PageSize   int     `json:"PageSize" xml:"PageSize"`
	PageNumber int     `json:"PageNumber" xml:"PageNumber"`
	SynDbs     []SynDb `json:"SynDbs" xml:"SynDbs"`
}

// CreateDescribeSynDbsRequest creates a request to invoke DescribeSynDbs API
func CreateDescribeSynDbsRequest() (request *DescribeSynDbsRequest) {
	request = &DescribeSynDbsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeSynDbs", "service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSynDbsResponse creates a response to parse from DescribeSynDbs response
func CreateDescribeSynDbsResponse() (response *DescribeSynDbsResponse) {
	response = &DescribeSynDbsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
