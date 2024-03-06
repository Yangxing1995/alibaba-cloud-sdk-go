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

// DescribeTables invokes the clickhouse.DescribeTables API synchronously
func (client *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
	response = CreateDescribeTablesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTablesWithChan invokes the clickhouse.DescribeTables API asynchronously
func (client *Client) DescribeTablesWithChan(request *DescribeTablesRequest) (<-chan *DescribeTablesResponse, <-chan error) {
	responseChan := make(chan *DescribeTablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTables(request)
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

// DescribeTablesWithCallback invokes the clickhouse.DescribeTables API asynchronously
func (client *Client) DescribeTablesWithCallback(request *DescribeTablesRequest, callback func(response *DescribeTablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTablesResponse
		var err error
		defer close(result)
		response, err = client.DescribeTables(request)
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

// DescribeTablesRequest is the request struct for api DescribeTables
type DescribeTablesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TableName            string           `position:"Query" name:"TableName"`
	SchemaName           string           `position:"Query" name:"SchemaName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeTablesResponse is the response struct for api DescribeTables
type DescribeTablesResponse struct {
	*responses.BaseResponse
	RequestId string                `json:"RequestId" xml:"RequestId"`
	Items     ItemsInDescribeTables `json:"Items" xml:"Items"`
}

// CreateDescribeTablesRequest creates a request to invoke DescribeTables API
func CreateDescribeTablesRequest() (request *DescribeTablesRequest) {
	request = &DescribeTablesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeTables", "service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTablesResponse creates a response to parse from DescribeTables response
func CreateDescribeTablesResponse() (response *DescribeTablesResponse) {
	response = &DescribeTablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
